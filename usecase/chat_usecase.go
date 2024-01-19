package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fadilahonespot/chatbot/entity"
	"github.com/fadilahonespot/chatbot/repository/cached"
	"github.com/fadilahonespot/chatbot/repository/http/chatgbt"
	"github.com/fadilahonespot/chatbot/repository/mysql"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/fadilahonespot/library/errors"
	"github.com/sashabaranov/go-openai"
)

type ChatUsecase interface {
	ChatQuestion(ctx context.Context, userId int, req dto.ChatQuestionRequest) (resp dto.ChatQuestionResponse, err error)
	GetHistoryChat(ctx context.Context, userId int) (resp []dto.ChatHistoryResponse, err error)
}

type defaultChatUsecase struct {
	userRepo      mysql.UserRepository
	chatRepo      mysql.ChatRepository
	openAiWrapper chatgbt.OpenAIWrapper
	cacheWrapper  cached.CacheWrapper
}

const (
	KeyHistory = "getHistory"
)

// NewChatUsecase creates a new instance of ChatUsecase
func NewChatUsecase(userRepo mysql.UserRepository, chatRepo mysql.ChatRepository, openAiWrapper chatgbt.OpenAIWrapper, cacheWrapper cached.CacheWrapper) ChatUsecase {
	return &defaultChatUsecase{
		userRepo:      userRepo,
		chatRepo:      chatRepo,
		openAiWrapper: openAiWrapper,
		cacheWrapper:  cacheWrapper,
	}
}

// ChatQuestion handles the chat question from the user
func (s *defaultChatUsecase) ChatQuestion(ctx context.Context, userId int, req dto.ChatQuestionRequest) (resp dto.ChatQuestionResponse, err error) {
	// get user data
	userData, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		logger.Error(ctx, "user not found")
		err = errors.SetError(http.StatusBadRequest, "user not found")
		return
	}

	// create the initial chat request
	reqChat := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Halo! Perkenalkan aku adalah ChatBot Assistant. Bagimana aku bisa membantumu hari ini?",
			},
		},
	}

	// append the user's question to the request
	reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Question,
	})

	// get the previous chat request from cache
	key := fmt.Sprintf("ChatBot_%v", userId)
	value, _ := s.cacheWrapper.Get(ctx, key)
	if value != "" {
		// unmarshall the previous chat request
		var reqData openai.ChatCompletionRequest
		err = json.Unmarshal([]byte(value), &reqData)
		if err != nil {
			fmt.Println("error unmarshalling chat: ", err.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		// append the user's question to the previous chat request
		reqData.Messages = append(reqData.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: req.Question,
		})

		// update the previous chat request with the new question
		reqChat = reqData
	} else {
		// get the history of chats with the user
		historyData, errRes := s.chatRepo.GetByUserId(ctx, userData.ID)
		if errRes != nil {
			logger.Error(ctx, "error getting chat history")
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		// append the history of chats to the initial chat request
		for i := 0; i < len(historyData); i++ {
			reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: historyData[i].Message,
			})
		}

		// append the user's question to the end of the chat request
		reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: req.Question,
		})
	}

	// generate the response from OpenAI
	dataResp, err := s.openAiWrapper.GenerateText(ctx, reqChat)
	if err != nil {
		fmt.Println("error generating text: ", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// append the response from OpenAI to the chat request
	reqChat.Messages = append(reqChat.Messages, dataResp.Choices[0].Message)

	// marshall the chat request and cache it
	dataByte, _ := json.Marshal(reqChat)
	s.cacheWrapper.Set(ctx, key, string(dataByte), time.Duration(60)*time.Minute)

	// get the response from OpenAI
	answer := dataResp.Choices[0].Message.Content

	// create the history of chats
	reqHistory := []entity.Chat{
		{
			UserID:  userData.ID,
			Name:    userData.Name,
			Message: req.Question,
		},
		{
			UserID:  userData.ID,
			Name:    "Bot",
			Message: answer,
		},
	}

	// start a transaction
	tx := s.chatRepo.BeginsTrans()
	// loop through the history of chats and create them
	for i := 0; i < len(reqHistory); i++ {
		err = s.chatRepo.Create(ctx, tx, &reqHistory[i])
		if err != nil {
			s.chatRepo.Rollback(tx)
			logger.Error(ctx, "error creating chat history", err.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	// commit the transaction
	s.chatRepo.Commit(tx)

	// delete the history of chats cache
	keyHistory := fmt.Sprintf("%v_%v", KeyHistory, userId)
	s.cacheWrapper.Delete(ctx, keyHistory)

	// set the response
	resp.Answer = answer
	return
}

// GetHistoryChat returns the history of chats between the user and the bot
func (s *defaultChatUsecase) GetHistoryChat(ctx context.Context, userId int) (resp []dto.ChatHistoryResponse, err error) {
	key := fmt.Sprintf("%v_%v", KeyHistory, userId)
	value, _ := s.cacheWrapper.Get(ctx, key)
	if value == "" {
		// get history from database
		historyData, errRes := s.chatRepo.GetHistoryChatByUserId(ctx, userId)
		if errRes != nil {
			logger.Error(ctx, "error getting chat history", errRes.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		// convert to dto
		for i := 0; i < len(historyData); i++ {
			resp = append(resp, dto.ChatHistoryResponse{
				Id:      historyData[i].ID,
				Name:    historyData[i].Name,
				Message: historyData[i].Message,
			})
		}

		// convert to json and cache
		dataByte, _ := json.Marshal(resp)
		s.cacheWrapper.Set(ctx, key, string(dataByte), time.Duration(60)*time.Minute)
		return
	}

	// unmarshall from cache
	err = json.Unmarshal([]byte(value), &resp)
	if err != nil {
		logger.Error(ctx, "error unmarshalling chat: ", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

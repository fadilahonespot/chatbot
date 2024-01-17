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

func NewChatUsecase(userRepo mysql.UserRepository, chatRepo mysql.ChatRepository, openAiWrapper chatgbt.OpenAIWrapper, cacheWrapper cached.CacheWrapper) ChatUsecase {
	return &defaultChatUsecase{
		userRepo:      userRepo,
		chatRepo:      chatRepo,
		openAiWrapper: openAiWrapper,
		cacheWrapper:  cacheWrapper,
	}
}

func (s *defaultChatUsecase) ChatQuestion(ctx context.Context, userId int, req dto.ChatQuestionRequest) (resp dto.ChatQuestionResponse, err error) {
	userData, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		logger.Error(ctx, "user not found")
		err = errors.SetError(http.StatusBadRequest, "user not found")
		return
	}
	reqChat := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Halo! Perkenalkan aku adalah ChatBot Assistant. Bagimana aku bisa membantumu hari ini?",
			},
		},
	}

	reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Question,
	})

	key := fmt.Sprintf("ChatBot_%v", userId)
	value, _ := s.cacheWrapper.Get(ctx, key)
	if value != "" {
		var reqData openai.ChatCompletionRequest
		err = json.Unmarshal([]byte(value), &reqData)
		if err != nil {
			fmt.Println("error unmarshalling chat: ", err.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		reqData.Messages = append(reqData.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: req.Question,
		})

		reqChat = reqData
	} else {
		historyData, errRes := s.chatRepo.GetByUserId(ctx, userData.ID)
		if errRes != nil {
			logger.Error(ctx, "error getting chat history")
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		for i := 0; i < len(historyData); i++ {
			reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: historyData[i].Message,
			})
		}

		reqChat.Messages = append(reqChat.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: req.Question,
		})
	}

	dataResp, err := s.openAiWrapper.GenerateText(ctx, reqChat)
	if err != nil {
		fmt.Println("error generating text: ", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	reqChat.Messages = append(reqChat.Messages, dataResp.Choices[0].Message)
	dataByte, _ := json.Marshal(reqChat)
	s.cacheWrapper.Set(ctx, key, string(dataByte), time.Duration(60)*time.Minute)

	answer := dataResp.Choices[0].Message.Content
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

	tx := s.chatRepo.BeginsTrans()
	for i := 0; i < len(reqHistory); i++ {
		err = s.chatRepo.Create(ctx, tx, &reqHistory[i])
		if err != nil {
			s.chatRepo.Rollback(tx)
			logger.Error(ctx, "error creating chat history", err.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	s.chatRepo.Commit(tx)
	keyHistory := fmt.Sprintf("%v_%v", KeyHistory, userId)
	s.cacheWrapper.Delete(ctx, keyHistory)
	resp.Answer = answer
	return
}

func (s *defaultChatUsecase) GetHistoryChat(ctx context.Context, userId int) (resp []dto.ChatHistoryResponse, err error) {
	key := fmt.Sprintf("%v_%v", KeyHistory, userId)
	value, _ := s.cacheWrapper.Get(ctx, key)
	if value == "" {
		historyData, errRes := s.chatRepo.GetHistoryChatByUserId(ctx, userId)
		if errRes != nil {
			logger.Error(ctx, "error getting chat history", errRes.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		for i := 0; i < len(historyData); i++ {
			resp = append(resp, dto.ChatHistoryResponse{
				Id:      historyData[i].ID,
				Name:    historyData[i].Name,
				Message: historyData[i].Message,
			})
		}

		dataByte, _ := json.Marshal(resp)
		s.cacheWrapper.Set(ctx, key, string(dataByte), time.Duration(60)*time.Minute)
		return
	}

	err = json.Unmarshal([]byte(value), &resp)
	if err != nil {
		logger.Error(ctx, "error unmarshalling chat: ", err.Error())
        err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        return
	}

	return
}

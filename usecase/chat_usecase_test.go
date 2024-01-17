package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/fadilahonespot/chatbot/entity"
	"github.com/fadilahonespot/chatbot/mocks"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
)

func Test_defaultChatUsecase_ChatQuestion(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()

	type args struct {
		ctx    context.Context
		userId int
		req    dto.ChatQuestionRequest
	}
	tests := []struct {
		name             string
		args             args
		getUserResp      *entity.User
		getUserErr       error
		cacheGetResp     string
		cacheGetErr      error
		getChatResp      []entity.Chat
		getChatErr       error
		generateTextResp openai.ChatCompletionResponse
		generateTextErr  error
		createChatErr    error
		wantResp         dto.ChatQuestionResponse
		wantErr          bool
	}{
		{
			name: "user not found",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			getUserErr: errors.New("user not found"),
			wantErr:    true,
		},
		{
			name: "data value in cache not empty but unmarshal error",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			cacheGetResp: "users value",
			wantErr:      true,
		},
		{
			name: "data value in cache is empty but get data chat by userId error",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			getUserResp: &entity.User{
				ID: 1,
			},
			getChatErr: errors.New("error getting data chat"),
			wantErr:    true,
		},
		{
			name: "data value in cache is empty but generate text error",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			getUserResp: &entity.User{
				ID: 1,
			},
			getChatResp: []entity.Chat{
				{
					ID:      1,
					Name:    "testing",
					Message: "testing",
				},
			},
			generateTextErr: errors.New("error generate text"),
			wantErr:         true,
		},
		{
			name: "data value in cache is empty but create chat error",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			getUserResp: &entity.User{
				ID: 1,
			},
			getChatResp: []entity.Chat{
				{
					ID:      1,
					Name:    "testing",
					Message: "testing",
				},
			},
			generateTextResp: openai.ChatCompletionResponse{
				Choices: []openai.ChatCompletionChoice{
					{
						Message: openai.ChatCompletionMessage{
							Role:    openai.ChatMessageRoleUser,
							Content: "cara memasak nasi goreng yaitu nasi harus di goreng",
						},
					},
				},
			},
			createChatErr: errors.New("create chat error"),
			wantErr:       true,
		},
		{
			name: "succes generate chat",
			args: args{
				ctx:    ctx,
				userId: 1,
				req: dto.ChatQuestionRequest{
					Question: "bagaimana cara memasak nasi goreng?",
				},
			},
			getUserResp: &entity.User{
				ID: 1,
			},
			getChatResp: []entity.Chat{
				{
					ID:      1,
					Name:    "testing",
					Message: "testing",
				},
			},
			generateTextResp: openai.ChatCompletionResponse{
				Choices: []openai.ChatCompletionChoice{
					{
						Message: openai.ChatCompletionMessage{
							Role:    openai.ChatMessageRoleUser,
							Content: "cara memasak nasi goreng yaitu nasi harus di goreng",
						},
					},
				},
			},
			wantResp: dto.ChatQuestionResponse{
				Answer: "cara memasak nasi goreng yaitu nasi harus di goreng",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)
			chatRepo := new(mocks.ChatRepository)
			openAiWrapper := new(mocks.OpenAIWrapper)
			cacheWrapper := new(mocks.CacheWrapper)

			mockDb := utils.MockGorm()

			userRepo.On("GetUserById", mock.Anything, mock.Anything).Return(tt.getUserResp, tt.getUserErr).Once()
			cacheWrapper.On("Get", mock.Anything, mock.Anything).Return(tt.cacheGetResp, tt.cacheGetErr).Once()
			chatRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(tt.getChatResp, tt.getChatErr).Once()
			openAiWrapper.On("GenerateText", mock.Anything, mock.Anything).Return(tt.generateTextResp, tt.generateTextErr).Once()
			chatRepo.On("BeginsTrans").Return(mockDb).Once()
			chatRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(tt.createChatErr)
			chatRepo.On("Rollback", mock.Anything).Return(nil).Once()
			chatRepo.On("Commit", mock.Anything).Return(nil).Once()
			cacheWrapper.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
			cacheWrapper.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

			s := NewChatUsecase(userRepo, chatRepo, openAiWrapper, cacheWrapper)
			gotResp, err := s.ChatQuestion(tt.args.ctx, tt.args.userId, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultChatUsecase.ChatQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultChatUsecase.ChatQuestion() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_defaultChatUsecase_GetHistoryChat(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()
	getChat := []entity.Chat{
		{
			ID:      1,
			Name:    "testing",
			Message: "testing",
		},
	}

	historyChatResp := []dto.ChatHistoryResponse{
		{
			Id:      1,
			Name:    "testing",
			Message: "testing",
		},
	}

	historyByte, _ := json.Marshal(historyChatResp)
	type args struct {
		ctx    context.Context
		userId int
	}
	tests := []struct {
		name           string
		args           args
		cacheGetResp   string
		cacheGetErr    error
		getHistoryResp []entity.Chat
		getHistoryErr  error
		wantResp       []dto.ChatHistoryResponse
		wantErr        bool
	}{
		{
			name: "get history error",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			cacheGetResp:  "",
			getHistoryErr: errors.New("error getting chat history"),
			wantErr:       true,
		},
		{
			name:           "cache is null but succes get data chat history",
			cacheGetResp:   "",
			getHistoryResp: getChat,
			wantResp:       historyChatResp,
			wantErr:        false,
		},
		{
			name: "cache is not empty but unmarshal error",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			cacheGetResp: "datas",
			wantErr:      true,
		},
		{
			name: "cache is not empty but succes get data chat history",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			cacheGetResp: string(historyByte),
			wantResp:     historyChatResp,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)
			chatRepo := new(mocks.ChatRepository)
			openAiWrapper := new(mocks.OpenAIWrapper)
			cacheWrapper := new(mocks.CacheWrapper)

			cacheWrapper.On("Get", mock.Anything, mock.Anything).Return(tt.cacheGetResp, tt.cacheGetErr).Once()
			chatRepo.On("GetHistoryChatByUserId", mock.Anything, mock.Anything).Return(tt.getHistoryResp, tt.getHistoryErr).Once()
			cacheWrapper.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

			s := NewChatUsecase(userRepo, chatRepo, openAiWrapper, cacheWrapper)
			gotResp, err := s.GetHistoryChat(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultChatUsecase.GetHistoryChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultChatUsecase.GetHistoryChat() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

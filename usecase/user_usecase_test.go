package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fadilahonespot/chatbot/entity"
	"github.com/fadilahonespot/chatbot/mocks"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/stretchr/testify/mock"
)

func Test_defaultUserUsecase_Register(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()

	type args struct {
		ctx context.Context
		req dto.RegisterRequest
	}
	tests := []struct {
		name          string
		args          args
		getUserResp   *entity.User
		getUserErr    error
		createUserErr error
		wantErr       bool
	}{
		{
			name: "email not valid",
			args: args{
				ctx: ctx,
				req: dto.RegisterRequest{
					Email: "thkshhdhdd",
				},
			},
			wantErr: true,
		},
		{
			name: "email already exist",
			args: args{
				ctx: ctx,
				req: dto.RegisterRequest{
					Email: "testing@gmail.com",
				},
			},
			getUserResp: &entity.User{
				Email: "testing@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "create user error",
			args: args{
				ctx: ctx,
				req: dto.RegisterRequest{
					Email: "testing@gmail.com",
				},
			},
			getUserResp:   &entity.User{},
			createUserErr: errors.New("create user error"),
			wantErr:       true,
		},
		{
			name: "create user success",
			args: args{
				ctx: ctx,
				req: dto.RegisterRequest{
					Email: "testing@gmail.com",
				},
			},
			getUserResp: &entity.User{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)
			userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(tt.getUserResp, tt.getUserErr).Once()
			userRepo.On("Create", mock.Anything, mock.Anything).Return(tt.createUserErr).Once()

			s := NewUserUsecase(userRepo)
			if err := s.Register(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultUserUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultUserUsecase_Login(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()

	type args struct {
		ctx context.Context
		req dto.LoginRequest
	}
	tests := []struct {
		name        string
		args        args
		getUserResp *entity.User
		getUserErr  error
		wantResp    dto.LoginResponse
		wantErr     bool
	}{
		{
			name: "email not valid",
			args: args{
				ctx: ctx,
				req: dto.LoginRequest{
					Email: "wdwdweewdd",
				},
			},
			wantResp: dto.LoginResponse{},
			wantErr:  true,
		},
		{
			name: "user email not found",
			args: args{
				ctx: ctx,
				req: dto.LoginRequest{
					Email: "testing@gmail.com",
				},
			},
			getUserErr: errors.New("not found"),
			wantResp:   dto.LoginResponse{},
			wantErr:    true,
		},
		{
			name: "compare password not same",
			args: args{
				ctx: ctx,
				req: dto.LoginRequest{
					Email:    "testing@gmail.com",
					Password: "12345",
				},
			},
			getUserResp: &entity.User{
				ID:       1,
				Email:    "testing@gmail.com",
				Name:     "testing",
				Password: "$2a$10$kpRxNQtm.QaT9hr4tPhfPuLkWE73bcnxKFJIMjHCxfWnfavaIAcAW",
			},
			wantResp: dto.LoginResponse{},
			wantErr:  true,
		},
		{
			name: "success login",
			args: args{
				ctx: ctx,
				req: dto.LoginRequest{
					Email:    "testing@gmail.com",
					Password: "123456",
				},
			},
			getUserResp: &entity.User{
				ID:       1,
				Email:    "testing@gmail.com",
				Name:     "testing",
				Password: "$2a$10$kpRxNQtm.QaT9hr4tPhfPuLkWE73bcnxKFJIMjHCxfWnfavaIAcAW",
			},
			wantResp: dto.LoginResponse{
				Id:    1,
				Email: "testing@gmail.com",
				Name:  "testing",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)
			userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(tt.getUserResp, tt.getUserErr).Once()

			s := NewUserUsecase(userRepo)
			gotResp, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultUserUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp.Email, tt.wantResp.Email) {
				t.Errorf("defaultUserUsecase.Login() = %v, want %v", gotResp.Email, tt.wantResp.Email)
			}
			if !reflect.DeepEqual(gotResp.Name, tt.wantResp.Name) {
				t.Errorf("defaultUserUsecase.Login() = %v, want %v", gotResp.Name, tt.wantResp.Name)
			}
			if !reflect.DeepEqual(gotResp.Id, tt.wantResp.Id) {
				t.Errorf("defaultUserUsecase.Login() = %v, want %v", gotResp.Id, tt.wantResp.Id)
			}
		})
	}
}

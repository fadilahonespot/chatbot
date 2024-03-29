// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/fadilahonespot/chatbot/usecase/dto"
	mock "github.com/stretchr/testify/mock"
)

// ChatUsecase is an autogenerated mock type for the ChatUsecase type
type ChatUsecase struct {
	mock.Mock
}

// ChatQuestion provides a mock function with given fields: ctx, userId, req
func (_m *ChatUsecase) ChatQuestion(ctx context.Context, userId int, req dto.ChatQuestionRequest) (dto.ChatQuestionResponse, error) {
	ret := _m.Called(ctx, userId, req)

	if len(ret) == 0 {
		panic("no return value specified for ChatQuestion")
	}

	var r0 dto.ChatQuestionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, dto.ChatQuestionRequest) (dto.ChatQuestionResponse, error)); ok {
		return rf(ctx, userId, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, dto.ChatQuestionRequest) dto.ChatQuestionResponse); ok {
		r0 = rf(ctx, userId, req)
	} else {
		r0 = ret.Get(0).(dto.ChatQuestionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, dto.ChatQuestionRequest) error); ok {
		r1 = rf(ctx, userId, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHistoryChat provides a mock function with given fields: ctx, userId
func (_m *ChatUsecase) GetHistoryChat(ctx context.Context, userId int) ([]dto.ChatHistoryResponse, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetHistoryChat")
	}

	var r0 []dto.ChatHistoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]dto.ChatHistoryResponse, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []dto.ChatHistoryResponse); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ChatHistoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatUsecase creates a new instance of ChatUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatUsecase {
	mock := &ChatUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

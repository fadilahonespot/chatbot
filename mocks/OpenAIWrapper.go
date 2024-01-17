// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	openai "github.com/sashabaranov/go-openai"
	mock "github.com/stretchr/testify/mock"
)

// OpenAIWrapper is an autogenerated mock type for the OpenAIWrapper type
type OpenAIWrapper struct {
	mock.Mock
}

// GenerateText provides a mock function with given fields: ctx, req
func (_m *OpenAIWrapper) GenerateText(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GenerateText")
	}

	var r0 openai.ChatCompletionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, openai.ChatCompletionRequest) openai.ChatCompletionResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(openai.ChatCompletionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, openai.ChatCompletionRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOpenAIWrapper creates a new instance of OpenAIWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOpenAIWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *OpenAIWrapper {
	mock := &OpenAIWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

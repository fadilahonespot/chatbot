// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fadilahonespot/chatbot/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// ChatRepository is an autogenerated mock type for the ChatRepository type
type ChatRepository struct {
	mock.Mock
}

// BeginsTrans provides a mock function with given fields:
func (_m *ChatRepository) BeginsTrans() *gorm.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BeginsTrans")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Commit provides a mock function with given fields: tx
func (_m *ChatRepository) Commit(tx *gorm.DB) error {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, tx, req
func (_m *ChatRepository) Create(ctx context.Context, tx *gorm.DB, req *entity.Chat) error {
	ret := _m.Called(ctx, tx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *entity.Chat) error); ok {
		r0 = rf(ctx, tx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUserId provides a mock function with given fields: ctx, userId
func (_m *ChatRepository) GetByUserId(ctx context.Context, userId int) ([]entity.Chat, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserId")
	}

	var r0 []entity.Chat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Chat, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Chat); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Chat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHistoryChatByUserId provides a mock function with given fields: ctx, userId
func (_m *ChatRepository) GetHistoryChatByUserId(ctx context.Context, userId int) ([]entity.Chat, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetHistoryChatByUserId")
	}

	var r0 []entity.Chat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Chat, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Chat); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Chat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollback provides a mock function with given fields: tx
func (_m *ChatRepository) Rollback(tx *gorm.DB) error {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChatRepository creates a new instance of ChatRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatRepository {
	mock := &ChatRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

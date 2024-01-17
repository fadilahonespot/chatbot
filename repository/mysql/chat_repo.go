package mysql

import (
	"context"

	"github.com/fadilahonespot/chatbot/entity"
	"github.com/fadilahonespot/chatbot/utils/paginate"
	"gorm.io/gorm"
)

type ChatRepository interface {
	BeginsTrans() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, req *entity.Chat) (err error) 
	GetByUserId(ctx context.Context, userId int) (resp []entity.Chat, err error)
	GetHistoryChatByUserId(ctx context.Context, userId int) (resp []entity.Chat, err error) 
}

type defaultChatRepo struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &defaultChatRepo{db}
}

func (s *defaultChatRepo) BeginsTrans() *gorm.DB {
	return s.db.Begin()
}

func (s *defaultChatRepo) Commit(tx *gorm.DB) error {
    return tx.Commit().Error
}

func (s *defaultChatRepo) Rollback(tx *gorm.DB) error {
    return tx.Rollback().Error
}

func (s *defaultChatRepo) Create(ctx context.Context, tx *gorm.DB, req *entity.Chat) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultChatRepo) GetByUserId(ctx context.Context, userId int) (resp []entity.Chat, err error) {
	err = s.db.WithContext(ctx).
		Scopes(paginate.Paginate(1, 20)).
		Find(&resp, "user_id = ?", userId).Error
	return
}

func (s *defaultChatRepo) GetHistoryChatByUserId(ctx context.Context, userId int) (resp []entity.Chat, err error) {
	err = s.db.WithContext(ctx).
		Scopes(paginate.Paginate(1, 20)).Order("id DESC").
		Find(&resp, "user_id = ?", userId).Error
	return
}

package mysql

import (
	"context"

	"github.com/fadilahonespot/chatbot/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, req *entity.User) (err error)
	GetUserById(ctx context.Context, id int) (resp *entity.User, err error)
	GetUserByEmail(ctx context.Context, email string) (resp *entity.User, err error)
}

type defultUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &defultUserRepo{db}
}

func (s *defultUserRepo) Create(ctx context.Context, req *entity.User) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defultUserRepo) GetUserById(ctx context.Context, id int) (resp *entity.User, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (s *defultUserRepo) GetUserByEmail(ctx context.Context, email string) (resp *entity.User, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "email = ?", email).Error
	return
}

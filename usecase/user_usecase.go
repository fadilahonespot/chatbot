package usecase

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/fadilahonespot/chatbot/entity"
	"github.com/fadilahonespot/chatbot/repository/mysql"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils/constrans"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/fadilahonespot/library/errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (err error)
	Login(ctx context.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error)
}

type defaultUserUsecase struct {
	userRepo mysql.UserRepository
}

func NewUserUsecase(userRepo mysql.UserRepository) UserUsecase {
	return &defaultUserUsecase{
		userRepo: userRepo,
	}
}

func (s *defaultUserUsecase) Register(ctx context.Context, req dto.RegisterRequest) (err error) {
	if !strings.Contains(req.Email, "@") {
		logger.Error(ctx, "email not valid")
		err = errors.SetError(http.StatusBadRequest, "email not valid")
		return
	}

	userData, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userData.Email != "" {
		logger.Error(ctx, "email already exists")
		err = errors.SetError(http.StatusBadRequest, "email already exists")
		return
	}

	createUser := entity.User{
		Email:    req.Email,
		Password: hashPassword(req.Password),
		Name:     req.Name,
	}
	err = s.userRepo.Create(ctx, &createUser)
	if err != nil {
		logger.Error(ctx, "error creating user", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	return
}

func (s *defaultUserUsecase) Login(ctx context.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error) {
	if !strings.Contains(req.Email, "@") {
		logger.Error(ctx, "email not valid")
		err = errors.SetError(http.StatusBadRequest, "email not valid")
		return
	}

	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "error getting user", err.Error())
		err = errors.SetError(http.StatusBadRequest, "user not found")
		return
	}

	err = comparePassword(userData.Password, req.Password)
	if err != nil {
		logger.Error(ctx, "password not valid")
		err = errors.SetError(http.StatusBadRequest, "password not valid")
		return
	}

	token := createToken(userData.ID)
	resp = dto.LoginResponse{
		Id:          userData.ID,
		Name:        userData.Name,
		Email:       userData.Email,
		AccessToken: token,
	}

	return
}

func createToken(userId int) string {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(constrans.JwtSecret))
	return t
}

func hashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func comparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

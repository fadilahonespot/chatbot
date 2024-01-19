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

// Register registers a new user.
//
// If the email is already registered, returns an error.
//
// Parameters:
//   - ctx: context.Context
//   - req: dto.RegisterRequest
//
// Returns:
//   - error: error if any
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

// Login authenticates a user and returns an access token.
//
// Parameters:
//   - ctx: context.Context
//   - req: dto.LoginRequest
//
// Returns:
//   - resp: dto.LoginResponse
//   - err: error
func (s *defaultUserUsecase) Login(ctx context.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error) {
	if !strings.Contains(req.Email, "@") {
		logger.Error(ctx, "email not valid")
		err = errors.SetError(http.StatusBadRequest, "email not valid")
		return
	}

	// Get the user data from the database.
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "error getting user", err.Error())
		err = errors.SetError(http.StatusBadRequest, "user not found")
		return
	}

	// Compare the provided password with the stored hash.
	err = comparePassword(userData.Password, req.Password)
	if err != nil {
		logger.Error(ctx, "password not valid")
		err = errors.SetError(http.StatusBadRequest, "password not valid")
		return
	}

	// Create a JSON Web Token and return it to the user.
	token := createToken(userData.ID)
	resp = dto.LoginResponse{
		Id:          userData.ID,
		Name:        userData.Name,
		Email:       userData.Email,
		AccessToken: token,
	}

	return
}

// createToken creates a JSON Web Token (JWT) with the given user ID.
func createToken(userId int) string {
	// Create the claims for the token.
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key.
	t, _ := token.SignedString([]byte(constrans.JwtSecret))
	return t
}

// hashPassword generates a salted and hashed password using the bcrypt algorithm
func hashPassword(password string) string {
	// Generate a salt using the bcrypt package
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// Convert the salt to a string and return it along with the hashed password
	return string(result)
}

// comparePassword compares a given password with a hashed password
func comparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

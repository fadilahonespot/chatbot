package handler

import (
	"net/http"

	"github.com/fadilahonespot/chatbot/usecase"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/fadilahonespot/chatbot/utils/request"
	"github.com/fadilahonespot/chatbot/utils/response"
	"github.com/fadilahonespot/library/errors"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 200 {object} dto.RegisterResponse
// @Failure 400,500 {object} errors.HTTPError
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req dto.RegisterRequest
    ctx := r.Context()
    err := request.GetRequestFromContext(ctx, &req)
    if err != nil {
        logger.Error(ctx, "failed get request", err.Error())
        err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        response.ResponseError(w, err)
        return
    }

    err = h.userUsecase.Register(ctx, req)
    if err != nil {
        response.ResponseError(w, err)
        return
    }

    response.ResponseSuccess(w, nil)
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400,500 {object} errors.HTTPError
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req dto.LoginRequest
    ctx := r.Context()
    err := request.GetRequestFromContext(ctx, &req)
    if err != nil {
        logger.Error(ctx, "failed get request", err.Error())
        err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        response.ResponseError(w, err)
        return
    }

    resp, err := h.userUsecase.Login(ctx, req)
    if err != nil {
        response.ResponseError(w, err)
        return
    }

    response.ResponseSuccess(w, resp)
}

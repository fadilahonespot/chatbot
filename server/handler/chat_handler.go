package handler

import (
	"net/http"

	"github.com/fadilahonespot/chatbot/usecase"
	"github.com/fadilahonespot/chatbot/usecase/dto"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/fadilahonespot/chatbot/utils/request"
	"github.com/fadilahonespot/chatbot/utils/response"
	"github.com/fadilahonespot/library/errors"
	"github.com/spf13/cast"
)

type ChatHandler struct {
	chatUsecase usecase.ChatUsecase
}

func NewChatHandler(chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		chatUsecase: chatUsecase,
	}
}

func (h *ChatHandler) Chat(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req dto.ChatQuestionRequest
		ctx := r.Context()
		err := request.GetRequestFromContext(ctx, &req)
		if err != nil {
			logger.Error(ctx, "failed get request", err.Error())
			err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			response.ResponseError(w, err)
			return
		}

		userId := ctx.Value("userId")
		resp, err := h.chatUsecase.ChatQuestion(ctx, cast.ToInt(userId), req)
		if err != nil {
			response.ResponseError(w, err)
			return
		}

		response.ResponseSuccess(w, resp)

	case http.MethodGet:
		ctx := r.Context()

		userId := ctx.Value("userId")
		resp, err := h.chatUsecase.GetHistoryChat(ctx, cast.ToInt(userId))
		if err != nil {
			response.ResponseError(w, err)
			return
		}

		response.ResponseSuccess(w, resp)

	default:
		err := errors.SetError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		response.ResponseError(w, err)
	}
}

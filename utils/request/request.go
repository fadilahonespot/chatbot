package request

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fadilahonespot/chatbot/utils/logger"
)

const (
	RequestBodyKey = "requestBody"
)

func SetRequestInContext(ctx context.Context, reqBody []byte) context.Context {
	ctx = context.WithValue(ctx, RequestBodyKey, reqBody)
	return ctx
}
func GetRequestFromContext(ctx context.Context, req any) error {
	requestData, ok := ctx.Value(RequestBodyKey).([]byte)
	if !ok {
		return errors.New("failed to get request body")
	}

	err := json.Unmarshal(requestData, req)
	if err != nil {
		return err
	}

	logger.Info(ctx, "[REQUEST]", req)
	return nil
}

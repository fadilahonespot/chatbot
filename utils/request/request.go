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

// SetRequestInContext sets the request data in the context.
func SetRequestInContext(ctx context.Context, reqBody []byte) context.Context {
    ctx = context.WithValue(ctx, RequestBodyKey, reqBody)
    return ctx
}
// GetRequestFromContext retrieves the request data from the context.
// It decodes the request data into the given struct.
func GetRequestFromContext(ctx context.Context, req any) error {
    // requestData is the request data stored in the context.
    requestData, ok := ctx.Value(RequestBodyKey).([]byte)
    if !ok {
        return errors.New("failed to get request body")
    }

    // Unmarshal decodes the JSON-encoded data in the request body into the value pointed to by req.
    err := json.Unmarshal(requestData, req)
    if err != nil {
        return err
    }

    // Log the request data.
    logger.Info(ctx, "[REQUEST]", req)
    return nil
}

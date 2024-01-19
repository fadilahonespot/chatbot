package chatgbt

import (
	"context"
	"fmt"
	"os"

	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/sashabaranov/go-openai"
)

type wrapper struct {
	client *openai.Client
}

// NewWrapper creates a new instance of the OpenAI wrapper
func NewWrapper() OpenAIWrapper {
	client := openai.NewClient(os.Getenv("OPEN_AI_TOKEN"))
	return &wrapper{
		client: client,
	}
}

// GenerateText creates a new instance of the OpenAI wrapper
func (w *wrapper) GenerateText(ctx context.Context, req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {
	// log the request
	logger.Info(ctx, "GenerateText REQUEST", req)

	// make the request to OpenAI
	resp, err = w.client.CreateChatCompletion(ctx, req)
	if err != nil {
		// handle the error
		err = fmt.Errorf("chat completion error: %s", err.Error())
		return
	}

	// log the response
	logger.Info(ctx, "GenerateText RESPONSE", resp)

	// return the response
	return
}

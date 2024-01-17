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

func NewWrapper() OpenAIWrapper {
	client := openai.NewClient(os.Getenv("OPEN_AI_TOKEN"))
	return &wrapper{
		client: client,
	}
}

func (w *wrapper) GenerateText(ctx context.Context, req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {
	logger.Info(ctx, "GeneratText REQUEST", req)
	
	resp, err = w.client.CreateChatCompletion(ctx, req)
	if err != nil {
		err = fmt.Errorf("chat completion error: %s", err.Error())
		return
	}

	logger.Info(ctx, "GenerateText RESPONSE", resp)

	return
}

package chatgbt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIWrapper interface {
	GenerateText(ctx context.Context, req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error)
}

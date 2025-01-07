package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	logger *zap.Logger
	client *genai.Client
}

type Cfg struct {
	Logger *zap.Logger
	ApiKey string
}

func (c Cfg) Validate() error {
	if c.Logger == nil {
		return fmt.Errorf("Logger is missing")
	}

	if c.ApiKey == "" {
		return fmt.Errorf("ApiKey is missing")
	}

	return nil
}

func NewGeminiClient(cfg Cfg) (GeminiClient, error) {
	if err := cfg.Validate(); err != nil {
		return GeminiClient{}, fmt.Errorf("cfg.Validate(): %w", err)
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.ApiKey))
	if err != nil {
		return GeminiClient{}, err
	}

	return GeminiClient{
		client: client,
		logger: cfg.Logger,
	}, nil
}

func (c GeminiClient) Close() {
	c.client.Close()
}

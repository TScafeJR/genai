package gemini

import (
	"context"
	"fmt"
	"sync"

	"github.com/TScafeJR/genai/ratelimit"
	"github.com/google/generative-ai-go/genai"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	logger      *zap.Logger
	client      *genai.Client
	rateLimiter *rate.Limiter
	mu          sync.Mutex
}

type Cfg struct {
	Logger    *zap.Logger
	ApiKey    string
	RateLimit *ratelimit.RateLimit
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

	if cfg.RateLimit == nil {
		cfg.RateLimit = &ratelimit.RateLimit{
			Timeframe: 10,
			MaxCalls:  100,
		}
	}

	ratePerSecond := float64(cfg.RateLimit.MaxCalls) / cfg.RateLimit.Timeframe.Seconds()

	return GeminiClient{
		client:      client,
		logger:      cfg.Logger,
		rateLimiter: rate.NewLimiter(rate.Limit(ratePerSecond), cfg.RateLimit.MaxCalls),
	}, nil
}

func (c *GeminiClient) limit(ctx context.Context) error {
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter error: %w", err)
	}
	return nil
}

func (c *GeminiClient) Close() {
	c.client.Close()
}

package deepseek

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type DeepseekClient struct {
	logger     *zap.Logger
	client     *Client
	url        string
	httpClient *http.Client
	apiKey     string
}

type Models struct {
	Text       string
	MultiModal string
}

type Cfg struct {
	Logger     *zap.Logger
	ApiKey     string
	Models     Models
	Local      bool
	HttpClient *http.Client
}

func (c Cfg) Validate() error {
	if c.Logger == nil {
		return fmt.Errorf("Logger is missing")
	}

	if !c.Local && c.ApiKey == "" {
		return fmt.Errorf("ApiKey is missing")
	}

	if c.HttpClient == nil {
		return fmt.Errorf("HttpClient is missing")
	}

	return nil
}

func pickServerUrl(local bool) string {
	if local {
		return "http://localhost:11434"
	}
	return "https://api.deepseek.com/v1"
}

func NewDeepseekClient(cfg Cfg) (DeepseekClient, error) {
	if err := cfg.Validate(); err != nil {
		return DeepseekClient{}, fmt.Errorf("cfg.Validate(): %w", err)
	}

	url := pickServerUrl(cfg.Local)
	client, err := NewClient(url, WithHTTPClient(cfg.HttpClient))
	if err != nil {
		return DeepseekClient{}, err
	}

	return DeepseekClient{
		logger: cfg.Logger,
		url:    url,
		client: client,
	}, nil
}

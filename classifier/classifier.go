package classifier

import (
	"context"
)

type Client interface {
	Close()
	GenerateContent(ctx context.Context, p Prompt) (Classification, error)
	CreateEmbedding(ctx context.Context, p Prompt) ([]float32, error)
}

type Image struct {
	Data    []byte
	ImgType string
}

type Prompt struct {
	Images []Image
	Text   []string
}

type Model interface {
	GenerateContent(ctx context.Context, p ...Prompt) (string, error)
}

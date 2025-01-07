package classifier

import (
	"context"
)

type FakeClient struct {
}

func (c FakeClient) Close() {
}

func (c FakeClient) GenerateContent(ctx context.Context, p Prompt) (Classification, error) {
	return Classification{}, nil
}

func (c FakeClient) CreateEmbedding(ctx context.Context, p Prompt) ([]float32, error) {
	return []float32{}, nil
}

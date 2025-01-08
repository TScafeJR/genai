package gemini

import (
	"context"
	"fmt"

	"github.com/TScafeJR/genai/classifier"
	"go.uber.org/zap"
)

func cleanPrompt(p classifier.Prompt) classifier.Prompt {
	return classifier.Prompt{
		Text: p.Text,
	}
}

func (g *GeminiClient) CreateEmbedding(ctx context.Context, p classifier.Prompt) ([]float32, error) {
	g.logger.Debug("CreateEmbedding", zap.Any("prompt", p))

	if err := g.limit(ctx); err != nil {
		return nil, err
	}
	// need to remove images from the prompt for the time being since our model doesn't support images
	prompt := toGenAIPrompt(cleanPrompt(p))
	em := g.client.EmbeddingModel("embedding-001")
	res, err := em.EmbedContent(ctx, prompt...)
	if err != nil {
		return []float32{}, fmt.Errorf("model.EmbedContent(): %w", err)
	}

	if res.Embedding == nil {
		return []float32{}, nil
	}

	return res.Embedding.Values, nil
}

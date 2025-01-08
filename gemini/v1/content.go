package gemini

import (
	"context"
	"fmt"

	"github.com/TScafeJR/genai/classifier"
	"github.com/google/generative-ai-go/genai"
	"go.uber.org/zap"
)

func getTextFromPart(p genai.Part) (string, bool) {
	if textPart, ok := p.(genai.Text); ok {
		return string(textPart), true
	}
	return "", false
}

func (c *GeminiClient) GenerateContent(ctx context.Context, p classifier.Prompt) (classifier.Classification, error) {
	c.logger.Debug("GenerateContent", zap.Any("prompt", p))
	if err := c.limit(ctx); err != nil {
		return classifier.Classification{}, err
	}
	var model *genai.GenerativeModel
	if len(p.Images) == 0 {
		model = c.client.GenerativeModel("gemini-1.5-flash")
	} else {
		model = c.client.GenerativeModel("gemini-pro-vision")
	}

	prompt := toGenAIPrompt(p)

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		return classifier.Classification{}, fmt.Errorf("model.GenerateContent(): %w", err)
	}

	classification := classifier.Classification{
		Parts: []string{},
	}

	if len(resp.Candidates) > 0 {
		for _, candidate := range resp.Candidates {
			if candidate.Content != nil {
				content := candidate.Content
				for _, part := range content.Parts {
					s, ok := getTextFromPart(part)
					if ok {
						classification.Parts = append(classification.Parts, s)
					}
				}
			}
		}
	}

	return classification, nil
}

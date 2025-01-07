package gemini

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/TScafeJR/genai/classifier"
	"github.com/google/generative-ai-go/genai"
)

func getTextFromPart(p genai.Part) (string, bool) {
	if textPart, ok := p.(genai.Text); ok {
		return string(textPart), true
	}
	return "", false
}

func parseTopics(s string) []string {
	if s == "" {
		return []string{}
	}
	topics := strings.Split(s, ", ")
	re := regexp.MustCompile(`\d+\.\s*`)

	for i, topic := range topics {
		topics[i] = re.ReplaceAllString(topic, "")
	}

	return topics
}

func (c GeminiClient) GenerateContent(ctx context.Context, p classifier.Prompt) (classifier.Classification, error) {
	var model *genai.GenerativeModel
	if len(p.Images) == 0 {
		model = c.client.GenerativeModel("gemini-pro")
	} else {
		model = c.client.GenerativeModel("gemini-pro-vision")
	}

	prompt := toGenAIPrompt(p)

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		return classifier.Classification{}, fmt.Errorf("model.GenerateContent(): %w", err)
	}

	var classification classifier.Classification

	if len(resp.Candidates) > 0 {
		for _, candidate := range resp.Candidates {
			if candidate.Content != nil {
				content := candidate.Content
				for _, part := range content.Parts {
					s, ok := getTextFromPart(part)
					if ok {
						topics := parseTopics(s)
						classification.Topics = append(classification.Topics, topics...)
					}
				}
			}
		}
	}

	return classification, nil
}

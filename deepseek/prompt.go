package deepseek

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/TScafeJR/genai/classifier"
	"go.uber.org/zap"
)

func promptToReq(p classifier.Prompt) GenerateContentJSONRequestBody {
	var prompt string

	for _, text := range p.Text {
		prompt += fmt.Sprintf("%s ", text)
	}

	return GenerateContentJSONRequestBody{
		Prompt: &prompt,
	}
}

func (d DeepseekClient) isLocal() bool {
	return strings.Contains(d.url, "localhost")
}

func respToClassification(_ *http.Response) (classifier.Classification, error) {
	return classifier.Classification{}, nil
}

func (d DeepseekClient) GenerateContent(ctx context.Context, p classifier.Prompt) (classifier.Classification, error) {
	d.logger.Debug("GenerateContent", zap.Any("prompt", p))

	req := promptToReq(p)
	resp, err := d.client.GenerateContent(ctx, req)
	if err != nil {
		return classifier.Classification{}, fmt.Errorf("client.GenerateContent(): %w", err)
	}

	aiResp, err := respToClassification(resp)
	if err != nil {
		return classifier.Classification{}, fmt.Errorf("respToClassification(): %w", err)
	}

	return aiResp, nil
}

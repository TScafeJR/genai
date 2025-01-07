package gemini

import (
	"github.com/TScafeJR/genai/classifier"
	"github.com/google/generative-ai-go/genai"
)

func toGenAIPrompt(p classifier.Prompt) []genai.Part {
	prompt := []genai.Part{}

	for _, txt := range p.Text {
		prompt = append(prompt, genai.Text(txt))
	}

	for _, img := range p.Images {
		prompt = append(prompt, genai.ImageData(img.ImgType, img.Data))
	}

	return prompt
}

package genai

import (
	"regexp"
	"strings"
)

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

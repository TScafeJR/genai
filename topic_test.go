package genai

import (
	"reflect"
	"testing"
)

func Test_ParseTopics(t *testing.T) {
	tests := []struct {
		desc  string
		input string
		want  []string
	}{
		{
			desc:  "empty string",
			input: "",
			want:  []string{},
		},
		{
			desc:  "simple csv field",
			input: "EmotionalExpression, Empathy, Communication, Relationships, Masculinity, Vulnerability, Gender, SupportSystems, EmotionalIntelligence, Openness, MentalHealth, ListeningSkills, EmotionalSupport, SupportiveRelationships, InterpersonalRelationships, NonJudgmentalListening, EmotionalWellbeing, HealthyRelationships, MaleEmotionalExpression",
			want: []string{
				"EmotionalExpression",
				"Empathy",
				"Communication",
				"Relationships",
				"Masculinity",
				"Vulnerability",
				"Gender",
				"SupportSystems",
				"EmotionalIntelligence",
				"Openness",
				"MentalHealth",
				"ListeningSkills",
				"EmotionalSupport",
				"SupportiveRelationships",
				"InterpersonalRelationships",
				"NonJudgmentalListening",
				"EmotionalWellbeing",
				"HealthyRelationships",
				"MaleEmotionalExpression",
			},
		},
		{
			desc:  "numbered csv field",
			input: "1. EmotionalExpression, 2. Empathy, 3. Communication, 4. Relationships, 5. Masculinity, 6. Vulnerability, 7. Gender, 8. SupportSystems, 9. EmotionalIntelligence, 10. Openness, 11. MentalHealth, 12. ListeningSkills, 13. EmotionalSupport, 14. SupportiveRelationships, 15. InterpersonalRelationships, 16. NonJudgmentalListening, 17. EmotionalWellbeing, 18. HealthyRelationships, 19. MaleEmotionalExpression",
			want: []string{
				"EmotionalExpression",
				"Empathy",
				"Communication",
				"Relationships",
				"Masculinity",
				"Vulnerability",
				"Gender",
				"SupportSystems",
				"EmotionalIntelligence",
				"Openness",
				"MentalHealth",
				"ListeningSkills",
				"EmotionalSupport",
				"SupportiveRelationships",
				"InterpersonalRelationships",
				"NonJudgmentalListening",
				"EmotionalWellbeing",
				"HealthyRelationships",
				"MaleEmotionalExpression",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := parseTopics(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("parseTopics(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

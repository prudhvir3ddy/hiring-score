package utils

import (
	"github.com/prudhvir3ddy/hiring-score/models"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		name     string
		input    models.Candidate
		expected float64
	}{
		{
			name: "High scoring full-time candidate",
			input: models.Candidate{
				WorkAvailability: []string{"full-time"},
				Skills:           []string{"Golang", "TypeScript", "React", "Next.js"},
				WorkExperiences: []models.WorkExperience{
					{Company: "A"}, {Company: "B"}, {Company: "C"},
					{Company: "D"}, {Company: "E"},
				},
				Education: models.Education{
					HighestLevel: "Master's Degree",
					Degrees:      []models.Degree{{IsTop25: true}},
				},
				AnnualSalaryExpectation: models.SalaryExpectation{FullTime: "$70000"},
			},
			expected: 65, // 15 (exp) + 20 (skills: 8 base + 12 priority) + 15 (edu: 10 masters + 5 top25) + 5 (availability) + 10 (salary)
		},
		{
			name: "Part-time only candidate",
			input: models.Candidate{
				WorkAvailability: []string{"part-time"},
				Skills:           []string{"Golang", "React"},
				WorkExperiences:  []models.WorkExperience{{Company: "A"}},
				Education: models.Education{
					HighestLevel: "Bachelor's Degree",
					Degrees:      []models.Degree{{IsTop25: true}},
				},
				AnnualSalaryExpectation: models.SalaryExpectation{FullTime: "$80000"},
			},
			expected: 31, // 3 (exp) + 10 (skills: 4 base + 6 priority) + 10 (edu: 5 bachelors + 5 top25) + 0 (availability) + 8 (salary)
		},
		{
			name: "Mid-level salary candidate",
			input: models.Candidate{
				WorkAvailability: []string{"full-time"},
				Skills:           []string{"JavaScript", "HTML", "CSS"},
				WorkExperiences:  []models.WorkExperience{{Company: "A"}, {Company: "B"}},
				Education: models.Education{
					HighestLevel: "Bachelor's Degree",
					Degrees:      []models.Degree{{IsTop50: true}},
				},
				AnnualSalaryExpectation: models.SalaryExpectation{FullTime: "$85000"},
			},
			expected: 36, // 6 (exp) + 9 (skills: 6 base + 3 priority) + 8 (edu: 5 bachelors + 3 top50) + 5 (availability) + 8 (salary)
		},
		{
			name: "Maximum score candidate",
			input: models.Candidate{
				WorkAvailability: []string{"full-time"},
				Skills: []string{
					"Golang", "TypeScript", "React", "Next.js", "TailwindCSS",
					"MongoDB", "PostgreSQL", "Python", "Docker", "AWS",
				},
				WorkExperiences: []models.WorkExperience{
					{Company: "A"}, {Company: "B"}, {Company: "C"},
					{Company: "D"}, {Company: "E"}, {Company: "F"},
					{Company: "G"}, {Company: "H"}, {Company: "I"}, {Company: "J"},
				},
				Education: models.Education{
					HighestLevel: "Master's Degree",
					Degrees:      []models.Degree{{IsTop25: true}},
				},
				AnnualSalaryExpectation: models.SalaryExpectation{FullTime: "$70000"},
			},
			expected: 100, // 30 (exp) + 40 (skills) + 15 (edu: 10 masters + 5 top25) + 5 (availability) + 10 (salary)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateScore(tt.input)
			if got != tt.expected {
				t.Errorf("CalculateScore() = %v, want %v", got, tt.expected)
			}
		})
	}
}

package utils

import (
	"github.com/prudhvir3ddy/hiring-score/models"
	"reflect"
	"testing"
)

func TestSortCandidates(t *testing.T) {
	tests := []struct {
		name     string
		input    []models.Candidate
		expected []models.Candidate
	}{
		{
			name: "Different scores",
			input: []models.Candidate{
				{Score: 80},
				{Score: 90},
				{Score: 70},
			},
			expected: []models.Candidate{
				{Score: 90},
				{Score: 80},
				{Score: 70},
			},
		},
		{
			name: "Same score, different work experience",
			input: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}}, // 2 experiences
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}}, // 1 experience
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}, {}}, // 3 experiences
				},
			},
			expected: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}, {}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}},
				},
			},
		},
		{
			name: "Same score and work experience, different skills count",
			input: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react", "typescript"},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go"},
				},
			},
			expected: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react", "typescript"},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go"},
				},
			},
		},
		{
			name: "Same score, work experience and skills, different education",
			input: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{IsTop50: true}}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{IsTop25: true}}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{}}},
				},
			},
			expected: []models.Candidate{
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{IsTop25: true}}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{IsTop50: true}}},
				},
				{
					Score:           80,
					WorkExperiences: []models.WorkExperience{{}, {}},
					Skills:          []string{"go", "react"},
					Education:       models.Education{Degrees: []models.Degree{{}}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortCandidates(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("SortCandidates() = %v, want %v", tt.input, tt.expected)
			}
		})
	}
}

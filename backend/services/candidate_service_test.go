// backend/services/candidate_service_test.go
package services

import (
	"testing"

	"github.com/prudhvir3ddy/hiring-score/models"
	"github.com/stretchr/testify/assert"
)

func TestGetFilteredCandidates(t *testing.T) {
	service := &CandidateService{
		candidates: []models.Candidate{
			{
				ID:       "1",
				Name:     "John Doe",
				Email:    "john@example.com",
				Location: "New York",
				Skills:   []string{"Go", "TypeScript"},
				Education: models.Education{
					Degrees: []models.Degree{
						{School: "MIT", Subject: "CS", Degree: "BS"},
					},
				},
				WorkExperiences: []models.WorkExperience{
					{Company: "Google", RoleName: "Developer"},
				},
			},
		},
	}

	tests := []struct {
		name          string
		query         string
		expectedCount int
	}{
		{"Empty query", "", 1},
		{"Match name", "john", 1},
		{"Match skill", "typescript", 1},
		{"Match location", "york", 1},
		{"Match school", "mit", 1},
		{"Match company", "google", 1},
		{"No match", "invalid", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.GetFilteredCandidates(tt.query)
			assert.Equal(t, tt.expectedCount, len(result))
		})
	}
}

func TestPaginateCandidates(t *testing.T) {
	service := &CandidateService{}
	candidates := []models.Candidate{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
		{ID: "4"},
		{ID: "5"},
	}

	tests := []struct {
		name        string
		page        int
		pageCount   int
		expectedLen int
		hasNextPage bool
		expectError bool
	}{
		{"Valid page", 1, 2, 2, true, false},
		{"Last page", 3, 2, 1, false, false},
		{"Invalid page", 0, 2, 0, false, true},
		{"Invalid pageCount", 1, 0, 0, false, true},
		{"Page out of bounds", 10, 2, 0, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.PaginateCandidates(candidates, tt.page, tt.pageCount)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedLen, len(result.Candidates))
			assert.Equal(t, tt.hasNextPage, result.HasNextPage)
		})
	}
}

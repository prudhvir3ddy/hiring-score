package services

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/prudhvir3ddy/hiring-score/models"
	"github.com/prudhvir3ddy/hiring-score/services/utils"
	"os"
	"strings"
)

type CandidateService struct {
	candidates []models.Candidate
}

func NewCandidateService() *CandidateService {
	return &CandidateService{}
}

func (s *CandidateService) LoadCandidates() error {
	data, err := os.ReadFile("./data/candidates.json")
	if err != nil {
		return err
	}

	var rawCandidates []models.Candidate
	if err := json.Unmarshal(data, &rawCandidates); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	emailMap := make(map[string]models.Candidate)
	for _, c := range rawCandidates {
		c.ID = uuid.New().String()
		c.Score = utils.CalculateScore(c)
		if existing, ok := emailMap[c.Email]; ok {
			if c.SubmittedAt.Time.After(existing.SubmittedAt.Time) {
				emailMap[c.Email] = c
			}
		} else {
			emailMap[c.Email] = c
		}
	}

	s.candidates = make([]models.Candidate, 0, len(emailMap))
	for _, c := range emailMap {
		s.candidates = append(s.candidates, c)
	}

	utils.SortCandidates(s.candidates)

	return nil
}

func (s *CandidateService) GetFilteredCandidates(query string) []models.Candidate {
	if query == "" {
		return s.candidates
	}

	query = strings.ToLower(query)
	filtered := make(map[string]models.Candidate)

	for _, candidate := range s.candidates {
		// Skip if already added
		if _, exists := filtered[candidate.ID]; exists {
			continue
		}

		// Check name, email, location
		if strings.Contains(strings.ToLower(candidate.Name), query) ||
			strings.Contains(strings.ToLower(candidate.Email), query) ||
			strings.Contains(strings.ToLower(candidate.Location), query) {
			filtered[candidate.ID] = candidate
			continue
		}

		// Check skills
		for _, skill := range candidate.Skills {
			if strings.Contains(strings.ToLower(skill), query) {
				filtered[candidate.ID] = candidate
				break
			}
		}

		// Skip if already added
		if _, exists := filtered[candidate.ID]; exists {
			continue
		}

		// Check education
		for _, degree := range candidate.Education.Degrees {
			if strings.Contains(strings.ToLower(degree.School), query) ||
				strings.Contains(strings.ToLower(degree.Subject), query) ||
				strings.Contains(strings.ToLower(degree.Degree), query) {
				filtered[candidate.ID] = candidate
				break
			}
		}

		// Skip if already added
		if _, exists := filtered[candidate.ID]; exists {
			continue
		}

		// Check work experience
		for _, work := range candidate.WorkExperiences {
			if strings.Contains(strings.ToLower(work.Company), query) ||
				strings.Contains(strings.ToLower(work.RoleName), query) {
				filtered[candidate.ID] = candidate
				break
			}
		}
	}

	// Convert map to slice
	result := make([]models.Candidate, 0, len(filtered))
	for _, candidate := range filtered {
		result = append(result, candidate)
	}

	// Sort by score
	utils.SortCandidates(result)

	return result
}

func (s *CandidateService) PaginateCandidates(candidates []models.Candidate, page, pageCount int) (*models.PaginatedResponse, error) {
	// Check for negative values
	if page <= 0 || pageCount <= 0 {
		return nil, fmt.Errorf("invalid pagination parameters: page and pageCount must be positive")
	}

	// Handle empty candidates
	if len(candidates) == 0 {
		return &models.PaginatedResponse{
			Candidates:  []models.Candidate{},
			HasNextPage: false,
		}, nil
	}

	start := (page - 1) * pageCount
	end := start + pageCount

	// Check if start index is out of bounds
	if start >= len(candidates) {
		return &models.PaginatedResponse{
			Candidates:  []models.Candidate{},
			HasNextPage: false,
		}, nil
	}

	// Adjust end index if needed
	if end > len(candidates) {
		end = len(candidates)
	}

	return &models.PaginatedResponse{
		Candidates:  candidates[start:end],
		HasNextPage: end < len(candidates),
	}, nil
}

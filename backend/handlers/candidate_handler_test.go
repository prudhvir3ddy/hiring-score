// backend/handlers/candidate_handler_test.go
package handlers

import (
	"encoding/json"
	"golang.org/x/crypto/openpgp/errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prudhvir3ddy/hiring-score/models"
	"github.com/stretchr/testify/assert"
)

// FakeCandidateService implements services.CandidateService interface
type FakeCandidateService struct {
	candidates []models.Candidate
}

func (f *FakeCandidateService) GetFilteredCandidates(query string) []models.Candidate {
	if query == "" {
		return f.candidates
	}
	return nil
}

func (f *FakeCandidateService) PaginateCandidates(candidates []models.Candidate, page, pageCount int) (*models.PaginatedResponse, error) {
	if page <= 0 || pageCount <= 0 {
		return nil, errors.ErrKeyIncorrect
	}

	return &models.PaginatedResponse{
		Candidates:  candidates,
		HasNextPage: false,
	}, nil
}

// Verify FakeCandidateService implements the interface
var _ CandidateServiceInterface = (*FakeCandidateService)(nil)

func TestGetCandidates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCandidates := []models.Candidate{
		{
			ID:       "1",
			Name:     "Test User",
			Email:    "test@example.com",
			Phone:    "+1234567890",
			Location: "Test City",
			SubmittedAt: models.CustomTime{Time: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2025-01-28T17:29:05Z")
				return t
			}()},
			WorkAvailability:        []string{"Full-time"},
			AnnualSalaryExpectation: models.SalaryExpectation{FullTime: "100000"},
			WorkExperiences:         []models.WorkExperience{{Company: "Test Company", RoleName: "Software Engineer"}},
			Education: models.Education{
				HighestLevel: "Bachelor's",
				Degrees:      []models.Degree{{Degree: "Bachelor of Science", Subject: "Computer Science", School: "Test University", GPA: "3.8", StartDate: "2018-09-01", EndDate: "2022-05-31", OriginalSchool: "Test University", IsTop50: true, IsTop25: false}}},
			Skills: []string{"Go", "React", "TypeScript"},
			Score:  85.5,
		},
	}

	tests := []struct {
		name          string
		query         string
		page          string
		pageCount     string
		expectedCode  int
		expectedError bool
	}{
		{
			name:          "Valid request",
			query:         "",
			page:          "1",
			pageCount:     "10",
			expectedCode:  200,
			expectedError: false,
		},
		{
			name:          "Invalid page",
			query:         "",
			page:          "invalid",
			pageCount:     "10",
			expectedCode:  400,
			expectedError: true,
		},
		{
			name:          "Invalid pageCount",
			query:         "",
			page:          "1",
			pageCount:     "invalid",
			expectedCode:  400,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeService := &FakeCandidateService{
				candidates: testCandidates,
			}
			handler := NewCandidateHandler(fakeService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("GET", "/api/candidates", nil)
			q := req.URL.Query()
			q.Add("query", tt.query)
			q.Add("page", tt.page)
			q.Add("page_count", tt.pageCount)
			req.URL.RawQuery = q.Encode()

			c.Request = req

			handler.GetCandidates(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			if !tt.expectedError {
				var response models.PaginatedResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(testCandidates), len(response.Candidates))
			}
		})
	}
}

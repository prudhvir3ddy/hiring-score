package models

import (
	"fmt"
	"strings"
	"time"
)

type Candidate struct {
	ID                      string            `json:"id"`
	Name                    string            `json:"name"`
	Email                   string            `json:"email"`
	Phone                   string            `json:"phone"`
	Location                string            `json:"location"`
	SubmittedAt             CustomTime        `json:"submitted_at"`
	WorkAvailability        []string          `json:"work_availability"`
	AnnualSalaryExpectation SalaryExpectation `json:"annual_salary_expectation"`
	WorkExperiences         []WorkExperience  `json:"work_experiences"`
	Education               Education         `json:"education"`
	Skills                  []string          `json:"skills"`
	Score                   float64           `json:"score"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	layouts := []string{
		time.RFC3339, // "2025-01-28T17:29:05Z"
		"2006-01-02 15:04:05.000000",
		"2006-01-02",
	}

	var err error
	for _, layout := range layouts {
		t, parseErr := time.Parse(layout, s)
		if parseErr == nil {
			ct.Time = t
			return nil
		}
		err = parseErr
	}
	return err
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(time.RFC3339))), nil
}

type SalaryExpectation struct {
	FullTime string `json:"full-time"`
}

type WorkExperience struct {
	Company  string `json:"company"`
	RoleName string `json:"roleName"`
}

type Education struct {
	HighestLevel string   `json:"highest_level"`
	Degrees      []Degree `json:"degrees"`
}

type Degree struct {
	Degree         string `json:"degree"`
	Subject        string `json:"subject"`
	School         string `json:"school"`
	GPA            string `json:"gpa"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	OriginalSchool string `json:"originalSchool"`
	IsTop50        bool   `json:"isTop50"`
	IsTop25        bool   `json:"isTop25"`
}

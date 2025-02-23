package utils

import (
	"fmt"
	"github.com/prudhvir3ddy/hiring-score/models"
	"strings"
)

var highPrioritySkills = map[string]bool{
	"Golang":                    true,
	"Kotlin":                    true,
	"Python":                    true,
	"JavaScript":                true,
	"TypeScript":                true,
	"Java":                      true,
	"React":                     true,
	"Next.js":                   true,
	"TailwindCSS":               true,
	"Django":                    true,
	"Redis":                     true,
	"Kafka":                     true,
	"Docker":                    true,
	"Amazon Web Services (AWS)": true,
	"MongoDB":                   true,
	"PostgreSQL":                true,
}

const (
	maxSkillPoints        = 40
	maxExperiencePoints   = 30
	maxEducationPoints    = 15
	maxAvailabilityPoints = 5
)

func CalculateScore(c models.Candidate) float64 {
	skillScore := float64(len(c.Skills) * 2)
	for _, skill := range c.Skills {
		if highPrioritySkills[skill] {
			skillScore += 3
		}
	}
	if skillScore > maxSkillPoints {
		skillScore = maxSkillPoints
	}

	workExpScore := float64(len(c.WorkExperiences) * 3)
	if workExpScore > maxExperiencePoints {
		workExpScore = maxExperiencePoints
	}

	educationScore := 0.0
	if c.Education.HighestLevel == "Master's Degree" {
		educationScore = 10
	} else if c.Education.HighestLevel == "Bachelor's Degree" {
		educationScore = 5
	}
	for _, d := range c.Education.Degrees {
		if d.IsTop25 {
			educationScore += 5
			break
		} else if d.IsTop50 {
			educationScore += 3
			break
		}
	}
	if educationScore > maxEducationPoints {
		educationScore = maxEducationPoints
	}

	salaryScore := 0.0
	salary := strings.TrimPrefix(c.AnnualSalaryExpectation.FullTime, "$")
	salaryFloat := 0.0
	fmt.Sscanf(salary, "%f", &salaryFloat)
	if salaryFloat <= 70000 {
		salaryScore += 10
	} else if salaryFloat <= 90000 {
		salaryScore += 8
	} else if salaryFloat <= 110000 {
		salaryScore += 5
	} else {
		salaryScore += 3
	}

	availabilityScore := 0.0
	for _, wa := range c.WorkAvailability {
		if wa == "full-time" {
			availabilityScore = maxAvailabilityPoints
		}
	}

	return workExpScore + skillScore + educationScore + availabilityScore + salaryScore
}

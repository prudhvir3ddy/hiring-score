package utils

import (
	"github.com/prudhvir3ddy/hiring-score/models"
	"sort"
)

func SortCandidates(candidates []models.Candidate) {
	sort.Slice(candidates, func(i, j int) bool {
		// First compare by score
		if candidates[i].Score != candidates[j].Score {
			return candidates[i].Score > candidates[j].Score
		}

		// If scores are equal, compare by work experience count
		iWorkExp := len(candidates[i].WorkExperiences)
		jWorkExp := len(candidates[j].WorkExperiences)
		if iWorkExp != jWorkExp {
			return iWorkExp > jWorkExp
		}

		// If work experience is equal, compare by skills count
		iSkills := len(candidates[i].Skills)
		jSkills := len(candidates[j].Skills)
		if iSkills != jSkills {
			return iSkills > jSkills
		}

		// If skills are equal, compare by education
		iEducation := getEducationScore(candidates[i].Education)
		jEducation := getEducationScore(candidates[j].Education)
		return iEducation > jEducation
	})
}

func getEducationScore(education models.Education) int {
	score := 0
	for _, degree := range education.Degrees {
		if degree.IsTop25 {
			score += 2
		} else if degree.IsTop50 {
			score += 1
		}
	}
	return score
}

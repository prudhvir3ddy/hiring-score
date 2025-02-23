package models

type PaginatedResponse struct {
	Candidates  []Candidate `json:"candidates"`
	HasNextPage bool        `json:"hasNextPage"`
}

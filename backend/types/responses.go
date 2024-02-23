package types

//
// This file will contain types that will define the response of request handlers
// This also represents what the front end expects in `data` field of it's response
//

type SolutionEntity struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Photos      []string `json:"photos"`
	Ingredients []string `json:"ingredients"`
	Description string   `json:"description"`
}

type IdentifiedDiseaseEntity struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type UploadedEntity struct {
	Id         string                   `json:"id"`
	Name       string                   `json:"name"`
	Images     []string                 `json:"images"`
	Severity   int                      `json:"severity"`
	Status     string                   `json:"status"`
	Solutions  []*SolutionEntity        `json:"solutions"`
	Identified *IdentifiedDiseaseEntity `json:"identified"`
}

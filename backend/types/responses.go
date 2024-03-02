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
type CausesEntity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type IdentifiedDiseaseEntity struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Solutions []*SolutionEntity `json:"solutions"`
	Causes    []*CausesEntity   `json:"causes"`
}

type UploadedEntity struct {
	Id         string                     `json:"id"`
	Name       []string                   `json:"name"`
	Severity   int                        `json:"severity"`
	Status     string                     `json:"status"`
	Images     []string                   `json:"images"`
	Roi        string                     `json:"roi"`
	Identified []*IdentifiedDiseaseEntity `json:"identified"`
}

type UploadListItemType struct {
	Id       string   `json:"id"`
	Name     []string `json:"name"`
	Severity int      `json:"severity"`
	Status   string   `json:"status"`
	Images   []string `json:"images"`
}

type HeatMapEntry struct {
	Id        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Weight    int     `json:"weight"`
}

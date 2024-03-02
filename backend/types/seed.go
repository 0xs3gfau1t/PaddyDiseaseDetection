package types

type SeedJson struct {
	Disease []DiseaseJson `json:"diseases"`
}
type DiseaseJson struct {
	Name      string         `json:"name"`
	Solutions []SolutionJson `json:"solutions"`
	Causes    []CauseJson    `json:"causes"`
}
type SolutionJson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Ingredients []string `json:"ingredients"`
}

type CauseJson struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

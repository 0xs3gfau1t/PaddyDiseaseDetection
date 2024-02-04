package types

type SeedJson struct {
	Disease []DiseaseJson `json:"diseases"`
}
type DiseaseJson struct {
	Name      string         `json:"name"`
	Solutions []SolutionJson `json:"solutions"`
}
type SolutionJson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Ingredients []string `json:"ingredients"`
}

package types

//
// This file will contain types that will define the response of request handlers
// This also represents what the front end expects in `data` field of it's response
//

type UploadedEntity struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Images   string `json:"images"`
	Severity int    `json:"severity"`
	Status   string `json:"status"`
}

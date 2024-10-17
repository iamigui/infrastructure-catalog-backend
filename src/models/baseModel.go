package models

// * Model Project
type Project struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	JSONData    map[string]interface{} `json:"jsonData"`
}

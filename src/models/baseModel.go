package models

// Project data model
type Project struct {
	ID          string                 `json:"id" bson:"id"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	JSONData    map[string]interface{} `json:"jsonData" bson:"jsondata"`
}

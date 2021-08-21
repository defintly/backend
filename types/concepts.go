package types

import "reflect"

type Concept struct {
	Id           int    `json:"id" db:"id"`
	Icon         string `json:"icon" db:"icon"`
	Type         string `json:"type" db:"type"`
	Concept      string `json:"concept" db:"concept"`
	Definition   string `json:"definition" db:"definition"`
	Author       string `json:"author" db:"author"`
	Source       string `json:"source" db:"source"`
	CollectionId int    `json:"collection_id" db:"collection_id"`
}

var ConceptType = reflect.TypeOf(&Concept{})

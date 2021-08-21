package types

import "reflect"

type Concept struct {
	Id         int    `json:"id" db:"id"`
	Icon       string `json:"icon" db:"icon"`
	Concept    string `json:"concept" db:"concept"`
	Definition string `json:"definition" db:"definition"`
	Author     string `json:"author" db:"author"`
	Source     string `json:"source" db:"source"`
}

var ConceptType = reflect.TypeOf(&Concept{})

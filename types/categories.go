package types

import "reflect"

type Category struct {
	Id          int    `json:"id" db:"id"`
	Icon        string `json:"icon" db:"icon"`
	Category    string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
	Type        string `json:"type" db:"type"`
}

var CategoryType = reflect.TypeOf(&Category{})

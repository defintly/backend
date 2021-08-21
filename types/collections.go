package types

import "reflect"

type Collection struct {
	Id         int    `json:"id" db:"id"`
	Icon       string `json:"icon" db:"icon"`
	Collection string `json:"collection" db:"collection"`
	Type       string `json:"type" db:"type"`
}

var CollectionType = reflect.TypeOf(&Collection{})

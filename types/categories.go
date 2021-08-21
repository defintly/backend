package types

import "reflect"

type Category struct {
	Id                            int    `json:"id" db:"id"`
	Icon                          string `json:"icon" db:"icon"`
	Category                      string `json:"category" db:"category"`
	Type                          string `json:"type" db:"type"`
	QualityCriteriaInThisCategory string `json:"quality_criteria_in_this_category" db:"quality_criteria_in_this_category"`
}

var CategoryType = reflect.TypeOf(&Category{})

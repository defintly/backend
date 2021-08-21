package types

import "reflect"

type Criteria struct {
	Id                                    int    `json:"id" db:"id"`
	Icon                                  string `json:"icon" db:"icon"`
	QualityCriterion                      string `json:"quality_criterion" db:"quality_criterion"`
	DescriptionShort                      string `json:"description_short" db:"description_short"`
	DescriptionLong                       string `json:"description_long" db:"description_long"`
	Example                               string `json:"example" db:"example"`
	Explanation                           string `json:"explanation" db:"explanation"`
	Type                                  string `json:"type" db:"type"`
	CategoryId                            int    `json:"category_id" db:"category_id"`
	OtherQualityCriteriaInTheSameCategory string `json:"other_quality_criteria_in_the_same_category" db:"other_quality_criteria_in_the_same_category"`
	References                            string `json:"references" db:"references"`
}

var CriteriaType = reflect.TypeOf(&Criteria{})

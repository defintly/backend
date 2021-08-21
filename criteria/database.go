package criteria

import (
	"errors"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

var NotFound = errors.New("criteria not found")

func GetAllCriteria() ([]*types.Criteria, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CriteriaType, "SELECT * FROM criteria")
	if err != nil {
		return nil, err
	}

	return slice.([]*types.Criteria), err
}

func GetCriteriaById(id int) (*types.Criteria, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CriteriaType,
		"SELECT * FROM criteria WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	criteria := slice.([]*types.Criteria)
	if len(criteria) == 0 {
		return nil, NotFound
	}

	return criteria[0], err
}

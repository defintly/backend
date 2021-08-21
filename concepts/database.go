package concepts

import (
	"errors"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

var NotFound = errors.New("concept not found")

func GetAllConcepts() ([]*types.Concept, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.ConceptType, "SELECT * FROM concepts")
	if err != nil {
		return nil, err
	}

	return slice.([]*types.Concept), err
}

func GetConceptById(id int) (*types.Concept, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.ConceptType,
		"SELECT * FROM concepts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	concepts := slice.([]*types.Concept)
	if len(concepts) == 0 {
		return nil, NotFound
	}

	return concepts[0], err
}

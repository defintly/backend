package collections

import (
	"errors"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

var NotFound = errors.New("collection not found")

func GetAllCollections() ([]*types.Collection, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CollectionType, "SELECT * FROM collections")
	if err != nil {
		return nil, err
	}

	return slice.([]*types.Collection), err
}

func GetCollectionById(id int) (*types.Collection, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CollectionType,
		"SELECT * FROM collections WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	collections := slice.([]*types.Collection)
	if len(collections) == 0 {
		return nil, NotFound
	}

	return collections[0], err
}

func Search(search string) ([]*types.Collection, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CollectionType,
		"SELECT * FROM collections WHERE collection LIKE $1", search+"%")

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Collection), nil
}

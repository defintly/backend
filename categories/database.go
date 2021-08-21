package categories

import (
	"errors"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

var NotFound = errors.New("category not found")

func GetAllCategories() ([]*types.Category, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CategoryType, "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	return slice.([]*types.Category), err
}

func GetCategoryById(id int) (*types.Category, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CategoryType,
		"SELECT * FROM categories WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	categories := slice.([]*types.Category)
	if len(categories) == 0 {
		return nil, NotFound
	}

	return categories[0], err
}

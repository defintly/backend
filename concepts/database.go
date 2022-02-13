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

func Search(search string) ([]*types.Concept, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CriteriaType,
		"SELECT * FROM concepts WHERE concept LIKE $1", search+"%")

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Concept), nil
}

func AddComment(conceptId int, userId int, text string, parentCommentId *int) (int, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"INSERT INTO concept_comments(concept_id, user_id, text, parent_id, reviewed) VALUES($1, $2, $3, $4, $5) RETURNING id",
		conceptId, userId, text, parentCommentId, false)

	if err != nil {
		return -1, err
	}

	return slice.([]*types.IdInformation)[0].Id, nil
}

func DeleteComment(commentId int) error {
	return database.PrepareAsync(database.DefaultTimeout, "DELETE FROM concept_comments WHERE id = $1", commentId)
}

func AllowComment(commentId int) error {
	return database.PrepareAsync(database.DefaultTimeout, "UPDATE concept_comments SET allowed = 1 WHERE id = $1",
		commentId)
}

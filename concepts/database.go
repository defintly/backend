package concepts

import (
	"errors"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

var (
	NotFound        = errors.New("concept not found")
	CommentNotFound = errors.New("comment not found")
)

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
	slice, err := database.QueryAsync(database.DefaultTimeout, types.ConceptType,
		"SELECT * FROM concepts WHERE concept LIKE $1", search+"%")

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Concept), nil
}

func AddComment(conceptId int, userId int, text string, parentCommentId *int) (int, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"INSERT INTO concept_comments(concept_id, user_id, text, parent_id, allowed) VALUES($1, $2, $3, $4, $5) "+
			"RETURNING id", conceptId, userId, text, parentCommentId, false)

	if err != nil {
		return -1, err
	}

	return slice.([]*types.IdInformation)[0].Id, nil
}

func GetCreatorUserIdOfComment(commentId int) (int, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"SELECT user_id AS id FROM (SELECT user_id FROM concept_comments WHERE id = $1) users", commentId)

	if err != nil {
		return -1, err
	}

	userIds := slice.([]*types.IdInformation)
	if len(userIds) == 0 {
		return -1, CommentNotFound
	}

	return userIds[0].Id, nil
}

func ListUnreviewedComments() ([]*types.Comment, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CommentType,
		"SELECT * FROM concept_comments WHERE allowed = false")

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Comment), nil
}

func DeleteComment(commentId int) error {
	return database.PrepareAsync(database.DefaultTimeout, "DELETE FROM concept_comments WHERE id = $1", commentId)
}

func AllowComment(commentId int) error {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"UPDATE concept_comments SET allowed = true WHERE id = $1 RETURNING id", commentId)

	if err != nil {
		return err
	}

	if len(slice.([]*types.IdInformation)) == 0 {
		return CommentNotFound
	}

	return nil
}

func GetComment(commentId int) (*types.Comment, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CommentType,
		"SELECT * FROM concept_comments WHERE id = $1", commentId)

	if err != nil {
		return nil, err
	}

	comments := slice.([]*types.Comment)

	if len(comments) == 0 {
		return nil, CommentNotFound
	}

	return comments[0], nil
}

func GetParentComments(id int) ([]*types.Comment, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.CommentType,
		"SELECT * FROM concept_comments WHERE concept_id = $1 AND parent_id IS NULL", id)

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Comment), nil
}

// GetCommentTree only allows a depth of one child - more children can be received by
// querying each child (and child of a child, ...)
func GetCommentTree(commentId int) (*types.CommentTree, error) {
	comment, err := GetComment(commentId)
	if err != nil {
		return nil, err
	}

	tree := &types.CommentTree{
		Comment:  comment,
		Children: nil,
	}

	slice, err := database.QueryAsync(database.DefaultTimeout, types.CommentType,
		"SELECT * FROM concept_comments WHERE parent_id = $1", comment.Id)

	if err != nil {
		return nil, err
	}

	tree.Children = slice.([]*types.Comment)
	return tree, nil
}

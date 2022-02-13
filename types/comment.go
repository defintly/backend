package types

import "reflect"

var (
	CommentType = reflect.TypeOf(&Comment{})
)

type Comment struct {
	Id              *int   `json:"id,omitempty" db:"id"`
	ConceptId       *int   `json:"concept_id,omitempty" db:"concept_id"`
	UserId          *int   `json:"user_id,omitempty" db:"user_id"`
	Text            string `json:"text" db:"text"`
	ParentCommentId *int   `json:"parent_comment_id,omitempty" db:"parent_comment_id"`
	Allowed         *bool  `json:"allowed,omitempty" db:"allowed"`
}

type CommentTree struct {
	Comment  *Comment   `json:"comment"`
	Children []*Comment `json:"children"`
}

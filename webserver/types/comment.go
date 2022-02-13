package types

type Comment struct {
	Id              *int   `json:"id,omitempty"`
	Text            string `json:"text"`
	ParentCommentId *int   `json:"parent_comment_id,omitempty"`
}

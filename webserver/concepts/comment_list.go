package concepts

import (
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/permissions"
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUnreviewedComments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)

		canAllowComments, err := users.HasPermission(user.Id, permissions.AllowComment)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to check user permission: ", err)
			return
		}

		if !canAllowComments {
			ctx.AbortWithStatusJSON(http.StatusForbidden, errors.NoPermission)
			return
		}

		commentList, err := concepts.ListUnreviewedComments()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to list unreviewed comments: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"comments": commentList})
	}
}

// GetCommentList only allows a depth of one for children - more children can be received by
// querying each child (and child of a child, ...)
func GetCommentList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentId := ctx.MustGet("id").(int)

		tree, err := concepts.GetCommentTree(commentId)
		if err != nil {
			if err == concepts.CommentNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CommentNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to list comment tree: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, tree)
	}
}

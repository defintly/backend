package concepts

import (
	"encoding/json"
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/permissions"
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetComments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id").(int)

		comments, err := concepts.GetParentComments(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get comments: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"comments": comments})
	}
}

func AddComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)
		id := ctx.MustGet("id").(int)
		defer ctx.Request.Body.Close()

		comment := &types.Comment{}

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		err = json.Unmarshal(body, comment)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		commentId, err := concepts.AddComment(id, user.Id, comment.Text, comment.ParentId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to add comment: ", err)
			return
		}

		allowCommentInternally(ctx, user.Id, commentId, false)

		comment.Id = &commentId

		ctx.JSON(http.StatusOK, comment)
	}
}

func DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)
		commentId := ctx.MustGet("id").(int)

		creatorId, err := concepts.GetCreatorUserIdOfComment(commentId)
		if err != nil {
			if err == concepts.CommentNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CommentNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get comment creator user id: ", err)
			}
			return
		}

		if user.Id == creatorId {
			deleteCommentInternally(ctx, commentId)
			return
		}

		canDeleteComments, err := users.HasPermission(user.Id, permissions.DeleteComment)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to check user permission: ", err)
			return
		}

		if canDeleteComments {
			deleteCommentInternally(ctx, commentId)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, errors.NoPermission)
	}
}

func AllowComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)
		commentId := ctx.MustGet("id").(int)

		allowCommentInternally(ctx, user.Id, commentId, true)
	}
}

func allowCommentInternally(ctx *gin.Context, userId int, commentId int, finishRequest bool) {
	canAllowComments, err := users.HasPermission(userId, permissions.AllowComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
		general.Log.Error("Failed to check user permission: ", err)
		return
	}

	if !canAllowComments {
		ctx.AbortWithStatusJSON(http.StatusForbidden, errors.NoPermission)
		return
	}

	if err := concepts.AllowComment(commentId); err != nil {
		if err == concepts.CommentNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CommentNotFound)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to allow comment: ", err)
		}
		return
	}

	if finishRequest {
		ctx.JSON(http.StatusOK, gin.H{"id": commentId, "status": "allowed"})
	}
}

func deleteCommentInternally(ctx *gin.Context, commentId int) {
	err := concepts.DeleteComment(commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
		general.Log.Error("Failed to delete command: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": commentId, "status": "deleted"})
}

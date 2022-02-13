package concepts

import (
	"encoding/json"
	"github.com/defintly/backend/concepts"
	generalTypes "github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/defintly/backend/webserver/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func AddComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*generalTypes.User)
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

		commentId, err := concepts.AddComment(id, user.Id, comment.Text, comment.ParentCommentId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			return
		}

		canAllowComments, err := users.HasPermission(user.Id, "comment.allow")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			return
		}

		if canAllowComments {
			if concepts.AllowComment(commentId) != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				return
			}
		}

		comment.Id = &commentId

		ctx.JSON(http.StatusOK, comment)
	}
}

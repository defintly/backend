package users

import (
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)
		id := ctx.MustGet("userId").(int)

		if id == user.Id {
			ctx.JSON(http.StatusOK, user)
			return
		}

		requestedUser, err := users.GetUserById(id)

		if err != nil {
			if err == users.UserNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.NotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get user: ", err)
			}
			return
		}

		requestedUser.MailAddress = nil
		ctx.JSON(http.StatusOK, requestedUser)
	}
}

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
		id := ctx.MustGet("userId").(int)
		if id == -1 {
			return
		}
		showMail := false

		user, ok := ctx.Get("user")
		if ok && user.(*types.User).Id == id {
			showMail = true
		}

		requestedUser, err := users.GetUserById(showMail, id)

		if err != nil {
			if err == users.UserNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.NotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get user: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, requestedUser)
	}
}

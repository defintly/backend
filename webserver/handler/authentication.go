package handler

import (
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const authKeyHeader = "X-Auth-Key"

func AuthenticationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authKey := ctx.Request.URL.Query().Get("token")

		if authKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.MissingAuthenticationInformation)
			return
		}

		user, err := users.GetUserByAuthenticationKey(authKey)
		if err != nil {
			if err == users.UserNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.MissingAuthenticationInformation)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)

				general.Log.Error("Failed to get authentication information of user: ", err)
			}
			return
		}

		ctx.Set("user", user)
	}
}

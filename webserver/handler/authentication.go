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
		AuthenticateInternally(ctx)
	}
}

func AuthenticateInternally(ctx *gin.Context) (continueExecution bool) {
	authKey := ctx.Request.Header.Get(authKeyHeader)

	if authKey == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.MissingAuthenticationInformation)
		return false
	}

	user, err := users.GetUserByAuthenticationKey(authKey)
	if err != nil {
		if err == users.UserNotFound {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.MissingAuthenticationInformation)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)

			general.Log.Error("Failed to get authentication information of user: ", err)
		}
		return false
	}

	ctx.Set("user", user)
	return true
}

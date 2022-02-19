package helper

import (
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/webserver/errors"
	"github.com/defintly/backend/webserver/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserIdWithMeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userId", extractUserIdWithMe(ctx))
	}
}

func ExtractUserId(ctx *gin.Context) int {
	userIdAsString := ctx.Param("userId")

	if userIdAsString == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
		return -1
	}

	userId, err := strconv.Atoi(userIdAsString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
		return -1
	}

	return userId
}

func extractUserIdWithMe(ctx *gin.Context) int {
	userIdAsString := ctx.Param("userId")

	if userIdAsString == "me" {
		if !handler.AuthenticateInternally(ctx) {
			return -1
		}
		user, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.MissingAuthenticationInformation)
			return -1
		}
		return user.(*types.User).Id
	}

	return ExtractUserId(ctx)
}

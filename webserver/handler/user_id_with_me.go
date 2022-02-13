package handler

import (
	"github.com/defintly/backend/webserver/helper"
	"github.com/gin-gonic/gin"
)

func UserIdWithMeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userId", helper.ExtractUserIdWithMe(ctx))
	}
}

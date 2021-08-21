package handler

import (
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Id() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idAsString := ctx.Param("id")

		if idAsString == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		id, err := strconv.Atoi(idAsString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		ctx.Set("id", id)
	}
}

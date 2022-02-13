package concepts

import (
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchString, ok := ctx.GetQuery("query")
		if !ok || len(searchString) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		foundConcepts, err := concepts.Search(searchString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to search for concepts: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"concepts": foundConcepts})
	}
}

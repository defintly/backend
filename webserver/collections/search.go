package collections

import (
	"github.com/defintly/backend/collections"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchString, ok := ctx.MustGet("query").(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		foundCollections, err := collections.Search(searchString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to search for collections: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"collections": foundCollections})
	}
}

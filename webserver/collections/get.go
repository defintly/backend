package collections

import (
	"github.com/defintly/backend/collections"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		collectionList, err := collections.GetAllCollections()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get collections: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"collections": collectionList})
	}
}

func GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.MustGet("id").(int)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		collection, err := collections.GetCollectionById(id)
		if err != nil {
			if err == collections.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CollectionNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get collections: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, collection)
	}
}

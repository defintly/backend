package concepts

import (
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conceptList, err := concepts.GetAllConcepts()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get concepts: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"concepts": conceptList})
	}
}

func GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.MustGet("id").(int)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		concept, err := concepts.GetConceptById(id)
		if err != nil {
			if err == concepts.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.ConceptNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get concepts: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, concept)
	}
}

package criteria

import (
	"github.com/defintly/backend/criteria"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		criteriaList, err := criteria.GetAllCriteria()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get criteria: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"criteria": criteriaList})
	}
}

func GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.MustGet("id").(int)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		criteriaObject, err := criteria.GetCriteriaById(id)
		if err != nil {
			if err == criteria.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CriteriaNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get criteria: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, criteriaObject)
	}
}

package categories

import (
	"github.com/defintly/backend/categories"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryList, err := categories.GetAllCategories()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get categories: ", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"categories": categoryList})
	}
}

func GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.MustGet("id").(int)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		category, err := categories.GetCategoryById(id)
		if err != nil {
			if err == categories.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.CategoryNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get categories: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, category)
	}
}

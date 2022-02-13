package criteria

import (
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/criteria"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"net/http"
)

func GenerateHTML() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id").(int)

		selectedCriteria, err := criteria.GetCriteriaById(id)
		if err != nil {
			if err == concepts.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.ConceptNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get selectedCriteria for export: ", err)
			}
			return
		}

		ctx.Data(http.StatusOK, "application/html", markdown.ToHTML([]byte(
			selectedCriteria.QualityCriterion+"\n"+selectedCriteria.DescriptionLong+"\n"+selectedCriteria.Example),
			nil, nil))
	}
}

package criteria

import (
	"github.com/defintly/backend/categories"
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/criteria"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"net/http"
	"strings"
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

		category, err := categories.GetCategoryById(selectedCriteria.CategoryId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
			general.Log.Error("Failed to get selectedCriteria's category for export: ", err)
			return
		}

		returnData := []byte("<html><head><title>" + selectedCriteria.QualityCriterion + "</title></head><body>")
		returnData = append(returnData,
			markdown.ToHTML([]byte(
				"# "+selectedCriteria.QualityCriterion+
					"\n\n"+
					"**Category:** "+category.Category+
					"\n\n"+
					strings.ReplaceAll(selectedCriteria.DescriptionLong, "\\n", "\\\n")+
					"\n\n"+
					strings.ReplaceAll(selectedCriteria.Example, "\\n", "\\\n")+
					"\n\n"+
					strings.ReplaceAll(selectedCriteria.References, "\\n", "\\\n"),
			), nil, nil)...)
		returnData = append(returnData, []byte("</body></html>")...)

		ctx.Data(http.StatusOK, "text/html; charset=UTF-8", returnData)
	}
}

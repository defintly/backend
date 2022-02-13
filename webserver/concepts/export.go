package concepts

import (
	"github.com/defintly/backend/concepts"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"net/http"
)

func GenerateHTML() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id").(int)

		concept, err := concepts.GetConceptById(id)
		if err != nil {
			if err == concepts.NotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, errors.ConceptNotFound)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to get concept for export: ", err)
			}
			return
		}

		ctx.Data(http.StatusOK, "application/html",
			markdown.ToHTML([]byte(concept.Definition+"\n"+concept.Source), nil, nil))
	}
}

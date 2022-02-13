package concepts

import (
	"github.com/defintly/backend/concepts"
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

		// TODO add "how to cite"?
		returnData := []byte("<html><head><title>" + concept.Concept + "</title></head><body>")
		returnData = append(returnData,
			markdown.ToHTML([]byte(
				"# "+concept.Concept+
					"\n\n"+
					strings.ReplaceAll(concept.Definition, "\\n", "\\\n")+
					"\n\n"+
					strings.ReplaceAll(concept.Source, "\\n", "\\\n"),
			), nil, nil)...)
		returnData = append(returnData, []byte("</body></html>")...)

		ctx.Data(http.StatusOK, "text/html; charset=UTF-8", returnData)

	}
}

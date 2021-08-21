package webserver

import (
	"fmt"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/categories"
	"github.com/defintly/backend/webserver/collections"
	"github.com/defintly/backend/webserver/concepts"
	"github.com/defintly/backend/webserver/criteria"
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"net/http"
)

func Run(hostname string, port int) {
	router := gin.New()
	router.Use(ginlogrus.Logger(general.Log), gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"name": "defintly API", "version": "2.0.0"})
	})

	initCategoryRoutes(router)
	initCollectionRoutes(router)
	initConceptRoutes(router)
	initCriteriaRoutes(router)

	panic(router.Run(fmt.Sprintf("%s:%d", hostname, port)))
}

func initCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/categories")

	categoryGroup.GET("", categories.GetAll())

	categoryIdGroup := categoryGroup.Group("/:id")
	categoryIdGroup.GET("", categories.GetById())
}

func initCollectionRoutes(router *gin.Engine) {
	collectionGroup := router.Group("/collections")

	collectionGroup.GET("", collections.GetAll())

	collectionIdGroup := collectionGroup.Group("/:id")
	collectionIdGroup.GET("", collections.GetById())
}

func initConceptRoutes(router *gin.Engine) {
	conceptGroup := router.Group("/concepts")

	conceptGroup.GET("", concepts.GetAll())

	conceptIdGroup := conceptGroup.Group("/:id")
	conceptIdGroup.GET("", concepts.GetById())
}

func initCriteriaRoutes(router *gin.Engine) {
	criteriaGroup := router.Group("/criteria")

	criteriaGroup.GET("", criteria.GetAll())

	criteriaIdGroup := criteriaGroup.Group("/:id")
	criteriaIdGroup.GET("", criteria.GetById())
}

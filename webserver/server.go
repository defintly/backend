package webserver

import (
	"fmt"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/webserver/authentication"
	"github.com/defintly/backend/webserver/categories"
	"github.com/defintly/backend/webserver/collections"
	"github.com/defintly/backend/webserver/concepts"
	"github.com/defintly/backend/webserver/criteria"
	"github.com/defintly/backend/webserver/handler"
	"github.com/defintly/backend/webserver/users"
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"net/http"
)

func Run(hostname string, port int) {
	router := gin.New()
	router.Use(ginlogrus.Logger(general.Log), gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"name": "defintly API", "version": "2.1.1"})
	})

	authHandler := handler.AuthenticationHandler()

	initCategoryRoutes(router)
	initCollectionRoutes(router)
	initConceptRoutes(router, authHandler)
	initCriteriaRoutes(router)
	initAuthRoutes(router, authHandler)
	initUserRoutes(router, authHandler)

	panic(router.Run(fmt.Sprintf("%s:%d", hostname, port)))
}

func initCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/categories")

	categoryGroup.GET("", categories.GetAll())
	categoryGroup.GET("/search", categories.Search())

	categoryIdGroup := categoryGroup.Group("/:id", handler.Id())
	categoryIdGroup.GET("", categories.GetById())
}

func initCollectionRoutes(router *gin.Engine) {
	collectionGroup := router.Group("/collections")

	collectionGroup.GET("", collections.GetAll())
	collectionGroup.GET("/search", collections.Search())

	collectionIdGroup := collectionGroup.Group("/:id", handler.Id())
	collectionIdGroup.GET("", collections.GetById())
}

func initConceptRoutes(router *gin.Engine, authHandler gin.HandlerFunc) {
	conceptGroup := router.Group("/concepts")

	conceptGroup.GET("", concepts.GetAll())
	conceptGroup.GET("/search", concepts.Search())

	conceptIdGroup := conceptGroup.Group("/:id", handler.Id())
	conceptIdGroup.GET("", concepts.GetById())
	conceptIdGroup.GET("/export", concepts.GenerateHTML())
	conceptIdGroup.GET("/comments", concepts.GetComments())
	conceptIdGroup.POST("/comments", authHandler, concepts.AddComment())

	conceptCommentGroup := conceptGroup.Group("/comments")
	conceptCommentGroup.GET("/list-unreviewed", authHandler, concepts.ListUnreviewedComments())

	conceptCommentIdGroup := conceptCommentGroup.Group("/:id", handler.Id())
	conceptCommentIdGroup.DELETE("/", authHandler, concepts.DeleteComment())
	conceptCommentIdGroup.PUT("/allow", authHandler, concepts.AllowComment())
	conceptCommentIdGroup.GET("/children", concepts.GetCommentList())
}

func initCriteriaRoutes(router *gin.Engine) {
	criteriaGroup := router.Group("/criteria")

	criteriaGroup.GET("", criteria.GetAll())
	criteriaGroup.GET("/search", criteria.Search())

	criteriaIdGroup := criteriaGroup.Group("/:id", handler.Id())
	criteriaIdGroup.GET("", criteria.GetById())
	criteriaIdGroup.GET("/export", criteria.GenerateHTML())
}

func initAuthRoutes(router *gin.Engine, authHandler gin.HandlerFunc) {
	authGroup := router.Group("/auth")

	authGroup.POST("/login", authentication.Login())
	authGroup.POST("/register", authentication.Register())
	authGroup.PUT("/mail", authHandler, authentication.ChangeMailAddress())
	authGroup.PUT("/password", authHandler, authentication.ChangePassword())
}

func initUserRoutes(router *gin.Engine, authHandler gin.HandlerFunc) {
	usersGroup := router.Group("/users", authHandler)

	usersIdGroup := usersGroup.Group("/:userId", handler.UserIdWithMeHandler())

	usersIdGroup.GET("", users.Get())
	usersIdGroup.PUT("/username", users.ChangeUsername())
	usersIdGroup.PUT("/firstname", users.ChangeFirstname())
	usersIdGroup.PUT("/lastname", users.ChangeLastname())
}

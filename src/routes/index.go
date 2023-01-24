package routes

import (
	"library-api-rest/src/controllers"

	"github.com/gin-gonic/gin"
	docs "library-api-rest/src/docs"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.Use(cors.Default())
	r.GET("/", controllers.Main)
	r.GET("/books", controllers.GetAll)
	r.GET("/books/:id", controllers.GetOne)
	r.POST("/books", controllers.Create)
	r.PUT("/books/:id", controllers.Update)
	r.DELETE("/books/:id", controllers.Delete)
	r.GET("/books/search/:title", controllers.FindByTitle)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}

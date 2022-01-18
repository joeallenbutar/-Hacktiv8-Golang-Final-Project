package routers

import (
	"Final-Project/database"
	"Final-Project/controllers"
	"Final-Project/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

)

func Setup() *gin.Engine {
	docs.SwaggerInfo.Title = "Example ToDo's Rest API"
	docs.SwaggerInfo.Description = "Documentation of Todo's Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.Default()
	api := &controllers.APIEnv{
		DB: database.GetDB(),
	}

	re := r.Group("/todos")
	{
		re.GET("", api.GetToDos)
		re.POST("", api.CreateToDo)
		re.GET("/:id", api.GetToDo)
		re.PUT("/:id", api.UpdateToDo)
		re.DELETE("/:id", api.DeleteToDo)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
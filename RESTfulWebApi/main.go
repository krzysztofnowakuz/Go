package main

import (
	"github.com/gin-gonic/gin"

	controllers "RestApi/Controllers"
	"RestApi/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	// Get/Api/Items
	r.GET("/Api/Items", controllers.GetItems)
	// Get/Api/Item/:id
	r.GET("/Api/Item/:id", controllers.GetItemById)
	// Post/Api/Item/
	r.POST("/Api/Item/", controllers.PostItem)
	// Put/Api/Item/:id
	r.PUT("/Api/Item/:id", controllers.PutItem)
	// Delete/Api/Item/:id
	r.DELETE("/Api/Item/:id", controllers.DeleteItem)

	r.Run()
}

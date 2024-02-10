package main

import (
	productcontrollers "github.com/fanes-setiawan/go-restapi-gin/controllers/product_controllers"
	"github.com/fanes-setiawan/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/api/product", productcontrollers.Index)
	r.GET("/api/product/:id", productcontrollers.Show)
	r.POST("/api/product", productcontrollers.Create)
	r.PUT("/api/product/:id", productcontrollers.Update)
	r.DELETE("/api/product", productcontrollers.Delete)

	r.Run()

}

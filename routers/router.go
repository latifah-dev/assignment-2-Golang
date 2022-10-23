package routers

import (
	"Tugas2/controllers"

	"github.com/gin-gonic/gin"
)

func RootHandler() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controllers.GetOrder)
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:orderID", controllers.GetOrderById)
	router.PATCH("/orders/:orderID", controllers.UdateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)
	return router
}

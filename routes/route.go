package routes

import (
	"golang3/controllers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine{
	router := gin.Default()

	router.GET("/",controllers.ViewBiodata)
	router.GET("/find/:id",controllers.ViewOneBiodata)
	router.POST("/create",controllers.InsertBiodata)
	router.GET("/update/:id",controllers.UpdateBiodata)
	router.GET("/delete/:id",controllers.DeleteBiodata)

	return router
}
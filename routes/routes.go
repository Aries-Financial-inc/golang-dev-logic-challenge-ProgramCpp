package routes

import (
	"github.com/aries-financial-inc/options-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		controllers.AnalysisHandler(c.Writer, c.Request)
	})

	return router
}

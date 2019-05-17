package config

import "github.com/gin-gonic/gin"
import "pine/app/controllers"

func Route(r *gin.Engine) {
	r.GET("/", controllers.IndexHandle)
}

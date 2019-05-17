package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context)  {
	c.String(http.StatusOK, "index handler")
}
package handler

import "github.com/gin-gonic/gin"

type Timeline interface {
	Get(c *gin.Context)
	NoRoute(c *gin.Context)
	NoMethod(c *gin.Context)
}

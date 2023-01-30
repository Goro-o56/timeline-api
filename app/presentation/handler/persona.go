package handler

import "github.com/gin-gonic/gin"

type Persona interface {
	Find(c *gin.Context)
}

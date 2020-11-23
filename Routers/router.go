package routers

import (
	v1 "go-gin-structure/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/test", v1.GetTags)
	apiv1.POST("/test", v1.AddTag)
	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kevinchfe/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.Article{}
	tag := v1.Tag{}

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.Get)
		apiv1.GET("/tags/:id", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/article/:id", article.Delete)
		apiv1.PUT("/article/:id", article.Update)
		apiv1.PATCH("/article/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}
	return r
}

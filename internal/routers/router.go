package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinchfe/blog-service/internal/middleware"
	v1 "github.com/kevinchfe/blog-service/internal/routers/api/v1"
	_ "github.com/kevinchfe/blog-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.Article{}
	tag := v1.Tag{}

	apiv1 := r.Group("/api/v1")
	{
		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("tags/:id", tag.Delete)
		// 更新标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 标签列表
		apiv1.GET("/tags", tag.List)
		// 获取指定标签
		apiv1.GET("/tags/:id", tag.Get)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/article/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/article/:id", article.Update)
		// 文章列表
		apiv1.GET("/articles", article.List)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
	}
	return r
}

package router

import (
    "github.com/gin-gonic/gin"

    "go-gin-blog/pkg/setting"
    "go-gin-blog/router/api/v1"
    "go-gin-blog/router/api"
    // "go-gin-blog/middleware/jwt"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    gin.SetMode(setting.ServerSetting.RunMode)

    r.GET("/auth", api.GetAuth)

    apiv1 := r.Group("/api/v1")
    // apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", v1.GetTags)
        //新建标签
        // http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=woo
        apiv1.POST("/tags", v1.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTag)
        //删除指定标签
        apiv1.DELETE("/tags/:id", v1.DeleteTag)

        //获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)
        //新建文章
        // http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=blog1&desc=blog-desc&content=blog-content&created_by=blog-created&state=1
        apiv1.POST("/articles", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
    }
    return r
}


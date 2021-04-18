package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaggerwoo/go-gin-blog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
        apiv1.GET("/tags", GetTags)
        //新建标签
        apiv1.POST("/tags", AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", EditTag)
        //删除指定标签
        apiv1.DELETE("/tags/:id", DeleteTag)
	}

    return r
}

//获取多个文章标签
func GetTags(c *gin.Context) {
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
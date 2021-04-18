package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-blog/pkg/setting"
	"go-gin-blog/pkg/e"
    "go-gin-blog/models"
    "go-gin-blog/pkg/util"
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
	name := c.Query("name")
	maps := make(map[string]interface{})
    data := make(map[string]interface{})
	if name != "" {
        maps["name"] = name
    }
	var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state
    }
	code := e.SUCCESS

    data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
    data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
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
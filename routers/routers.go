package routers

import (
	"github.com/gin-gonic/gin"

	"cumtBook/middleware/jwt"
	"cumtBook/pkg/setting"
	"cumtBook/routers/api"
	v1 "cumtBook/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取预约列表
		apiv1.GET("/books", v1.GetBooks)
		// 获取指定预约
		apiv1.GET("/books/:book_id", v1.GetBook)
		// 新增预约
		apiv1.POST("/books", v1.AddBook)
		// 删除指定预约
		apiv1.DELETE("/books/:book_id", v1.DelBook)
		// 更新指定预约
		// apiv1.PUT("/books/:book_id", v1.EditBook)
	}

	return r
}

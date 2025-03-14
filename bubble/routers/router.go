package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	//配置静态文件
	r.Static("/static", "static")

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//v1

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.PostATodoList)

		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改
		v1Group.PUT("/todo/:id", controller.PutATodoList)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodoList)
	}

	return r
}

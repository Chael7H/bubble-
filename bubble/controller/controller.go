package controller

import (
	"bubble/dao"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func PostATodoList(c *gin.Context) {
	//	前端界面点击提交的数据会发送到此处
	//1.从请求中拿到数据
	var todo models.Todo
	_ = c.ShouldBind(&todo)
	//2.存入到数据库,响应到前端
	if err := dao.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
			"msg":  "success",
		})
	}

}

func GetTodoList(c *gin.Context) {
	//查询todos表里的所有数据
	if todoList, err := dao.GetAllTodos(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func PutATodoList(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := dao.GetATodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "绑定JSON数据失败: " + err.Error(),
		})
		return
	}
	if err := dao.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodoList(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := dao.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "delete",
		})
	}
}

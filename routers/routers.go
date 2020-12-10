package routers

import (
	"bubble/controller"
	"bubble/dao"
	"bubble/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "static")
	// 加载模板
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/index", controller.IndexHandler)
	// v1
	v1Group := r.Group("v1")
	// 待办事项
	// 添加
	v1Group.POST("/todo", controller.CreateTodo)
	// 查看
	// 查看所有代办事项
	v1Group.GET("/todo", controller.GetAllTodo)
	// 查看某一个待办事项
	v1Group.GET("/todo/:id", func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			fmt.Printf("no id failed")
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		todoId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("no id failed")
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		var thing modles.Todo
		dao.DB.Model(&modles.Todo{}).First(&thing).Where("id = ?", todoId)
		if thing.ID != 0 {
			c.JSON(http.StatusOK, gin.H{
				"thing": thing,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "no such id",
			})
		}

	})
	// 修改事项状态
	v1Group.PUT("/todo/:id", controller.UpdateTodo)

	// 删除
	v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	return r
}

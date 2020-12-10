package controller

import (
	"bubble/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
// url --> controller --> logic/service --> model
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写代办事项，点击提交，会发送请求到这里
	// 1.从请求中把数据拿出来
	//title := c.PostForm("title")
	var todo modles.Todo
	c.BindJSON(&todo)
	//thing := Todo{
	//	Title:  title,
	//	Status: false,
	//}

	if err := modles.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"thing": todo,
		})
	}

}

func GetAllTodo(c *gin.Context) {
	things, err := modles.GetTodoList()
	if  err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, things)
	}

}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error": "id不存在",
		})
		return
	}
	i, _ := strconv.Atoi(id)
	todo, err := modles.GetATodo(i)
	if  err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.BindJSON(todo)
	err = modles.UpdateATodo(todo)
	if  err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todo)
	}

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		fmt.Printf("no id failed")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := modles.DeleteATodo(id)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"msg": "success",
		})
	}
}
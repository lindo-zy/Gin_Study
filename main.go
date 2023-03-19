package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//默认初始化
	r := gin.Default()
	//原生初始化
	//r := gin.New()

	//get请求
	r.GET("/", func(context *gin.Context) {
		context.JSONP(200, gin.H{
			"message": "hello world!",
		})
	})

	//路径参数,会匹配到/user/tony,匹配成功后输出
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(200, gin.H{
			"name": name,
		})
	})

	//路径参数，贪婪匹配,/items/ 和/items/action都可以匹配,匹配成功会输出带/的字符串
	r.GET("/items/*action", func(context *gin.Context) {
		action := context.Param("action")
		fmt.Println(action)
		context.String(http.StatusOK, action)

	})

	//run方法
	err := r.Run("127.0.0.1:8088")
	if err != nil {
		return
	}
}

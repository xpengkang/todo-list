package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 链接数据库
// 接受用户请求
// 一个代办事项作为一个结构体，存放在数据库。
//需要创建数据库，
// 需要操作的对象
type Todo struct {
	//ID int 'json:"id"'
	//Title string 'json:"Titel"'
	//Status bool 'json:"Status"'
	// 反引号
	ID int `json:"id"`
	TiTle string `json:"titel"`
	Status bool `json:"status"`

}
func main() {
	// 链接数据库
	//sql


	// gin 框架 // 接受用户请求
	r := gin.Default()

	// 获取静态资源, // 把前端的代码拷贝过来，npm
	r.Static("/static", "static")
	r.LoadHTMLGlob("temp/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run() // 创建路由, 本地打开127.0.0.1:8080


	//分析：crud 用gorm 框架, RestFul 风格接口
	v1Group := r.Group("v1"){
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			
		})
		//查看all
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//改
		v1Group.PUT("todo/:id", func(c *gin.Context) {

		})

		//删除
		v1Group.DELETE("todo/:id", func(c *gin.Context) {

		})
	}



	

}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 全局对象
var DB *gorm.DB

// 链接数据库
// 接受用户请求
// 一个代办事项作为一个结构体，存放在数据库。
//需要创建数据库，
// 需要操作的 model
type Todo struct {
	//ID int 'json:"id"'
	//Title string 'json:"Titel"'
	//Status bool 'json:"Status"'
	// 反引号
	ID int `json:"id"`
	TiTle string `json:"titel"`
	Status bool `json:"status"`

}
// 数据库只有链接成功了，才有后续的操作
func initMySQL() (err error){
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) // 这里不需要:=
	if err != nil{
		return
	}
	return DB.DB().Ping() // 如果连接成功了，会返回nil
}


func main() {
	// 链接数据库
	//sql
	err := initMySQL()
	if err != nil {
		panic(err) // 有问题，直接不后续操作了
	}
	defer DB.Close() // 完整的退出

	// 模型绑定；将model 和 表 关联
	DB.AutoMigrate(&Todo{})



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

package main

import (
	"github.com/gin-gonic/gin"

	"MyBlog/server/views/articles"
	"MyBlog/server/views/users"
)

func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()

	// 接口
	router.POST("/api/users/sign_up", users.SignUp)
	router.POST("/api/users/login", users.Login)
	router.POST("/api/users/update_password", users.UpdatePassword)

	router.POST("/api/articles/find", articles.Find)
	router.POST("/api/articles/get", articles.GetArticle)
	router.POST("/api/articles/post", articles.Post)
	router.POST("/api/articles/delete", articles.DeleteArticle)

	// 指定地址和端口号
	router.Run(":34182")

}

// Package view
package users

import (
	errors2 "MyBlog/server/views/errors"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/server"
)

// User 请求参数
type User struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

// Response 返回参数
type Response struct {
	Ok bool `json:"ok"`
}

// PostHandle 处理请求函数
func SignUp(c *gin.Context) {

	request := &User{}
	err := binding.JSON.Bind(c.Request, request)
	if err != nil {
		c.JSON(200, errors2.MakeErrorResponse(err))
		return
	}

	err = request.ValidateRequestParams()
	if err != nil {
		c.JSON(200, errors2.MakeErrorResponse(err))
		return
	}

	err = server.SignUp(request.UserName, request.Gender, request.Phone, request.Password)
	if err != nil {
		fmt.Println("出错了", err)
		c.JSON(200, Response{Ok: false})
		return
	}

	c.JSON(200, Response{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *User) ValidateRequestParams() error {

	if g.UserName == "" {
		return errors.New("invalid_param.miss_user_name")
	}
	if g.Password == "" {
		return errors.New("invalid_param.password")
	}
	if g.Phone == "" {
		return errors.New("invalid_param.miss_phone")
	}
	return nil
}

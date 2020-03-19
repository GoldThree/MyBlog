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

// LoginRequest 请求参数
type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Response 返回参数
type LoginResponse struct {
	Ok bool `json:"ok"`
}

// PostHandle 处理请求函数
func Login(c *gin.Context) {

	request := &LoginRequest{}
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

	err = server.Login(request.Phone, request.Password)
	if err != nil {
		fmt.Println("出错了", err)
		c.JSON(200, Response{Ok: false})
		return
	}

	c.JSON(200, Response{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *LoginRequest) ValidateRequestParams() error {

	if g.Password == "" {
		return errors.New("invalid_param.password")
	}
	if g.Phone == "" {
		return errors.New("invalid_param.miss_phone")
	}
	return nil
}

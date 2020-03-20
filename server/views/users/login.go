// Package view
package users

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/server"
	e "MyBlog/server/views/errors"
)

// LoginRequest 请求参数
type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Response 返回参数
type LoginResponse struct {
	Token string `json:"token"`
	Ok    bool   `json:"ok"`
}

// PostHandle 处理请求函数
func Login(c *gin.Context) {

	request := &LoginRequest{}
	err := binding.JSON.Bind(c.Request, request)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	err = request.ValidateRequestParams()
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	token, err := server.Login(request.Phone, request.Password)
	if err != nil {
		fmt.Println("出错了", err)
		c.JSON(200, Response{Ok: false})
		return
	}

	c.JSON(200, LoginResponse{Ok: true, Token: token})

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

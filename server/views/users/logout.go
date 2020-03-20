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

// LogoutRequest 请求参数
type LogoutRequest struct {
	Uuid  string `json:"uuid"`
	Token string `json:"token"`
}

// Response 返回参数
type LogoutResponse struct {
	Ok bool `json:"ok"`
}

// PostHandle 处理请求函数
func Logout(c *gin.Context) {

	request := &LogoutRequest{}
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

	err = server.Logout(request.Uuid, request.Token)
	if err != nil {
		fmt.Println("出错了", err)
		c.JSON(200, LogoutResponse{Ok: false})
		return
	}

	c.JSON(200, LogoutResponse{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *LogoutRequest) ValidateRequestParams() error {

	if g.Uuid == "" {
		return errors.New("invalid_param.uuid")
	}
	if g.Token == "" {
		return errors.New("invalid_param.token")
	}
	return nil
}

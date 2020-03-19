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

// UpdatePasswordRequest 请求参数
type UpdatePasswordRequest struct {
	UserUuid    string `json:"user_uuid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// Response 返回参数
type UpdatePasswordResponse struct {
	Ok bool `json:"ok"`
}

// UpdatePassword 处理请求函数
func UpdatePassword(c *gin.Context) {

	request := &UpdatePasswordRequest{}
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

	err = server.UpdatePassword(request.UserUuid, request.OldPassword, request.NewPassword)
	if err != nil {
		fmt.Println("出错了", err)
		c.JSON(200, Response{Ok: false})
		return
	}

	c.JSON(200, Response{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *UpdatePasswordRequest) ValidateRequestParams() error {

	if g.OldPassword == "" {
		return errors.New("invalid_param.old_password")
	}
	if g.NewPassword == "" {
		return errors.New("invalid_param.new_password")
	}
	if g.UserUuid == "" {
		return errors.New("invalid_param.user_uuid	")
	}
	return nil
}

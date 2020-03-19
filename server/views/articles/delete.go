package articles

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/server"
	e "MyBlog/server/views/errors"
)

// DeleteArticlesRequest 请求参数
type DeleteArticlesRequest struct {
	Uuid string `json:"uuid"`
}

// DeleteArticlesResponse 返回参数
type DeleteArticlesResponse struct {
	Ok bool `json:"ok"`
}

// DeleteArticle 处理请求函数
func DeleteArticle(c *gin.Context) {

	request := &DeleteArticlesRequest{}
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

	err = server.DeleteArticle(request.Uuid)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	c.JSON(200, DeleteArticlesResponse{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *DeleteArticlesRequest) ValidateRequestParams() error {

	if g.Uuid == "" {
		return errors.New("invalid_param.uuid")
	}
	return nil
}


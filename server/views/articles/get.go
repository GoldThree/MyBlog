package articles

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/models"
	"MyBlog/server/server"
	e "MyBlog/server/views/errors"
)

// PostArticlesRequest 请求参数
type GetArticlesRequest struct {
	Uuid string `json:"uuid"`
}

// PostArticlesResponse 返回参数
type GetArticlesResponse struct {
	Article models.Article `json:"article"`
}

// PostHandle 处理请求函数
func GetArticle(c *gin.Context) {

	request := &GetArticlesRequest{}
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

	article, err := server.GetArticle(request.Uuid)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	c.JSON(200, GetArticlesResponse{Article: article})

}

// ValidateRequestParams 参数检查
func (g *GetArticlesRequest) ValidateRequestParams() error {

	if g.Uuid == "" {
		return errors.New("invalid_param.uuid")
	}
	return nil
}

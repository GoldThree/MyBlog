package articles

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/server"
	e "MyBlog/server/views/errors"
)

// PostArticlesRequest 请求参数
type PostArticlesRequest struct {
	Title      string `json:"title"`
	Token      string `json:"token"`
	Content    string `json:"content"`
	AuthorUuid string `json:"author_uuid"`
}

// PostArticlesResponse 返回参数
type PostArticlesResponse struct {
	Ok bool `json:"ok"`
}

// PostHandle 处理请求函数
func Post(c *gin.Context) {

	request := &PostArticlesRequest{}
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

	err = server.Post(request.Title, request.Content, request.AuthorUuid, request.Token)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	c.JSON(200, PostArticlesResponse{Ok: true})

}

// ValidateRequestParams 参数检查
func (g *PostArticlesRequest) ValidateRequestParams() error {

	if g.Title == "" {
		return errors.New("invalid_param.title")
	}
	if g.Content == "" {
		return errors.New("invalid_param.content")
	}
	if g.AuthorUuid == "" {
		return errors.New("invalid_param.author")
	}
	if g.Token == "" {
		return errors.New("invalid_param.token")
	}
	return nil
}

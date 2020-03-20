package articles

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"MyBlog/server/models"
	"MyBlog/server/server"
	e "MyBlog/server/views/errors"
)

// FindArticlesRequest 请求参数
type FindArticlesRequest struct {
	Uuid   string `json:"uuid"`
	Token  string `json:"token"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

// FindArticlesResponse 返回参数
type FindArticlesResponse struct {
	Articles []models.Article `json:"articles"`
}

// PostHandle 处理请求函数
func Find(c *gin.Context) {

	request := &FindArticlesRequest{}
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

	articles, err := server.Find(request.Offset, request.Limit, request.Uuid, request.Token)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	c.JSON(200, FindArticlesResponse{Articles: articles})

}

// ValidateRequestParams 参数检查
func (g *FindArticlesRequest) ValidateRequestParams() error {

	if g.Uuid == "" {
		return errors.New("invalid_param.uuid")
	}
	if g.Token == "" {
		return errors.New("invalid_param.token")
	}
	if g.Offset < 0 {
		return errors.New("invalid_param.offset")
	}
	if g.Limit < 0 {
		return errors.New("invalid_param.limit")
	}
	return nil
}

// FindUserArticlesRequest 请求参数
type FindUserArticlesRequest struct {
	Uuid   string `json:"uuid"`
	Token  string `json:"token"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

// FindUserArticlesResponse 返回参数
type FindUserArticlesResponse struct {
	Articles []models.Article `json:"articles"`
}

// PostHandle 处理请求函数
func FindByUuid(c *gin.Context) {

	request := &FindUserArticlesRequest{}
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

	articles, err := server.FindByUuid(request.Uuid, request.Token, request.Offset, request.Limit)
	if err != nil {
		c.JSON(200, e.MakeErrorResponse(err))
		return
	}

	c.JSON(200, FindUserArticlesResponse{Articles: articles})

}

// ValidateRequestParams 参数检查
func (g *FindUserArticlesRequest) ValidateRequestParams() error {

	if g.Uuid == "" {
		return errors.New("invalid_param.uuid")
	}
	if g.Token == "" {
		return errors.New("invalid_param.token")
	}
	if g.Offset < 0 {
		return errors.New("invalid_param.offset")
	}
	if g.Limit < 0 {
		return errors.New("invalid_param.limit")
	}
	return nil
}

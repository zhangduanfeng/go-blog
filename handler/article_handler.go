package handler

import (
	"go-blog/dto"
	"go-blog/errno"
	"go-blog/model"
	"go-blog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var req = &dto.CreateArticleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	id, err := service.CreateArticle(req.Title, req.Content, req.CreateId)
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &dto.CreateArticleResponse{
		ArticleId: id,
	}))
	return
}

func ListArticles(c *gin.Context) {
	// request params
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// query from db
	articles, total, err := service.ListArticles(int64(pageNum), int64(pageSize))
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	// conv resp
	var result = make([]*dto.Article, 0)
	for _, item := range articles {
		result = append(result, convArticleDO2DTO(item))
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &dto.ListArticlesResponse{
		Articles: result,
		PageNo:   int64(pageNum),
		PageSize: int64(pageSize),
		Total:    total,
	}))
	return
}

func convArticleDO2DTO(article *model.Article) *dto.Article {
	var result = &dto.Article{
		Id:         article.Id,
		CreateTime: article.CreateTime,
		CreateId:   article.CreateId,
		UpdateId:   article.UpdateId,
		UpdateTime: article.UpdateTime,
		Title:      article.Title,
		Content:    article.Content,
	}
	return result
}
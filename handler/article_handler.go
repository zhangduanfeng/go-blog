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

}

func ListArticles(c *gin.Context) {
	// request params
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// query from db
	articles, err := service.ListArticles(int64(pageNum), int64(pageSize))
	if err != nil {
		// TODO: log
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	// conv resp
	var result = make([]*dto.Article, 0)
	for _, item := range articles {
		result = append(result, convArticleDO2DTO(item))
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &dto.ListArticlesResponse{Articles: result}))
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

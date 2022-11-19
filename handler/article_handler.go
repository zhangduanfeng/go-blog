package handler

import (
	"go-blog/errno"
	"go-blog/model"
	"go-blog/service"
	"go-blog/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var req = &vo.CreateArticleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	id, err := service.CreateArticle(req.Title, req.Content, req.CreateId)
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.CreateArticleResponse{
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
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	// conv resp
	var result = make([]*vo.Article, 0)
	for _, item := range articles {
		result = append(result, convArticleDO2VO(item))
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.ListArticlesResponse{
		Articles: result,
		PageNo:   int64(pageNum),
		PageSize: int64(pageSize),
		Total:    total,
	}))
	return
}

func PreviewArticle(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.DefaultQuery("articleId", "0"))

	article, err := service.GetArticleById(int64(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.Article{
		Id:         article.Id,
		CreateTime: article.CreateTime.String(),
		CreateId:   article.CreateId,
		UpdateId:   article.UpdateId,
		UpdateTime: article.UpdateTime.String(),
		Title:      article.Title,
		Content:    article.Content,
	}))
	return
}

func convArticleDO2VO(article *model.Article) *vo.Article {
	var result = &vo.Article{
		Id:         article.Id,
		CreateTime: article.CreateTime.String(),
		CreateId:   article.CreateId,
		UpdateId:   article.UpdateId,
		UpdateTime: article.UpdateTime.String(),
		Title:      article.Title,
		Content:    article.Content,
	}
	return result
}

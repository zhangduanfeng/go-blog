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

/**
 * @Description 发布文章
 * @Param
 * @return
 **/
func CreateArticle(c *gin.Context) {
	var req = &vo.CreateArticleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	if req.CreateId == 0 {
		req.CreateId = 10
	}
	id, err := service.CreateArticle(req.Title, req.Content, req.CreateId)
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.CreateArticleResponse{
		ArticleId: id,
	}))
}

/**
 * @Description 文章列表
 * @Param
 * @return
 **/
func ListArticles(c *gin.Context) {
	// request params
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	searchTitle := c.DefaultQuery("searchTitle", "")

	// query from db
	articles, total, err := service.ListArticles(int64(pageNum), int64(pageSize), searchTitle)
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
}

/**
 * @Description 文章详情
 * @Param
 * @return
 **/
func ArticleDetails(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.DefaultQuery("articleId", "0"))
	//根据文章id查询标签
	tagIds, err := service.GetTagByArticleId(int64(articleId))
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	var articleTags = make([]*vo.ArticleTag, 0)

	if len(tagIds) != 0 {
		tagNames, err := service.GetTagNameById(tagIds)
		if err != nil {
			c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
			return
		}
		for _, tag := range tagNames {
			var result = &vo.ArticleTag{
				Id:      tag.Id,
				TagName: tag.Name,
			}
			articleTags = append(articleTags, result)
		}
	}

	article, err := service.GetArticleById(int64(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.Article{
		Id:          article.Id,
		CreateTime:  article.CreateTime.String(),
		CreateId:    article.CreateId,
		UpdateId:    article.UpdateId,
		UpdateTime:  article.UpdateTime.String(),
		Title:       article.Title,
		Content:     article.Content,
		ArticleTags: articleTags,
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

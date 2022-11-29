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
	id, err := service.CreateArticle(req.Title, 10)
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.CreateArticleResponse{
		ArticleId: id,
	}))
}

func SaveArticle(c *gin.Context) {
	var req = &vo.SaveArticleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	if err := service.UpdateArticle(req.ArticleId, &req.Title, &req.Content, &req.Summary, &req.CoverImg, &req.Cate, 10, 0); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", nil))
}

func PublishArticle(c *gin.Context) {
	var req = &vo.PublishArticleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	if err := service.UpdateArticle(req.ArticleId, &req.Title, &req.Content, &req.Summary, &req.CoverImg, &req.Cate, 10, 1); err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.ConstructResp("", "", nil))
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
	tagId, _ := strconv.Atoi(c.DefaultQuery("tagId", ""))
	searchTitle := c.DefaultQuery("searchTitle", "")
	articleIds, err := service.GetArticleByTagId(int64(tagId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}
	// query from db
	articles, total, err := service.ListArticles(articleIds, int64(pageNum), int64(pageSize), searchTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	var cateIds = make([]int64, 0)
	// 后续ES/缓存优化
	for _, item := range articles {
		cateIds = append(cateIds, item.CateId)
	}

	cateInfos, err := service.GetCategoryByIds(cateIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	// conv resp
	var result = make([]*vo.Article, 0)
	var cateMap = make(map[int64]*model.Category)
	for _, cate := range cateInfos {
		cateMap[cate.Id] = cate
	}
	for _, item := range articles {
		var cateInfo = &model.Category{}
		cateInfo = cateMap[item.CateId]
		result = append(result, convArticleDO2VO(item, cateInfo, nil))
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

	cateInfo, err := service.GetCategoryById(article.CateId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.ConstructResp("", "", convArticleDO2VO(article, cateInfo, articleTags)))
}

func convArticleDO2VO(article *model.Article, cateInfo *model.Category, tagsInfos []*vo.ArticleTag) *vo.Article {
	var result = &vo.Article{
		Id:          article.Id,
		CreateTime:  article.CreateTime.Format("2006-01-02 15:04:05"),
		CreateId:    article.CreateId,
		UpdateId:    article.UpdateId,
		UpdateTime:  article.UpdateTime.Format("2006-01-02 15:04:05"),
		Title:       article.Title,
		Content:     article.Content,
		CoverImg:    article.CoverImg,
		Summary:     article.Summary,
		ArticleTags: tagsInfos,
	}
	if cateInfo != nil {
		result.CateInfo = &vo.Category{
			Id:       cateInfo.Id,
			Name:     cateInfo.Name,
			Sequence: cateInfo.Sequence,
		}
	}
	return result
}

package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/src/model"
	"go-blog/src/store"
	"go-blog/src/utils/errmsg"
	"net/http"
	"strconv"
)

/**
 * @Description blog前台首页接口
 * @Author duanfeng.zhang
 * @Date 2022/8/12 16:53
 **/
func GetArticles(c *gin.Context) {
	//分页
	var total int
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	articles := make([]model.Article, 0)
	if err := store.DB.Model(articles).Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": err.Error(),
		})
		return
	}
	//在这里判断是为了减少一次
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":     1,
			"data":     articles,
			"total":    total,
			"page":     pageNum,
			"pageSize": pageSize,
		})
		return
	}
	offset := (pageNum - 1) * pageSize
	result := store.DB.Debug().Order("id DESC").Offset(offset).Limit(pageSize).Find(&articles)
	if result.Error != nil && result.Error.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    0,
			"message": errmsg.GetErrmsg(errmsg.ERROR),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
		"data":     articles,
	})
}

package bg

import (
	"github.com/gin-gonic/gin"
	"go-blog/src/model"
	"net/http"
)

/**
 * @Description 发布文章
 * @Author duanfeng.zhang
 * @Date 2022/8/19 16:54
 **/
func AddArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"message": "参数解析异常",
		})
	}
}

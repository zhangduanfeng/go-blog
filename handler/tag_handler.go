package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-blog/errno"
	"go-blog/model"
	"go-blog/service"
	"go-blog/vo"
	"net/http"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/11/14 17:55
 **/
func ListCategoryAndTag(c *gin.Context) {
	logrus.Info("哈哈哈哈")
	categories, err := service.ListCategories()

	tags, err := service.ListTags()
	if err != nil {
		c.JSON(http.StatusOK, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	var tagResponse = &vo.TagsResponse{}
	var resultTag = make([]*vo.Tag, 0)
	var resultCategory = make([]*vo.Category, 0)

	for _, item := range tags {
		resultTag = append(resultTag, convTagDO2VO(item))
	}

	for _, item := range categories {
		resultCategory = append(resultCategory, convCategoryDO2VO(item))
	}

	tagResponse.Categories = resultCategory
	tagResponse.Tags = resultTag

	c.JSON(http.StatusOK, errno.ConstructResp("", "", &vo.TagsResponse{
		Tags:       resultTag,
		Categories: resultCategory,
	}))
	return
}

func convTagDO2VO(tag *model.Tag) *vo.Tag {
	var result = &vo.Tag{
		Id:       tag.Id,
		Name:     tag.Name,
		Code:     tag.Code,
		Sequence: tag.Sequence,
	}
	return result
}

func convCategoryDO2VO(category *model.Category) *vo.Category {
	var result = &vo.Category{
		Id:       category.Id,
		Name:     category.Name,
		Sequence: category.Sequence,
	}
	return result
}

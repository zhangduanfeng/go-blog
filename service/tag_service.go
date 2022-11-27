package service

import (
	"go-blog/dal/db"
	"go-blog/model"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/11/14 17:59
 **/
func ListCategories() ([]*model.Category, error) {
	categories, err := db.QueryCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func ListTags() ([]*model.Tag, error) {
	tags, err := db.QueryTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetCategoryById(cateId int64) (*model.Category, error) {
	category, err := db.GetCategoryById(cateId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetCategoryByIds(cateIds []int64) ([]*model.Category, error) {
	if len(cateIds) == 0 {
		return nil, nil
	}
	categories, err := db.GetCategoriesByIds(cateIds)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetArticleByTagId(tagId int64) ([]int64, error) {
	articleIds, err := db.GetArticleByTagId(tagId)
	if err != nil {
		return nil, err
	}
	return articleIds, nil
}

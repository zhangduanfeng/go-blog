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

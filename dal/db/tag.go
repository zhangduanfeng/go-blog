package db

import (
	"go-blog/model"
	"go-blog/store"
)

/**
 * @Description 获取所有文章标签
 * @Author duanfeng.zhang
 * @Date 2022/11/14 18:04
 **/
func QueryTags() ([]*model.Tag, error) {
	var tags = make([]*model.Tag, 0)
	err := store.DB.Debug().Order("id").Find(&tags).Error
	if err != nil {
		if err.Error() == "record not found" {
			return tags, nil
		}
		return nil, err
	}

	return tags, nil
}

/**
 * @Description 获取分类
 * @Author duanfeng.zhang
 * @Date 2022/11/14 18:04
 **/
func QueryCategory() ([]*model.Category, error) {
	var categorys = make([]*model.Category, 0)
	err := store.DB.Debug().Order("id").Find(&categorys).Error
	if err != nil {
		if err.Error() == "record not found" {
			return categorys, nil
		}
		return nil, err
	}

	return categorys, nil
}

func GetCategoryById(id int64) (*model.Category, error) {
	var category = &model.Category{}
	err := store.DB.Debug().Where("id = ?", id).Find(&category).Error
	if err != nil {
		if err.Error() == "record not found" {
			return category, nil
		}
		return nil, err
	}
	return category, nil
}

func GetCategoriesByIds(cateIds []int64) ([]*model.Category, error) {
	var categorys = make([]*model.Category, 0)
	err := store.DB.Debug().Order("id").Where("id in (?)", cateIds).Find(&categorys).Error
	if err != nil {
		if err.Error() == "record not found" {
			return categorys, nil
		}
		return nil, err
	}

	return categorys, nil
}
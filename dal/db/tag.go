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

func GetTagNameById(tagIds []int64) ([]*model.Tag, error) {
	var tags = make([]*model.Tag, 0)
	err := store.DB.Debug().Table("tag").Select("id,name").Where("id in (?)", tagIds).Find(&tags).Error
	if err != nil {
		if err.Error() == "record not found" {
			return tags, nil
		}
		return nil, err
	}

	return tags, nil
}

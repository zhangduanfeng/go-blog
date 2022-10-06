package db

import (
	"go-blog/config"
	"go-blog/model"
	"go-blog/store"
)

func CountArticles() (int64, error) {
	var total int64
	var articles = make([]*model.Article, 0)
	if err := store.DB.Model(&articles).Where("valid = ?", config.IS_VALID).Count(&total); err != nil {
		return 0, nil
	}
	return total, nil
}

func PageQueryArticles(offset, limit int64) ([]*model.Article, error) {
	var articles = make([]*model.Article, 0)
	err := store.DB.Debug().Order("id DESC").Offset(offset).Limit(limit).Find(&articles).Error
	if err.Error() == "record not found" {
		return articles, nil
	}
	if err != nil {
		return nil, err
	}

	return articles, nil
}

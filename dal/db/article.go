package db

import (
	"go-blog/config"
	"go-blog/model"
	"go-blog/store"
)

func CountArticles() (int64, error) {
	var total int64
	if err := store.DB.Debug().Table("article").Where("valid = ?", config.IS_VALID).Count(&total).Error; err != nil {
		return 0, nil
	}
	return total, nil
}

func PageQueryArticles(offset, limit int64) ([]*model.Article, error) {
	var articles = make([]*model.Article, 0)
	err := store.DB.Debug().Order("id DESC").Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		if err.Error() == "record not found" {
			return articles, nil
		}
		return nil, err
	}

	return articles, nil
}

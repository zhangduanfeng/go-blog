package db

import (
	"go-blog/config"
	"go-blog/model"
	"go-blog/store"
)

func CountArticles(searchTitle string) (int64, error) {
	var total int64
	if err := store.DB.Debug().Table("article").Where("title like ? and valid = ?", searchTitle, config.IS_VALID).Count(&total).Error; err != nil {
		return 0, nil
	}
	return total, nil
}

func CreateArticles(article *model.Article) (int64, error) {
	if err := store.DB.Debug().Table("article").Create(article).Error; err != nil {
		return 0, err
	}
	return article.Id, nil
}

func PageQueryArticles(searchTitle string, offset, limit int64) ([]*model.Article, error) {
	var articles = make([]*model.Article, 0)
	err := store.DB.Debug().Where("title like ?", searchTitle).Order("id DESC").Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		if err.Error() == "record not found" {
			return articles, nil
		}
		return nil, err
	}

	return articles, nil
}

func GetArticleById(id int64) (*model.Article, error) {
	var article = &model.Article{}
	err := store.DB.Debug().Where("id = ?", id).Find(&article).Error
	if err != nil {
		if err.Error() == "record not found" {
			return article, nil
		}
		return nil, err
	}
	return article, nil
}

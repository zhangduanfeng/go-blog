package service

import (
	"go-blog/dal/db"
	"go-blog/model"
)

func ListArticles(pageNo, pageSize int64) ([]*model.Article, error) {
	var articles = make([]*model.Article, 0)

	//分页
	count, err := db.CountArticles()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return articles, nil
	}

	articles, err = db.PageQueryArticles((pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

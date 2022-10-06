package service

import (
	"go-blog/dal/db"
	"go-blog/model"
	"time"
)

func ListArticles(pageNo, pageSize int64) ([]*model.Article, int64, error) {
	var articles = make([]*model.Article, 0)

	//分页
	count, err := db.CountArticles()
	if err != nil {
		return nil, 0, err
	}

	if count == 0 {
		return articles, 0, nil
	}

	articles, err = db.PageQueryArticles((pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

func CreateArticle(title, content string, createId int64) (int64, error) {
	var article = &model.Article{
		CreateId:   createId,
		UpdateId:   createId,
		Title:      title,
		Content:    content,
		CreateTime: time.Now().String(),
		UpdateTime: time.Now().String(),
	}
	return db.CreateArticles(article)
}

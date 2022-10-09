package service

import (
	"errors"
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

func GetArticleById(id int64) (*model.Article, error) {
	article, err := db.GetArticleById(id)
	if err != nil {
		return nil, err
	}
	if article.Id == 0 {
		return nil, errors.New("此文章不存在")
	}
	return article, nil
}

func CreateArticle(title, content string, createId int64) (int64, error) {
	var article = &model.Article{
		CreateId:   createId,
		UpdateId:   createId,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return db.CreateArticles(article)
}

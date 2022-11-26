package db

import (
	"go-blog/store"
)

func GetTagByArticleId(articleId int64) ([]int64, error) {
	var tagIds = make([]int64, 0)
	err := store.DB.Debug().Table("article_tag").Where("article_id = ?", articleId).Pluck("tag_id", &tagIds).Error
	if err != nil {
		if err.Error() == "record not found" {
			return tagIds, nil
		}
		return nil, err
	}
	return tagIds, nil
}

package model

import "time"

type ArticleTag struct {
	Id         int64     `json:"id"`
	ReqSource  int64     `json:"req_source" label:"打标来源 0用户自选，1系统打标，2管理员打标"`
	ArticleId  int64     `json:"article_id" label:"文章id"`
	UserId     int64     `json:"user_id" label:"用户id"`
	Code       string    `json:"code" label:"标签Code"`
	TagId      int64     `json:"tag_id" label:"标签id"`
	Extra      string    `json:"extra" label:"拓展信息"`
	CreateTime time.Time `json:"createTime"`
	CreateId   int64     `json:"createId"`
	UpdateTime time.Time `json:"updateTime"`
}

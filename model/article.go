package model

import "time"

/**
 * @Description 文章结构体
 * @Author duanfeng.zhang
 * @Date 2022/8/12 16:56
 **/
type Article struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `json:"createTime"`
	CreateId   int64     `json:"createId"`
	UpdateId   int64     `json:"updateId"`
	UpdateTime time.Time `json:"updateTime"`
	Title      string    `json:"title" validate:"required" label:"文章标题"`
	Content    string    `json:"content" validate:"required" label:"文章内容"`
	Summary    string    `json:"summary"`
	CoverImg   string    `json:"cover_img"`
	CateId     int64     `json:"cate_id"`
	Valid      int64     `json:"valid" label:"是否有效0-有效1-无效"`
	Status     int64     `json:"status"`
}

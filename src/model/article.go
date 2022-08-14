package model

/**
 * @Description 文章结构体
 * @Author duanfeng.zhang
 * @Date 2022/8/12 16:56
 **/
type Article struct {
	Id          int64  `json:"id"`
	CreateTime  string `json:"createTime"`
	CreateId    int64  `json:"createId"`
	UpdateId    int64  `json:"updateId"`
	UpdateTime  string `json:"updateTime"`
	Title       string `json:"title" validate:"required" label:"文章标题"`
	Content     string `json:"content" validate:"required" label:"文章内容"`
	TagId       int64  `json:"tag_id" validate:"required" label:"标签id"`
	LikeCount   int64  `json:"like_count" validate:"required" label:"点赞数"`
	BrowseCount int64  `json:"browse_count" validate:"required" label:"浏览数"`
	DeleteFlag  int    `json:"delete_flag" label:"是否有效0-有效1-无效"`
}

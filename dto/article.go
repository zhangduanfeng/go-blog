package dto

type ListArticlesResponse struct {
	Articles []*Article
}

type Article struct {
	Id          int64  `json:"id"`
	CreateTime  string `json:"createTime"`
	CreateId    int64  `json:"createId"`
	UpdateId    int64  `json:"updateId"`
	UpdateTime  string `json:"updateTime"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

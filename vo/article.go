package vo

type ListArticlesResponse struct {
	Articles []*Article
	PageNo   int64
	PageSize int64
	Total    int64
}

type Article struct {
	Id          int64         `json:"id"`
	CreateTime  string        `json:"createTime"`
	CreateId    int64         `json:"createId"`
	UpdateId    int64         `json:"updateId"`
	UpdateTime  string        `json:"updateTime"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	Summary     string        `json:"summary"`
	CoverImg    string        `json:"cover_img"`
	ArticleTags []*ArticleTag `json:"articleTags"`
	CateInfo    *Category     `json:"cate_info"`
}

type CreateArticleRequest struct {
	Title string `json:"title"`
}

type SaveArticleRequest struct {
	ArticleId int64  `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Summary   string `json:"summary"`
	CoverImg  string `json:"cover_img"`
	Cate      int64  `json:"cate"`
}

type PublishArticleRequest struct {
	ArticleId int64  `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Summary   string `json:"summary"`
	CoverImg  string `json:"cover_img"`
	Cate      int64  `json:"cate"`
}

type ArticleTag struct {
	Id      int64  `json:"id"`
	TagName string `json:"tagName"`
}

type CreateArticleResponse struct {
	ArticleId int64
}

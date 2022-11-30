package vo

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/11/14 18:25
 **/
type TagsResponse struct {
	Tags       []*Tag      `json:"tags"`
	Categories []*Category `json:"categories"`
}

type Tag struct {
	Id       int64  `json:"id"`
	Code     string `json:"code"`
	Sequence int64  `json:"sequence"`
	Name     string `json:"name"`
}

type Category struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Sequence int64  `json:"sequence"`
}

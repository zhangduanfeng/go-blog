package model

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/11/14 18:06
 **/
type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Sequence   int64  `json:"sequence"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

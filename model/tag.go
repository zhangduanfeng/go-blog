package model

type Tag struct {
	Id         int64  `json:"id"`
	Type       int64  `json:"type"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Valid      int64  `json:"valid"`
	Sequence   int64  `json:"sequence"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

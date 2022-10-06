package model

type User struct {
	Id         int64  `json:"id"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Username   string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password   string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	RoleId     int64  `json:"role" validate:"required" label:"角色ID"`
	Valid      int    `json:"Valid" label:"是否有效0-有效1-无效"`
}

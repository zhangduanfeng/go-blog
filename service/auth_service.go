package service

import (
	"context"
	"errors"
	"go-blog/dal/db"
	"go-blog/model"
	"time"

)

/**
 * @Author huchao
 * @Description //
 * @Date 16:42 2022/2/12
 **/
func CreateUser(ctx context.Context, userName, passWord string) error {
	existUser, err := db.GetUserByUserName(ctx, userName)
	if err != nil {
		return err
	}
	if existUser.Id != 0 {
		return errors.New("已存在的用户")
	}

	var user = &model.User{
		Username: userName,
		Password: passWord,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := db.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}
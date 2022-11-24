package db

import (
	"context"
	"go-blog/model"
	"go-blog/store"
)

func GetUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	var user = &model.User{}
	err := store.DB.Debug().Where("user_name = ?", userName).Find(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return user, nil
		}
		return nil, err
	}
	return user, nil
}

func CreateUser(ctx context.Context, user *model.User) error {
	if err := store.DB.Debug().Table("user").Create(user).Error; err != nil {
		return err
	}
	return nil
}

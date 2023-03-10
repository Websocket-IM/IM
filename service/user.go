package service

import (
	"fmt"
	"ginchat/common"
	"ginchat/dao"
	"ginchat/model"
)

// 展示用户列表
func GetUserList() ([]*model.User, error) {
	users := make([]*model.User, 10)
	err := dao.GetUserList(users)

	for _, v := range users {
		fmt.Println(v)
	}
	return users, err
}

// 新增用户
func CreateUser(user *model.RegisterUserReq) error {
	if err := dao.CreateUser(user); err != nil {
		common.SugarLogger.Error("新增用户错误, err: %v ", err)
		fmt.Println(err, 11111111111111)
		return err
	}
	return nil
}

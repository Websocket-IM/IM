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

// 删除用户
func DeleteUser(user *model.User) error {
	if err := dao.DeleteUser(user); err != nil {
		common.SugarLogger.Error("删除用户错误，err: %v", err.Error())
		return err
	}
	return nil
}

// 修改用户资料
func UpadateUser(user *model.UpadateUserRep) error {
	fmt.Println(555555)
	if err := dao.UpadateUser(user); err != nil {
		common.SugarLogger.Error("修改用户资料错误： %v", err.Error())
		fmt.Println(err)
		return err
	}

	return nil
}

// 通过字段查询用户
func FindBy(field string, value interface{}) ([]model.User, error) {
	users, err := dao.FindBy(field, value)
	if err != nil {
		common.SugarLogger.Error("查询用户%v错误： %v", field, err)
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

// 用户登录
func LoginUser(rep *model.LoginUserRep) error {
	if err := dao.LoginUser(rep); err != nil {
		common.SugarLogger.Error("用户登录错误，err:%v", err)
		return err
	}
	return nil
}

// 手机号验证码登录
func LoginByPhoneCode(phoneCode *model.LoginByPhoneCode) (model.User, error) {
	err, user := dao.LoginByPhoneCode(phoneCode)
	if err != nil {
		common.SugarLogger.Error("用户短信验证错误，err:%v", err)
		fmt.Println(3333333)
		return model.User{}, err
	}
	return user, nil
}

// redis验证
func SaveCode(phone, code string) {
	err := dao.SaveCode(phone, code)
	if err != nil {
		common.SugarLogger.Error("redis缓存验证码错误，err:%v", err)

		panic(err)
	}
}

// 添加用户
func AddUser(user *model.User) error {
	if err := dao.AddUser(user); err != nil {
		fmt.Println(err)
		common.SugarLogger.Error("添加用户时出错：%v", err)
		return err
	}
	fmt.Println("添加用户成功")
	return nil
}

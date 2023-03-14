package dao

import (
	"errors"
	"fmt"
	"ginchat/common"
	"ginchat/model"
	"ginchat/utils"
	"math/rand"
	"time"
)

// 展示用户列表
func GetUserList(users []*model.User) error {
	result := common.DB.Find(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 新增用户
func CreateUser(rep *model.RegisterUserReq) error {
	salt := fmt.Sprintf("%08d", rand.Int31())
	user := model.User{
		Username:      rep.Username,
		Nickname:      utils.RandNickname(),
		Password:      utils.Md5Password(rep.Password, salt),
		Salt:          salt,
		LoginTime:     time.Now(),
		LoginOutTime:  time.Now(),
		HeartbeatTime: time.Now(),
	}
	res := common.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 删除用户
func DeleteUser(user *model.User) error {
	res := common.DB.Delete(user, "id = ?", user.ID)
	fmt.Println(user.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 修改用户资料
func UpadateUser(user *model.UpadateUserRep) error {
	fmt.Println(6666)
	res := common.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(model.User{Nickname: user.Nickname, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
	if res.Error != nil {
		return res.Error
	}
	fmt.Println(2222)
	return nil
}

// 通过字段查询用户
func FindBy(field string, value interface{}) ([]model.User, error) {
	var users []model.User
	if res := common.DB.Where(field+" = ?", value).Find(&users); res.Error != nil {
		return nil, res.Error

	} else {
		return users, nil
	}
}

// 用户登录
func LoginUser(rep *model.LoginUserRep) error {
	var user model.User
	res := common.DB.Where("username = ?", rep.Username).Select("salt").First(&user)
	fmt.Println(user)
	fmt.Println("salt:", user.Salt)
	if res.Error != nil {
		fmt.Println("err:", res.Error)
		return res.Error
	}
	var user2 model.User
	res = common.DB.Where("username = ? and password = ?", rep.Username, utils.Md5Password(rep.Password, user.Salt)).Find(&user2)
	if res.Error != nil || user2.ID == 0 {
		return errors.New("密码错误！")
	}
	return nil
}

// 手机号验证码登录
func LoginByPhoneCode(phoneCode *model.LoginByPhoneCode) (error, model.User) {
	value, err := common.RDB.Get(common.CTX, phoneCode.Phone).Result()
	if err != nil {
		return err, model.User{}
	}
	if value == "" {
		return errors.New("找不到验证码"), model.User{}
	}
	if value != phoneCode.Code {
		return errors.New("验证码错误！"), model.User{}
	}
	// 主动注册
	salt := fmt.Sprintf("%08d", rand.Int31())
	user := model.User{
		Phone:         phoneCode.Phone,
		Nickname:      utils.RandNickname(),
		Password:      utils.Md5Password("faweifafw", salt), // 默认随机密码
		Salt:          salt,
		LoginTime:     time.Now(),
		LoginOutTime:  time.Now(),
		HeartbeatTime: time.Now(),
	}
	res := common.DB.Create(&user)
	if res.Error != nil {
		return res.Error, model.User{}
	}
	return nil, user
}

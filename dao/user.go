package dao

import (
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

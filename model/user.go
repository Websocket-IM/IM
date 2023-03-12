package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username      string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

// 注册用户
type RegisterUserReq struct {
	Username   string `form:"username" json:"username"  comment:"用户名" validate:"required,min=6,max=20"`
	Password   string `form:"password" json:"password"   comment:"密码" validate:"required,min=6,max=20"`
	RePassword string `form:"repassword" json:"repassword"  comment:"确认密码" validate:"required,eqfield=Password"`
}

// 更新用户资料
type UpadateUserRep struct {
	ID     uint   `gorm:"primarykey"`
	Phone  string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email  string `valid:"email"`
	Avatar string //头像
}
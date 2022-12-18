package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	fmt.Println("用户输入的密码：", password)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL AvatarUrl 封面地址
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}

package service

// 上传文件服务

import (
	"fmt"
	"go-mall/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, id uint, name string) (filePath string, err error) {
	bId := strconv.Itoa(int(id))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"

	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	// 用户头像路径
	avatarPath := basePath + name + ".jpg" // todo : 提取file的后缀
	// 转换文件类型
	content, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		fmt.Println("写入本地文件")
		return
	}

	return "user" + bId + "/" + name + ".jpg", nil
}

func UploadProductToLocalStatic(file multipart.File, id uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(id))
	basePath := "." + conf.ProductPath + "boss" + bId + "/"

	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	// 用户头像路径
	productPath := basePath + productName + ".jpg" // todo : 提取file的后缀
	// 转换文件类型
	content, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		fmt.Println("写入本地文件")
		return
	}

	return "boss" + bId + "/" + productName + ".jpg", nil
}

// DirExistOrNot 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}

package service

import (
	"context"
	"go-mall/conf"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

// UserService 接受参数列表
type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	// 前端验证key
	Key string `json:"key" form:"key"`
}

type SendEmailService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// 1 绑定邮箱 2 解绑邮箱 3 修改密码
	OperationType uint `json:"Operation_type" form:"operation_type"`
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// Register 用户注册逻辑
func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足！",
		}
	}

	// 金额 10000  ---> 密文存储，对称加密操作
	util.Encrypt.SetKey(service.Key)

	// dao层编写
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExitUser
		return serializer.Response{
			Status: e.ErrorExitUser,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
		Email:    service.NickName,
		Avatar:   "avatar.jpg",
		Status:   model.Active,
		// 初始金额的加密
		Money: util.Encrypt.AesEncoding("10000"),
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登录
func (service UserService) Login(ctx context.Context) serializer.Response {
	//定义一个用户对象
	var user *model.User
	// 初始化响应状态码变量值
	code := e.Success
	// 获取userDao对象
	userDao := dao.NewUserDao(ctx)
	//检查用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil || !exist {
		code = e.ErrorExitUserNotFound
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在！请先注册",
		}
	}
	//检查密码
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: e.ErrorNotCompare,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}
	// 因为http是无状态的协议，所以用户登录过后，让其携带一个登录凭证 使用jwt
	// token 签发
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// Update 用户信息更新
func (service UserService) Update(ctx context.Context, uId uint) serializer.Response {
	// 定义一个用户对象
	var user *model.User
	// 初始化响应状态值
	code := e.Success
	// 获取userDao对象
	userDao := dao.NewUserDao(ctx)
	// 获取这个用户
	user, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: e.ErrorExitUser,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在！",
		}
	}

	// 修改昵称nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}

// Post 更新用户头像
func (service UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	// 获取连接对象
	userDao := dao.NewUserDao(ctx)
	//fmt.Println("准备查询用户信息……")
	// 先查询用户信息
	user, err = userDao.GetUserById(uId)
	//fmt.Println("查询出的用户信息：", user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: e.ErrorExitUserNotFound,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 保存图片到本地函数
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: e.Error,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新用户信息
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Send 发送邮件
func (service SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {

	code := e.Success
	var address string
	var notice *model.Notice // 模板通知
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)

	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 模板dao NoticeDao
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//发送方
	address = conf.ValidEmail + token
	// 邮件发送相关
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证邮箱
func (service ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	code := e.Success
	var userId uint
	var email string
	var password string
	var operationType uint

	// 验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	// 返回失败响应
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//成功之后，获取用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if operationType == 1 {
		// 绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 解绑邮箱
		user.Email = ""
	} else {
		// 更新密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			util.LogrusObj.Info("err: ", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	// 更新个人信息
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Show 展示用户金额
func (service ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)

	if err != nil {
		code = e.Error
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}

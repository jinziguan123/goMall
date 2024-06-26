package service

import (
	"context"
	"goMall/backend/config"
	"goMall/backend/consts"
	"goMall/backend/pkg/e"
	"goMall/backend/pkg/utils"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"
	"mime/multipart"
	"strings"
	"time"

	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
)

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
	Token string `form:"token" json:"token"`
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥错误",
		}
	}
	utils.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"), // 初始金额
	}
	// 加密密码
	if err = user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if config.UploadModel == consts.UploadModelOss {
		user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	} else {
		user.Avatar = "avatar.JPG"
	}
	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { // 如果查询不到，返回相应的错误
		logging.Info(err)
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// Update 用户修改信息
func (service UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.SUCCESS
	// 找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var user *model.User
	var err error

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var path string
	if config.UploadModel == consts.UploadModelLocal { // 兼容两种存储方式
		path, err = utils.UploadAvatarToLocalStatic(file, uId, user.UserName)
	} else {
		path, err = utils.UploadToQiNiu(file, fileSize)
	}
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}

	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// Send 发送邮件
func (service *SendEmailService) Send(ctx context.Context, id uint) serializer.Response {
	code := e.SUCCESS
	var address string
	var notice *model.Notice

	token, err := utils.GenerateEmailToken(id, service.OperationType, service.Email, service.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = config.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", config.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		logging.Info(err)
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证内容
func (service ValidEmailService) Valid(ctx context.Context) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS

	// 验证token
	if service.Token == "" {
		code = e.InvalidParams
	} else {
		claims, err := utils.ParseEmailToken(service.Token)
		if err != nil {
			logging.Info(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		// 1:绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 2：解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		// 3：修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 成功则返回用户的信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

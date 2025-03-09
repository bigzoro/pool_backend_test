package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"pool/common/constant"
	"pool/dao"
	"pool/forms"
	"pool/log"
	"pool/middlewares"
	"pool/models"
	"pool/utils"
	"time"
)

func EmailRegister(registerForm *forms.EmailRegisterForm) error {
	// 参数校验
	var newUser models.User
	newUser.Username = registerForm.Email
	newUser.Email = registerForm.Email

	// 加密用户的密码
	salt := utils.RandString(8)
	newUser.Salt = salt
	newUser.Password = utils.GenPassword(registerForm.Password, salt)
	// 检查验证码
	//_, err := global.RedisClient.Get(context.Background(), registerForm.Username).Result()
	//if err != nil {
	//	return err
	//}

	// 判断用户是否存在
	exist, err := dao.UserExist(newUser.Username)
	if err != nil {
		log.SystemLog().Warnf("[Register]: ", err)
		return err
	}
	if exist {
		return errors.New(constant.ErrorUserExist)
	}

	// todo 如果用户使用邀请码，就验证邀请码是否有效

	// todo 添加关系表

	// 邀请部分
	// 1. 为用户生成邀请码
	inviteCode, err := utils.GenerateInviteCode()
	if err != nil {
		log.SystemLog().Warnf("[Register]: ", err)
		return err
	}
	// todo: 判断邀请码是否存在

	newUser.InviteCode = inviteCode

	// 2. 判断用户是否使用邀请码
	if registerForm.InviteCode != "" {
		// 使用邀请码查询上一级用户 ID
		referrerUser, err := dao.GetUserByInviteCode(registerForm.InviteCode)
		if err != nil {
			log.SystemLog().Warnf("[Register]: ", err)
			return err
		}

		// 3. 如果使用，就绑定关系
		err = dao.AddReferral(int(referrerUser.ID), int(newUser.ID))
		if err != nil {
			log.SystemLog().Warnf("[Register]: ", err)
			return err
		}
	}

	// 添加用户
	err = dao.AddUser(&newUser)
	if err != nil {
		log.SystemLog().Warnf("[Register]: ", err)
		return err
	}

	// 添加到邀请表

	// 添加到邀请码表
	err = dao.AddInviteCode(newUser.ID, inviteCode)
	if err != nil {
		log.SystemLog().Warnf("[Register]: ", err)
		return err
	}

	// 设置 token
	return nil
}

func Login(loginForm *forms.EmailLoginForm) (*models.User, string, error) {

	var user *models.User
	var err error
	// 判断登录类型
	switch loginForm.LoginType {
	case constant.PasswordType:
		// 校验用户是否存在
		if len(loginForm.Username) == 0 {
			return nil, "", errors.New(constant.ErrorParameter)
		}
		user, err = dao.GetUserByEmail(loginForm.Username)
		if err != nil {
			log.SystemLog().Info("[Login]: ", err)
			return nil, "", errors.New(constant.ErrorInternal)
		}
		// todo: 可能一直不为空
		if user == nil {
			return nil, "", errors.New(constant.ErrorUserNotExist)
		}
		// 校验密码是否正确
		if loginForm.Password != user.Password {
			return nil, "", errors.New(constant.ErrorPassword)
		}
	case constant.MobileType:
		// 校验手机号是否正确
		// 校验用户是否存在
	case constant.EmailType:
		// 校验邮箱是否正确
		// 检测用户是否存在
		// 检测校验码是否正确
		// 放行
		if len(loginForm.Username) == 0 {
			return nil, "", errors.New(constant.ErrorParameter)
		}
		user, err = dao.GetUserByEmail(loginForm.Username)
		if err != nil {
			log.SystemLog().Info("[Login]: ", err)
			return nil, "", errors.New(constant.ErrorInternal)
		}
		// todo: 可能一直不为空
		if user == nil {
			return nil, "", errors.New(constant.ErrorUserNotExist)
		}
		// 校验密码是否正确
		if loginForm.Password != user.Password {
			return nil, "", errors.New(constant.ErrorPassword)
		}
	default:
	}

	// 生成 token
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:       uint(user.ID),
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "plum",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		log.SystemLog().Warnf("[Login]: ", err)
		return nil, "", nil
	}

	return user, token, nil
}

func UpdatePassword(params *forms.UpdatePasswordForm) error {
	// 获取用户的信息
	user, err := dao.QueryUserById(params.UserId)
	if err != nil {
		return err
	}

	// 生成新的密码
	password := utils.GenPassword(params.Password, user.Salt)

	err = dao.UpdatePassword(params.UserId, password)
	if err != nil {
		log.SystemLog().Warnf("[UpdatePassword]: %v", err)
		return err
	}

	return nil
}

func UpdateWithdrawPassword(params *forms.UpdateWithdrawPasswordForm) error {
	// 获取用户的信息
	user, err := dao.QueryUserById(params.UserId)
	if err != nil {
		return err
	}

	// 生成新的密码
	withdrawPassword := utils.GenPassword(params.WithdrawPassword, user.Salt)

	err = dao.UpdateWithdrawPassword(params.UserId, withdrawPassword)
	if err != nil {
		log.SystemLog().Warnf("[UpdatePassword]: %v", err)
		return err
	}

	return nil
}

func UpdateEmail(params *forms.UpdateEmailForm) error {
	err := dao.UpdateEmail(params.UserId, params.Email)
	if err != nil {
		log.SystemLog().Warnf("[UpdatePassword]: %v", err)
		return err
	}

	return nil
}

func GetUserInfo(userId int) (*models.User, error) {
	return dao.QueryUserById(userId)
}

package model

import "errors"

//根据逻辑业务自定义一些错误
var (
	Error_User_NotExists=errors.New("用户不存在")
	Error_User_Exists=errors.New("用户已存在")
	Error_User_pwd=errors.New("密码不正确")
)

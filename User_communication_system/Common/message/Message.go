package message

import (
	"communicate/Common/User"
)

//定义每个消息类型的绰号
const (
	LoginMesType         = "LoginMes"
	LoginRetType         = "LoginRetMes"
	RegisterMesType      = "RegisterMes"
	RegisterRetMesType   = "RegisterRerMes"
	NotifyUserStatusType = "NotifyUserStatus"
	SmsMesGroupType      = "SmsMesGroup"
	SmsOneMesType        = "SmsOneMes"
)

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// Message 这是总消息，就是客户端发给服务器的总消息
type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

// LoginMes 定义两个发送的消息..后面再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

// LoginRetMes  服务器给客户端返回的消息
type LoginRetMes struct {
	Code    int    `json:"code"`  //返回的一个id，状态码500表示该用户未注册，200表示登陆成功
	Error   string `json:"error"` //返回的错误信息，若没有则返回nil
	UserIds []int  //增加字段，保存用户id的一个切片，代表登录返回一堆user的id，使用户知道谁在线
}

// RegisterMes 注册消息
type RegisterMes struct {
	User User.User
}
type RegisterRetMes struct {
	Code  int    `json:"code"`  //400表示该用户已经占用，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

// NotifyUserStatusMes 为了配合服务器端推送通知别人状态变化信息，定义一个信息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"Status"`
}

// SmsGroupMes 增加一个群聊消息发送的对象
type SmsGroupMes struct {
	Content string `json:"content"`
	User    User.User
}

// SmsOneMes 增加一个私聊消息的对象
type SmsOneMes struct {
	Content string `json:"content"`
	SendUser    User.User
	ReceiverUser User.User
}

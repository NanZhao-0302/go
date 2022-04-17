package process

import (
	"communicate/Client/model"
	"communicate/Common/User"
	"communicate/Common/message"
	"fmt"
)

// 客户端要维护的map
var OnlineUsers = make(map[int]*User.User, 10)

//全局一个当前用户
var CurUser model.CurUser	//我们在登陆成功后，完成对CurUser的初始化,相当于新建一个人

// UpdateUserStatus 编写处理返回的NotifyUserStatus
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := OnlineUsers[notifyUserStatusMes.UserId]
	//原来没有这个
	if !ok {
		user = &User.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	//原来有
	user.UserStatus = notifyUserStatusMes.Status
	OnlineUsers[notifyUserStatusMes.UserId] = user
	showOnlineUsers()
}
//在客户端显示当前在线的用户
func showOnlineUsers() {
	//遍历onlineUsers完事
	fmt.Println("当前在线用户列表")
	for id,_:=range OnlineUsers{
		fmt.Println("用户id:\t",id)
	}
}
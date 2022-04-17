package process

import (
	utils2 "communicate/Client/utils"
	"communicate/Common/User"
	"communicate/Common/message"
	"communicate/Server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

// Login 登录函数来验证
//最好返回值不是真假，因为不能具体知道是什么错误
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//todo 需要开始定协议
	//fmt.Printf("id是：%d\n 密码是：%s",userId,userPwd)
	//return nil

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败,err=", err)
		return
	}
	defer conn.Close()
	//连接上了，准备发送消息
	//2.先定义一个消息
	var mes message.Message
	mes.Type = message.LoginMesType //消息类型为登录的消息，消息包中有定义过
	//3.创建loginmes结构体
	var loginmes message.LoginMes
	loginmes.UserId = userId
	loginmes.UserPwd = userPwd
	//4.将loginmes结构体序列化，然后再放入message结构体当中
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("序列化出错了,err=", err)
		return
	}
	mes.Data = string(data) //mes结构体中的data变为loginMes

	//5.mes可以进行序列化了！！
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化出错了,err=", err)
		return
	}
	err = utils2.WritePkg(conn, data)
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.Readpkg()
	if err != nil {
		fmt.Println("转换失败,err=", err)
		return
	}
	//将mes的data反序列化
	var loginRetMes message.LoginRetMes
	err = json.Unmarshal([]byte(mes.Data), &loginRetMes)
	if loginRetMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn=conn
		CurUser.User.UserId=userId
		CurUser.User.UserStatus=message.UserOnline
		//可以显示在线用户列表
		fmt.Println("当前用户列表如下：")
		for _, v := range loginRetMes.UserIds {
			fmt.Println("用户id:\t", v)
			//完成客户端 onlineUsers 的初始化
			user := &User.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[v] = user
		}
		fmt.Println()
		fmt.Println()
		//这里需要再客户端启动一个协程
		//该协程保持和服务器的通讯，如果服务器有数据推送个客户端
		//接受并显示在客户端的终端上
		go processServer(conn)

		//1.显示登录成功后的界面
		for {
			showMenu()
		}

	} else {
		fmt.Println(loginRetMes.Error)
	}
	return

}
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败,err=", err)
		return
	}
	defer conn.Close()
	//连接上了，准备发送消息
	//2.先定义一个消息
	var mes message.Message
	mes.Type = message.RegisterMesType //消息类型为注册的消息，消息包中有定义过

	//3.创建RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserName = userName
	registerMes.User.UserPwd = userPwd
	//4.将registerMes结构体序列化，然后再放入message结构体当中
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化出错了,err=", err)
		return
	}
	mes.Data = string(data) //mes结构体中的data变为registerMes
	//5.mes可以进行序列化了！！
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化出错了,err=", err)
		return
	}
	//这个时候，data就是我们要发送的序列化后的消息
	tf := utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送包错误，err=", err)
		return
	}
	mes, err = tf.Readpkg() //mes为RegisterRetMes
	if err != nil {
		fmt.Println("转换失败,err=", err)
		return
	}

	var RegisterResMes message.RegisterRetMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功，重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(RegisterResMes.Error)
		os.Exit(0)
	}
	return
}

package process

//关于用户的信息处理
import (
	"communicate/Common/message"
	"communicate/Server/model"
	"communicate/Server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	//加一个字段，表示conn是哪个用户
	UserId int
	Conn   net.Conn
}

// NotifyOtherOnlineUser 编写展示所有在线用户给别人的方法
//传过来的id通知其他所有人自己上线了
func(this *UserProcessor)NotifyOtherOnlineUser(userId int){
	//遍历用户列表map，一个个发送 每一个人的上线及其资料
	for id,up:=range userMgr.OnlineUsers{
		//过滤自己
		if id==userId{
			continue
		}
		//开始通知，单独写个方法通知
		up.NotifyMeToOthers(userId)
	}
}

// NotifyMeToOthers 单独通知的方法
func (this *UserProcessor)NotifyMeToOthers(userId int){
	//组装消息
	var mes message.Message
	mes.Type= message.NotifyUserStatusType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId=userId
	notifyUserStatusMes.Status=message.UserOnline

	//序列化消息
	data,err:=json.Marshal(notifyUserStatusMes)
	if err!=nil{
		fmt.Println("序列化失败，err=",err)
		return
	}
	mes.Data=string(data)
	//mes序列化
	data,err =json.Marshal(mes)
	if err!=nil{
		fmt.Println("序列化失败")
		return
	}
	tf:=utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err!=nil{
		fmt.Println("通知别人失败,err=",err)
		return
	}
}





// ServerProcessLogin 专门处理客户端发来的登录信息的请求
func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出它的data，并直接反序列化成loginmes
	var loginMes message.LoginMes
	//mes此时已经反序列化完成，但里面内容依旧是字符串，取出里面的data，将data变为byte切片并反序列化然后装入loginmes中
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	//声明返回的mes
	var resMes message.Message
	resMes.Type = message.LoginRetType
	//再声明一个只为登录而返回的loginResMes
	var loginRetMes message.LoginRetMes
	//1.使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.LoginCheck(loginMes.UserId, loginMes.UserPwd)
	//各种错误信息可以自己定义
	if err != nil {
		if err == model.Error_User_NotExists {
			loginRetMes.Code = 500
			loginRetMes.Error = err.Error()
		} else if err == model.Error_User_pwd {
			loginRetMes.Code = 403
			loginRetMes.Error = err.Error()
		} else {
			loginRetMes.Code = 505
			loginRetMes.Error = "服务器内部错误"
		}

	} else {
		loginRetMes.Code = 200
		//登陆成功，所以要把用户放入到userMgr当中
		//将登录成功的id赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在线用户，我上线了
		this.NotifyOtherOnlineUser(loginMes.UserId)
		//将当前的用户的id放入到loginRetMes中
		for id,_:= range userMgr.OnlineUsers {
			loginRetMes.UserIds= append(loginRetMes.UserIds, id)
		}
		fmt.Println(user, "登录成功")
	}

	//将返回的结构体LoginRetMes序列化，才能进行返回
	data, err := json.Marshal(loginRetMes)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
		return
	}
	//将这个序列化后的data放入resMes结构体中
	resMes.Data = string(data)
	//对resMes再进行序列化，然后准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	//write函数发送消息
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}







func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出它的data，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	//mes此时已经反序列化完成，但里面内容依旧是字符串，取出里面的data，将data变为byte切片并反序列化然后装入registerMes中
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	//声明返回的RetMes
	var resMes message.Message
	resMes.Type = message.RegisterRetMesType
	//再声明一个只为注册而返回的
	var RegisterRetMes message.RegisterRetMes
	//使用model.MyUserDao到redis去验证,完成注册
	err = model.MyUserDao.RegisterCheck(&registerMes.User)
	if err != nil {
		if err == model.Error_User_Exists {
			RegisterRetMes.Code = 505
			RegisterRetMes.Error = model.Error_User_Exists.Error()
		} else {
			RegisterRetMes.Code = 506
			RegisterRetMes.Error = "注册发生未知错误.."
		}
	} else {
		RegisterRetMes.Code = 200
	}
	//将返回的结构体LoginRetMes序列化，才能进行返回
	data, err := json.Marshal(RegisterRetMes)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
		return
	}
	//将这个序列化后的data放入resMes结构体中
	resMes.Data = string(data)
	//对resMes再进行序列化，然后准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	//write函数发送消息
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

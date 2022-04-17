package process

import (
	"communicate/Common/message"
	"communicate/Server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// Smsprocess 关于短信息（发送的消息）的处理
type Smsprocess struct {
}
//转发消息给所有客户端
func(this *Smsprocess)SendGroupMes(mes *message.Message){
	//遍历服务器端的在线用户（map）,将消息转发给每一个人
	//序列化一下mes，不拆开内容，直接转发给客户端们
	var smsMes message.SmsGroupMes
	err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	if err!=nil{
		fmt.Println("反序列化失败,err=",err)
	}
	data,err:=json.Marshal(mes)
	if err!=nil{
		fmt.Println("反序列化失败,err=",err)
	}
	for id,up:=range userMgr.OnlineUsers{
		//过滤掉自己，不要再发回给自己
		if id==smsMes.User.UserId{
			continue
		}
		this.SendMesOnline(data,up.Conn)
	}
}
func(this *Smsprocess)SendMesOnline(data []byte ,conn net.Conn){
	tf:=utils.Transfer{
		Conn: conn,
	}
	err:=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("转发消息失败,err=",err)
		return
	}
}
//点对点聊天
func(this *Smsprocess)SendOneMes(mes *message.Message){
	//遍历服务器端的在线用户（map）,将消息转发给每一个人
	//序列化一下mes，不拆开内容，直接转发给客户端们
	var smsMes message.SmsOneMes
	err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	if err!=nil{
		fmt.Println("反序列化失败,err=",err)
	}
	data,err:=json.Marshal(mes)
	if err!=nil{
		fmt.Println("反序列化失败,err=",err)
	}
	for id,up:=range userMgr.OnlineUsers{
		//找到你要发的选手
		if id==smsMes.ReceiverUser.UserId{
			this.SendMesOnline (data,up.Conn)
		}else{
			continue
		}

	}
}
package process

import (
	"communicate/Common/message"
	"communicate/Server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// SendGroupMes 发送群聊的消息
func (this *SmsProcess)SendGroupMes(content string)(err error){
	//1.创建一个message
	var mes message.Message
	mes.Type=message.SmsMesGroupType
	//2.创建SmsMes实例
	var smsMes message.SmsGroupMes
	smsMes.Content=content //内容
	smsMes.User.UserId=CurUser.User.UserId
	smsMes.User.UserStatus=CurUser.User.UserStatus
	//3.序列化smsMes
	data,err:=json.Marshal(smsMes)
	if err!=nil {
		fmt.Println("序列化失败，err=",err)
		return
	}
	mes.Data=string(data)
	//4.对mes序列化
	data,err =json.Marshal(mes)
	if err!=nil {
		fmt.Println("序列化失败，err=",err)
		return
	}
	//5.将mes发送给服务器
	tf:=utils.Transfer{
		Conn: CurUser.Conn,
	}
	//6.发送
	err=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("发送fail，err=",err)
		return
	}
	return
}

// SendOneMes 增加一个能点对点聊天的短信息处理
func (this *SmsProcess)SendOneMes(TargetId int,content string)(err error){
	//1.创建一个message
	var mes message.Message
	mes.Type=message.SmsOneMesType
	//2.创建SmsMes实例
	var smsOneMes message.SmsOneMes
	smsOneMes.Content=content //内容
	smsOneMes.SendUser.UserId=CurUser.User.UserId
	smsOneMes.SendUser.UserId=CurUser.User.UserId
	smsOneMes.SendUser.UserStatus=CurUser.User.UserStatus
	smsOneMes.ReceiverUser.UserId=TargetId
	//3.序列化smsMes
	data,err:=json.Marshal(smsOneMes)
	if err!=nil {
		fmt.Println("SmsOneMes序列化失败，err=",err)
		return
	}
	mes.Data=string(data)
	//4.对mes序列化
	data,err =json.Marshal(mes)
	if err!=nil {
		fmt.Println("序列化失败，err=",err)
		return
	}
	//5.将mes发送给服务器
	tf:=utils.Transfer{
		Conn: CurUser.Conn,
	}
	//6.发送
	err=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("发送fail，err=",err)
		return
	}
	return
}
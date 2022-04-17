package pro

//总控

import (
	"communicate/Common/message"
	process2 "communicate/Server/process"
	"communicate/Server/utils"
	"fmt"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// ServerProcessMes 根据客户端发送的消息不同来解析用哪个函数来处理
func (this *Processor)ServerProcessMes(mes *message.Message) (err error) {

	//看看是否能接收到客户端发送的群发消息
	fmt.Println("mes=",mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		up := process2.UserProcessor{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册逻辑
		up := process2.UserProcessor{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesGroupType:
		//处理群聊任务
		smsProcess:=&process2.Smsprocess{}
		smsProcess.SendGroupMes(mes)
	case message.SmsOneMesType:
		//处理点对点聊天
		smsProcess:=&process2.Smsprocess{}
		smsProcess.SendOneMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

// MainProcess 	循环处理信息，大的总控函数
func(this *Processor)MainProcess()(err error){
	for {
		tf:=utils.Transfer{
			Conn: this.Conn,
		}
		//这里我们将读取数据包，直接封装成一个函数readpkg(),返回message，err
		mes, err :=  tf.Readpkg()
		if err != nil {
			//可以自定义出错方式
			fmt.Println("read pkg 长度出错,err=",err)
			return err
		}
		//fmt.Println("message =", mes)
		err=this.ServerProcessMes(&mes)
		if err!=nil{
			return err
		}
	}
}
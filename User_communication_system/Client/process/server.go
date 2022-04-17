package process

import (
	"communicate/Common/message"
	"communicate/Server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func showMenu() {
		fmt.Println("------------恭喜",CurUser.User.UserId,"号登陆成功-----------------")
		fmt.Println("------------1.显示用户在线列表---------------")
		fmt.Println("------------2.发送群聊消息---------------------")
		fmt.Println("------------3.私聊---------------------")
		fmt.Println("------------4.退出系统---------------------")
		fmt.Println("请选择1-4")
		n := 0
		var content string
		var TargetId int
		smsprocess:= &SmsProcess{}
		fmt.Scanln(&n)
		switch n {
		case 1:
			showOnlineUsers()
		case 2:
			fmt.Println("请输入群聊消息：")
			fmt.Scanln(&content)
			smsprocess.SendGroupMes(content)
		case 3:
			//点对点聊天功能
			fmt.Println("请输入要发送的用户的id：")
			fmt.Scanln(&TargetId)
			fmt.Println("请输入要发送的内容：")
			fmt.Scanln(&content)
			// 进入方法处理
			smsprocess.SendOneMes(TargetId,content)
		case 4:
			fmt.Println("您退出了系统,拜拜~")
			os.Exit(0)
		default:
			fmt.Println("您输入的不正确")

		}

}
//和服务器端保持通讯
func processServer(conn net.Conn){
	//创建一个transfer实例，不停地读取服务器发送的消息
	tf:=utils.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("客户端正在等待读取服务器发来的消息")
		mes,err:=tf.Readpkg()
		if err!=nil{
			fmt.Println("read出错了，err=",err)
			return
		}
		//如果读取到消息，下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusType://有人上线了
		//1.拿到消息后，取出NotifyUserStatus
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
		//2. 加入到客户端维护的那个map[int]User里面去
			UpdateUserStatus(&notifyUserStatusMes)

		case message.SmsMesGroupType:    //有人群发消息了
			ShowGroupMes(&mes)
		case message.SmsOneMesType:		// 有人给你发了私聊消息
			ShowOneMes(&mes)
		default:
			fmt.Println("服务器返回了一个未知类型")
		}
	}
}

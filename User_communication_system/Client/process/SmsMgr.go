package process

import (
	"communicate/Common/message"
	"encoding/json"
	"fmt"
)

func ShowGroupMes(mes *message.Message){
	//显示即可
	//1.反序列化mes.data
	var smsMes message.SmsGroupMes
	 err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	 if err!=nil{
		 fmt.Println("反序列化失败,err=",err)
		 return
	 }
	 fmt.Printf("用户id(%d)\t对大家说：\t%s",
		 smsMes.User.UserId,smsMes.Content)
	 fmt.Println()
}
func ShowOneMes(mes *message.Message){
	var smsMes message.SmsOneMes
	err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	if err!=nil{
		fmt.Println("反序列化失败,err=",err)
		return
	}
	fmt.Printf("用户id(%d)\t对你说：\t%s",
		smsMes.SendUser.UserId,smsMes.Content)
	fmt.Println()

}

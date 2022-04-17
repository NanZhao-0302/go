package utils

//工具文件
import (
	"communicate/Common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// Transfer 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8064]byte	//传输时使用的缓冲

}


// WritePkg 发送。。。。，将其封装到一个写包的函数中
func (this *Transfer)WritePkg( data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, pkglen)
	//发送长度
	n, err := this.Conn.Write(bytes)
	if err != nil || n != 4 {
		fmt.Println("conn write(len(data)) 失败，err=", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if uint32(n) != pkglen || err != nil {
		fmt.Println("发送失败")
		return
	}
	return
}

//将客户端发来的消息进行反序列化的过程（读取）
func (this *Transfer)Readpkg()(mes message.Message, err error) {
	fmt.Println("等待读取客户端发送的数据..")
	_, err =this.Conn.Read(this.Buf[:4]) //n为实际上读了几个字节
	if err != nil {
		return
	}
	//根据bytes[:4]转成一个uint32类型，看看到底自己要读多少个
	var pkglen uint32
	//放入[]byte，然后返回一个uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkglen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn read fail,err=", err)
		return
	}
	//把pkglen反序列化成 -> message.Message
	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	return

}

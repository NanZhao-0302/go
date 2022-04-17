package utils

import (
	"communicate/Common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// WritePkg 发送。。。。，将其封装到一个写包的函数中
func WritePkg(conn net.Conn, data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, pkglen)
	//发送长度
	n, err := conn.Write(bytes)
	if err != nil || n != 4 {
		fmt.Println("conn write(len(data)) 失败，err=", err)
	}
	fmt.Println("客户端发送数据的长度ok", len(data))
	//发送data本身
	n, err = conn.Write(data)
	if uint32(n) != pkglen || err != nil {
		fmt.Println("发送失败")
		return
	}
	return
}

//将客户端发来的消息进行反序列化的过程（读取）
func Readpkg(conn net.Conn) (mes message.Message, err error) {
	bytes := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据..")
	_, err = conn.Read(bytes[:4]) //n为实际上读了几个字节
	if err != nil {
		fmt.Println("读错了，err=", err)
		return
	}
	//根据bytes[:4]转成一个uint32类型，看看到底自己要读多少个
	var pkglen uint32
	//放入[]byte，然后返回一个uint32
	pkglen = binary.BigEndian.Uint32(bytes[:4])
	//根据pkglen读取消息内容
	n, err := conn.Read(bytes[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn read fail,err=", err)
		return
	}
	//把pkglen反序列化成 -> message.Message
	err = json.Unmarshal(bytes[:pkglen], &mes)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	return

}

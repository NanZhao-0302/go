package utils

import (
	"customerManager/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8064]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据")

	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		return

	}

	var pkhLen uint32
	pkhLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkhLen])
	if uint32(n) != pkhLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
	}

	err = json.Unmarshal(this.Buf[:pkhLen], &mes)
	if err != nil {
		err = errors.New("read pkg body error")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}

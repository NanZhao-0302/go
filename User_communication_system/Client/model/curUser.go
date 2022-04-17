package model

import (
	"communicate/Common/User"
	"net"
)

// CurUser 客户端当中，很多地方使用，将其作为全局
type CurUser struct {
	Conn net.Conn
	User User.User
}

package User

//定义用户结构体
type User struct {
	UserId int	`json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"` //用户的状态（在线离线离开等）
}



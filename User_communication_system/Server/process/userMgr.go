package process

import "errors"

//管理所有用户，所有客户端
// UserMgr 因为userMgr在服务器只有一个，所以为全局变量
var (
	userMgr *UserMgr
)
type UserMgr struct {
	OnlineUsers map[int] *UserProcessor
}
//初始化userMgr
func init(){
	userMgr=&UserMgr{
		OnlineUsers: make(map[int]*UserProcessor,1024),
	}
}
func (this *UserMgr)AddOnlineUser(up *UserProcessor)  {
	this.OnlineUsers[up.UserId]=up
}
func (this *UserMgr)DeleteOnlineUser(userId int)  {
	delete(this.OnlineUsers,userId)
}

// GetAllOnlineUsers 得到所有用户
func (this *UserMgr)GetAllOnlineUsers()map[int] *UserProcessor{
	return this.OnlineUsers
}

// GetOnlineUserById 根据id来返回一个userProcess,找到一个用户可以传东西之类
func (this *UserMgr)GetOnlineUserById(userId int) (up *UserProcessor,err error){
	//如何从map中取出一个值，带检测方式
	up,ok:=this.OnlineUsers[userId]
	if !ok{
		//说明你要查找的用户的当前不在线
		errors.New("用户不存在")
		return
	}else{
		return
	}


}

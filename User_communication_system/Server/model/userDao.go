package model

import (
	"communicate/Common/User"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// MyUserDao 服务器启动时就初始化一个userDao实例
var (
	MyUserDao *UserDao
)

// UserDao 定义一个userDao结构体，完成对user这个结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 创建userDao的构造方法
func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	userdao = &UserDao{
		pool: pool,
	}
	return
}

//根据用户id，返回一个用户实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User.User, err error) {
	//通过id，去redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err == redis.ErrNil { //表示再user这个哈希中，我们没有找到对应的id
		err = Error_User_NotExists
		return
	}
	user = &User.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("反序列化出错,err=", err)
		return
	}
	return
}

//todo 完成登录的校验
//1.完成对用户的验证
//2.如果用户的id和pwd都正确则返回一个user对象
//3.如果用户id或pwd有错误，则返回一个err

func (this *UserDao) LoginCheck(userId int, userPwd string) (user *User.User, err error) {
	//先从userDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这时用户是获取到了，但是密码不一定正确
	if user.UserPwd != userPwd {
		err = Error_User_pwd
		return
	}
	return

}

// RegisterCheck  注册校验
func (this *UserDao) RegisterCheck(user *User.User) (err error) {
	//先从userDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	//这里应该是err==Nil才可以，因为如果找到了用户信息，说明用户已存在，注册失败
	if err == nil {
		err = Error_User_Exists
		return
	}
	//这时说明redis没有这个人，可以创建新用户
	data, err := json.Marshal(&user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err!=nil{
		fmt.Println("保存注册用户资料错误，err=",err)
		return
	}
	return
}

package zdpgo_docker

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

// 操作docker的结构体
type Docker struct {
	Host string // 主机地址
	Port int // SSH端口号
	Username string // 登录用户名
	Password string // 登录密码
	SSH *zdpgo_ssh.SSH // ssh连接对象
}

// 连接Docker
func (d *Docker)Connect(){
	if d.SSH == nil{
		ssh := zdpgo_ssh.New(d.Host, d.Username, d.Password, d.Port)
		ssh.Connect()
		d.SSH = ssh
	}
}
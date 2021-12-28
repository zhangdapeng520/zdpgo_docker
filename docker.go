package zdpgo_docker

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

// 操作docker的结构体
type Docker struct {
	host        string         // 主机地址
	port        int            // SSH端口号
	username    string         // 登录用户名
	password    string         // 登录密码
	ssh         *zdpgo_ssh.SSH // ssh连接对象
	log         *zdpgo_log.Log // 日志对象
	logFilePath string         // 日志文件路径
	debug       bool           // 是否为debug模式
}

// docker配置对象
type DockerConfig struct {
	Host        string // 主机地址
	Port        int    // SSH端口号
	Username    string // 登录用户名
	Password    string // 登录密码
	LogFilePath string // 日志文件路径
}

// 创建Docker
func New(config DockerConfig) *Docker {
	d := Docker{}

	// 初始化日志
	if config.LogFilePath == "" {
		d.log = zdpgo_log.New("zdpgo_docker.log")
	} else {
		d.logFilePath = config.LogFilePath
		d.log = zdpgo_log.New(config.LogFilePath)
	}

	// 初始化配置
	if config.Host == "" {
		d.log.Panic("host不能为空")
	} else {
		d.host = config.Host
	}

	if config.Port == 0 {
		d.log.Panic("port不能为空或0")
	} else {
		d.port = config.Port
	}

	if config.Username == "" {
		d.log.Panic("username不能为空")
	} else {
		d.username = config.Username
	}

	if config.Password == "" {
		d.log.Panic("password不能为空")
	} else {
		d.password = config.Password
	}

	// 初始化ssh连接
	ssh := zdpgo_ssh.New(config.Host, config.Username, config.Password, config.Port)
	ssh.Connect()
	d.ssh = ssh

	return &d
}

// 设置debug模式
func (d *Docker) SetDebug(debug bool) {
	d.debug = debug
}

// 判断是否为debug模式
func (d *Docker) IsDebug() bool {
	return d.debug
}

// 连接Docker
func (d *Docker) Connect() {
	if d.ssh == nil {
		ssh := zdpgo_ssh.New(d.host, d.username, d.password, d.port)
		ssh.Connect()
		d.ssh = ssh
	}
}

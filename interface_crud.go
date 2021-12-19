package zdpgo_docker

// 容器的CRUD接口
type CRUD interface{
	Create(image, name, port []string, volume []string) // 创建容器
	Delete() // 删除容器
	Update() // 修改容器
	Find() // 查询容器
	Log() // 查询容器日志
}
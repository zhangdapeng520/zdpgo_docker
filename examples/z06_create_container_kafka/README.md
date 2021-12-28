## 使用zdpgo_docker框架安装kafka
```go
package main

import (
	"fmt"
	
	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// 创建对象
	docker := zdpgo_docker.Docker{
		Host:     "192.168.18.101",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}
	
	// 创建zookeeper容器
	result, err := docker.CreateContainerZookeeperDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	
	// 创建kafka容器
	result, err = docker.CreateContainerKafka("kafka", "zookeeper", "192.168.18.101", 9092)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```

## 测试是否安装成功
启动kafka镜像生成容器
```shell
docker run -d --name kafka --publish 9092:9092 --link zookeeper --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 --env KAFKA_ADVERTISED_HOST_NAME=192.168.18.101 --env KAFKA_ADVERTISED_PORT=9092 --volume /etc/localtime:/etc/localtime wurstmeister/kafka
```

创建topic
```shell
docker exec kafka kafka-topics.sh --create --zookeeper 192.168.18.101:2181 --replication-factor 1 --partitions 1 --topic test
docker exec kafka kafka-topics.sh --list --zookeeper 192.168.18.101:2181
```

启动生产者
```shell
docker exec -it kafka kafka-console-producer.sh --broker-list 192.168.18.101:9092 --topic test
```

启动消费者
```shell
docker exec -it kafka kafka-console-consumer.sh --bootstrap-server 192.168.18.101:9092 --topic test --from-beginning
```

## 移除kafka
```go
package main

import (
	"fmt"
	
	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// 创建对象
	docker := zdpgo_docker.Docker{
		Host:     "192.168.18.101",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}
	
	// 删除容器
	result, err := docker.Remove("zookeeper")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	
	result, err = docker.Remove("kafka")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```
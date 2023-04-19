package main

import (
	"log"

	"github.com/surflabom/pag"
	"github.com/surflabom/pag/pkg/prometheus"
)

func main() {
	//新建一个 prometheus client
	config := prometheus.Config{
		BaseURL:          "http://127.0.0.1:9090",
		TargetConfigPath: "",
	}
	client, _ := pag.NewPrometheusClient(config)

	//开启一个 HTTP Server 为 prometheus 提供配置读取服务
	client.Server(":8090", "/remoteTargets")

	//初始化一个 Target
	testJobName := "testjob"
	TargetConfig := prometheus.TargetConfig{
		Targets: []string{},
		Labels: map[string]string{
			"Job": testJobName,
		},
	}

	//添加一个 Target
	client.AddTarget(TargetConfig)

	//获取所有的 Target
	targets, err := client.GetAllTarget()
	if err != nil {
		log.Println(err)
	}
	log.Println(targets)

	//为 Target 添加一个 EndPoint/Node
	EndPointAddress := "localhost:9090"
	endpoint := prometheus.EndPoint{
		JobName: testJobName,
		Address: EndPointAddress,
	}
	if err := client.AddTargetEndPoint(endpoint); err != nil {
		log.Println(err)
	}

	//为指定的 Target 删除一个 EndPoint/Node
	if err := client.DeleteTargetEndPoint(endpoint); err != nil {
		log.Println(err)
	}

}

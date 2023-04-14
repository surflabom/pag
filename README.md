## Grafana Prometheus Alertmanager Client SDK

***

这是一个为 Grafana Prometheus Alertmanager 修改配置的
SDK


### 安装
```go
go get  github.com/surflabom/pag
```


###   初始化客户端

***

```go
  package main

  import (
	"pag"
  )

func main()  {
	  // baseURL 是 grafana 服务位置 ,apiKey 是grafana 管理页面 - Configuration - API Keys 所生成的身份验证 Token
	  apiKey := "eyJrIjoiT3hacE5YclNSa29BWnJsemU5TjRndnZTSVZzWWdSY00iLCJuIjoiZGF3IiwiaWQiOj"
	  baseURL := "http://127.0.0.1:3000"
	  grafanaClient :=pag.NewGrafanaClient(baseURL, apiKey)

	  // baseURL 是 alertmanager 服务位置
	  baseURL := "http://127.0.0.1:9093"
	  alertClient := pag.NewAlertmanagerClient(baseURL)

	  
	  
	  // baseURL 是 prometheus 服务位置,TargetConfigPath 是远程配置读取所需文件的路径,如果为空,
	  // 则默认在程序根目录下生成 `configs.json`
	  baseURL := "http://127.0.0.1:9090"
	  TargetConfigPath:= "/dir/file"
	  prometheusClient := pag.NewPrometheusClient(baseURL, TargetConfigPath)
	  
	  // 开启一个 Web API 接口服务,利用 Prometheus 能够读取远程配置的功能 来修改所需的 Target 配置
	  // 修改 Prometheus 程序运行目录下的 prometheus.yml 中的 scrape_configs 配置项
	  //   - job_name: "remote_target_config"
      //       http_sd_configs:
	  //       - url: 'http://127.0.0.1:8090/targetConfig'
	  prometheusClient.Server(":8090","/targetConfig")
 }
```



### 示例

---

- 为 Alertmanager 创建 Silence 
```go
slience := alertmanger.Silence{
			CreatedBy: "zyl",
			Comment:   "create a silence",
			StartsAt:  time.Now().UTC(),
			EndsAt:    time.Now().UTC().Add(1 * time.Hour),
			Matchers: []alertmanger.Matcher{
				{
					Name:    "env",
					Value:   "production",
					IsRegex: false,
					IsEqual: true,
				},
			},
			Status: alertmanger.State{
				State:       "active",
				InhibitedBy: []string{},
				SilencedBy:  []string{},
			},
		}
		silences, err := alertClient.NewSilences(slience)
		
		log.Println(silences)
```

- 给 prometheus  添加一个 Target 远程配置和 一个 node
```go
       conf := prometheus.Config{
			Targets: []string{},
			Labels: map[string]string{
				"Job": "testJob",
			},
		}

	   if	err := prometheusClient.AddTarget(conf);err!=nil{
	        log.Println(err)
			return err
       }


	   endpoint := prometheus.EndPoint{
	         JobName: testJobName,
	         Address: EndPointAddress,
	   }
	   
	   if err := prometheusClient.AddTargetEndPoint(target);err!=nil{
	         log.Println(err)
	         return err
	   }
```

- 添加一个 DataSources
```go
   DS := grafana.AddDataSourceCommand{
     Name:      "Prometheus",
      Type:      "prometheus",
      Access:    "proxy",
      URL:       "http://127.0.0.1:9090",
      IsDefault: true,
      UID:       "6fDMIrLVz",
      OrgID:     1,
   }
   
   DataSources, err := grafanaClient.CreateDataSources(DS)
```













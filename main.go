package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"influxdbQuery/config"
	"influxdbQuery/handlers"
	"influxdbQuery/influxdb"
	"log"
	"net/http"
	"os"
)

func main() {

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	//添加路由
	ws.Route(ws.GET("/queryservice").To(startQuery))

	restful.Add(ws)
	fmt.Println("开始监听8080")
	//启动http服务
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func startQuery(req *restful.Request, resp *restful.Response) {
	// 判断是否在 Kubernetes 集群内
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		// 如果在 Kubernetes 集群内，修改 config 中的 InfluxDB URL
		config.Config.InfluxDBURL = "http://influxdb2.influxdb.svc.cluster.local:80"
	}

	// ✅ 读取配置
	config.LoadConfig()

	//调试输出
	fmt.Println(config.Config.InfluxDBURL)

	// ✅ 初始化 InfluxDB 客户端
	influxdb.NewClient()

	// ✅ 配置 HTTP 服务器
	//http.HandleFunc("/query", handlers.HandleQuery)

	handlers.HandleQuery(resp, req.Request)
	// ✅ 启动服务
	//port := ":8080"
	//fmt.Println("🚀 API 服务器已启动", port)
	//if err := http.ListenAndServe(port, nil); err != nil {
	//	fmt.Printf("❌ 服务器启动失败: %v\n", err)
	//}
}

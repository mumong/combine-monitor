package influxdb

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"influxdbQuery/config"
)

type Client struct {
	client   influxdb2.Client
	queryAPI api.QueryAPI
}

func NewClient() *Client {
	client := influxdb2.NewClient(config.Config.InfluxDBURL, config.Config.InfluxToken)
	queryAPI := client.QueryAPI(config.Config.InfluxOrg)

	return &Client{
		client:   client,
		queryAPI: queryAPI,
	}
}

func BuildFluxQuery(metric string, duration string) (string, error) {
	if metric == "" || duration == "" {
		return "", fmt.Errorf("metric 和 duration 是必填参数")
	}

	// 从配置读取参数
	bucket := config.Config.InfluxBucket
	org := config.Config.InfluxOrg
	url := config.Config.InfluxDBURL
	token := config.Config.InfluxToken

	log.Printf("📌 构建查询: metric=%s, duration=%s", metric, duration)

	query := fmt.Sprintf(`from(bucket: "%s") |> range(start: -%s) |> filter(fn: (r) => r._measurement == "prometheus" and r._field == "%s")`, bucket, duration, metric)

	log.Printf("📌 最终生成的查询: %s", query)
	log.Printf("📌 InfluxDB Host: %s, Org: %s, Token: %s", url, org, token[:5]+"***") // token只打印前5字符

	return query, nil
}

// ✅ 添加对 QueryData 方法的定义
func (c *Client) QueryData(query string) (*api.QueryTableResult, error) {
	result, err := c.queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

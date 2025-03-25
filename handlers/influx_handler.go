package handlers

import (
	"fmt"
	"influxdbQuery/influxdb"
	"log"
	"net/http"
)

var client = influxdb.NewClient()

func HandleQuery(w http.ResponseWriter, r *http.Request) {
	metric := r.URL.Query().Get("metric")
	duration := r.URL.Query().Get("duration")

	if metric == "" || duration == "" {
		http.Error(w, "metric 和 duration 都是必填参数", http.StatusBadRequest)
		return
	}

	// ✅ 构建 Flux 查询语句
	query, err := influxdb.BuildFluxQuery(metric, duration)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ 构建查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// ✅ 执行查询
	result, err := client.QueryData(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ 查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// ✅ 解析查询结果（动态获取所有字段，不做任何假设）
	var response string

	for result.Next() {
		record := result.Record()
		values := record.Values()

		response += "Record:\n"
		for key, value := range values {
			response += fmt.Sprintf("  %s: %v\n", key, value)
		}
		response += "\n"
	}

	if result.Err() != nil {
		http.Error(w, fmt.Sprintf("❌ 解析结果失败: %v", result.Err()), http.StatusInternalServerError)
		return
	}

	// ✅ 返回结果
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	log.Println("✅ 查询成功！")

}

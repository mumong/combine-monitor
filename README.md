# README.md
- en [English](README.EN.md)
- zh [Chinese](README.md)
- jp [Japanses](README.JP.md)

# 📦 combine-monitor

这是一个基于 **Prometheus + Telegraf + InfluxDB** 的个性化指标处理项目。  
项目通过 **Telegraf 抓取 Prometheus 的监控数据**，在采集阶段进行 **处理与优化**，再写入 InfluxDB 中以供查询使用。

最终封装为一个 **Pod 服务** 对外提供访问接口。  
**⚠️ 注意：需要提前部署并配置好 Prometheus、InfluxDB 和 Telegraf。**

---

## 🧩 项目架构与流程

1. **Telegraf** 从 Prometheus 抓取指标数据；
2. 在 Telegraf 内部通过插件进行部分字段裁剪、字段映射等个性化处理；
3. 处理后的数据被写入 **InfluxDB**；
4. 本项目中的服务（Pod）提供统一的接口，用于外部查询数据；
5. 主逻辑通过 `/queryservice` 路由暴露，结合 `config` 下的配置文件，将用户请求转化为 InfluxDB 查询命令；
6. 可查看 `telegraf.conf` 了解数据的采集与裁剪配置细节。

---

## 🌐 示例请求 & 查询效果

服务部署后，可以通过如下方式访问：

```bash
curl "http://192.168.3.76:30077/queryservice?metric=node_memory_MemFree_bytes&duration=1m"

Record:
  _stop: 2025-03-25 02:48:28.103879791 +0000 UTC
  _time: 2025-03-25 02:47:30 +0000 UTC
  _value: 1.2134998016e+10
  _field: node_memory_MemFree_bytes
  result: _result
  table: 0
  _start: 2025-03-25 02:47:28.103879791 +0000 UTC
  _measurement: prometheus
  host: telegraf-polling-service

Record:
  _measurement: prometheus
  host: telegraf-polling-service
  _start: 2025-03-25 02:47:28.103879791 +0000 UTC
  _stop: 2025-03-25 02:48:28.103879791 +0000 UTC
  _value: 1.212106752e+10
  _field: node_memory_MemFree_bytes
  result: _result
  table: 0
  _time: 2025-03-25 02:47:35 +0000 UTC
```
![image](https://github.com/user-attachments/assets/0221566f-9c6c-440c-b107-611817653acc)

## ✅**标准请求格式：**
`/queryservice?metric=<指标名>&duration=<时间范围>`

---
## 🛠️ 构建方式
1. 执行`make docker` 编译构建镜像。
2. 执行`bash ./tools/upload.sh` 将docker镜像上传到本地containerd仓库中
3. 执行`kubectl apply -f deployment.yaml` 应用文件部署服务，需要配置对应的rbac确保硬件资源的发现。




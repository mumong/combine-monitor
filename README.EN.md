# README.EN.md
- [English](README.EN.md)
- [Chinese](README.md)
- [Japanses](README.JP.md)

# 📦 combine-monitor

This is a personalized monitoring metrics processing project based on **Prometheus + Telegraf + InfluxDB**.  
The project uses **Telegraf to scrape metrics from Prometheus**, performs **custom processing and optimization**, and writes the results into **InfluxDB** for further querying.

The entire system is packaged as a **Kubernetes Pod service**, exposing a unified API interface.  
> ⚠️ Note: You must pre-deploy and configure **Prometheus**, **InfluxDB**, and **Telegraf** properly.

---

## 🧩 Architecture & Workflow

1. **Telegraf** scrapes metrics from Prometheus;
2. Inside Telegraf, custom plugins handle field filtering, renaming, and mapping;
3. Processed data is written into **InfluxDB**;
4. This project provides a **unified API** via a Pod to expose query capabilities;
5. The main logic is served under the `/queryservice` route, using configurations under the `config` directory to build InfluxDB queries;
6. Refer to `telegraf.conf` for detailed metric collection and filtering rules.

---

## 🌐 Example Request & Response

After deployment, you can query data using:

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

## ✅** Standard Query Format:：**
`/queryservice?metric=<指标名>&duration=<时间范围>`

---
## 🛠️ Build & Deploy
1. Run `make docker` to compile and build the Docker image.
2. Run `bash ./tools/upload.sh` to upload the Docker image to the local containerd repository.
3. Run `kubectl apply -f deployment.yaml` to deploy the service using the application file.



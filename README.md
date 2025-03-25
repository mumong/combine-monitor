# combine-monitor



这是一个基于prometheus，telegraf，influxdb的个性化处理监控指标的项目。
项目是通过telegraf抓取 prometheus的数据并且在telegraf进行一些个性化的处理和优化进而将结果存储到influxdb中用于后续的数据抓取。


整个项目可以封装为一个pod服务对外提供服务。需要提前部署好prometheus,influxdb。telegraf并且进行配置好。


通过queryservice 路由进行主函数的访问，将访问的逻辑封装为接口方便后续的扩展。
通过config下的配置文件与输入的请求命令拼接为一个infulxdb的查询命令，在对应的bucket中查询所需数据，可以查看考telegraf.conf中的配置里面具体的设置了如何将prometheus的数据存储到了influxdb并且进行一些裁剪处理。


以下是一个直观的示例运行起来后如何查询以及效果。
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


只截取部分数据，标准格式为 curl + url/queryservice?metric=需要查询的指标&duration=时间
----
## 构建方式
1. 执行`make docker` 构建编译镜像
2. 执行`bash ./tools/upload.sh` 将docker镜像保存到containerd镜像仓库中
3. 应用`deployment.yaml` 文件部署pod服务。




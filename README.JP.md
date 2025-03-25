# README.JP.md
- [English](README.EN.md)
- [Chinese](README.md)
- [Japanses](README.JP.md)

# 📦 combine-monitor

これは **Prometheus + Telegraf + InfluxDB** に基づくパーソナライズされたメトリクス処理プロジェクトです。  
プロジェクトでは、**Telegraf を使用して Prometheus の監視データを取得** し、収集段階で **処理と最適化** を行い、その後 InfluxDB に書き込んでクエリに利用します。

最終的に **Pod サービス** としてカプセル化され、外部アクセス用のインターフェースを提供します。  
**⚠️ 注意：Prometheus、InfluxDB、Telegraf を事前にデプロイおよび設定する必要があります。**

---

## 🧩 プロジェクトの構造とフロー

1. **Telegraf** が Prometheus からメトリクスデータを取得します；
2. Telegraf 内部でプラグインを使用して、一部のフィールドのトリミングやフィールドマッピングなどのパーソナライズされた処理を行います；
3. 処理されたデータが **InfluxDB** に書き込まれます；
4. このプロジェクトのサービス（Pod）が統一されたインターフェースを提供し、外部からのデータクエリを可能にします；
5. 主なロジックは `/queryservice` ルートを通じて公開され、`config` ディレクトリ内の設定ファイルと組み合わせて、ユーザーのリクエストを InfluxDB クエリコマンドに変換します；
6. データの収集とトリミング設定の詳細については、`telegraf.conf` を参照してください。

---

## 🌐 サンプルリクエストとクエリ結果

サービスがデプロイされた後、以下のようにアクセスできます：

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

## ✅**標準リクエスト形式：**
`/queryservice?metric=<メトリクス名>&duration=<時間範囲>`

---
## 🛠️ 構築方法
1. `make docker` を実行して Docker イメージをコンパイルおよびビルドします。
2. `bash ./tools/upload.sh` を実行して Docker イメージをローカルの containerd リポジトリにアップロードします。
3. `kubectl apply -f deployment.yaml` を実行してアプリケーションをデプロイします。


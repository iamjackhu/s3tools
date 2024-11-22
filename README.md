# S3Tools
S3 工具集

## 并发创建 buckets

### Config.yaml
```
s3Endpoint:
  endpoint: "10.0.1.88:9000"
  accessKey: "ZhAmNIEBEfTNpWGR"
  secretKey: "znN1kMoSlEh8NTy4h3skqakSrg45R2a8"
  region: "default"
bucketPrefile: "ppbk"
concurrent: 10
total: 100
```
这个配置表示将会由10个并发携程，总共创建 100 个bucket，bucket的前缀名为 “ppbk”

### 运行
```
make run
```
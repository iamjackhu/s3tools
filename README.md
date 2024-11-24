# S3Tools
S3 工具集

## 并发创建 buckets

### Config.yaml
```
s3Endpoint:
  endpoint: "10.0.1.71:9000"
  accessKey: "ZhAmNIEBEfTNpWGR"
  secretKey: "znN1kMoSlEh8NTy4h3skqakSrg45R2a8"
  region: "default"
operation: PUT  # PUT: create; DELETE: remove
bucketPrefile: "ppbk"
concurrent: 10
total: 100
```
这个配置表示将会由10个并发携程，总共创建 100 个bucket，bucket的前缀名为 “ppbk”

### 运行
```
make run
rm -rf ./build/bucketPerf
go mod tidy
go build -ldflags "-X main._version=''" -o ./build/bucketPerf *.go
./build/bucketPerf
2024/11/23 21:50:56 INFO Starting
2024/11/23 21:50:56 INFO [Operation 0] time(ms)=13
2024/11/23 21:50:56 INFO [Operation 1] time(ms)=14
2024/11/23 21:50:56 INFO [Operation 2] time(ms)=18
2024/11/23 21:50:56 INFO [Operation 3] time(ms)=18
2024/11/23 21:50:56 INFO [Operation 4] time(ms)=19
2024/11/23 21:50:56 INFO [Operation 5] time(ms)=19
2024/11/23 21:50:56 INFO [Operation 6] time(ms)=19
2024/11/23 21:50:56 INFO [Operation 7] time(ms)=19
2024/11/23 21:50:56 INFO [Operation 8] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 9] time(ms)=21
2024/11/23 21:50:56 INFO [Operation 10] time(ms)=21
2024/11/23 21:50:56 INFO [Operation 11] time(ms)=11
2024/11/23 21:50:56 INFO [Operation 12] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 13] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 14] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 15] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 16] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 17] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 18] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 19] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 20] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 21] time(ms)=12
2024/11/23 21:50:56 INFO [Operation 22] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 23] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 24] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 25] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 26] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 27] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 28] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 29] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 30] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 31] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 32] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 33] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 34] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 35] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 36] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 37] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 38] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 39] time(ms)=4
2024/11/23 21:50:56 INFO [Operation 40] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 41] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 42] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 43] time(ms)=4
2024/11/23 21:50:56 INFO [Operation 44] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 45] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 46] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 47] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 48] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 49] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 50] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 51] time(ms)=5
2024/11/23 21:50:56 INFO [Operation 52] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 53] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 54] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 55] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 56] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 57] time(ms)=4
2024/11/23 21:50:56 INFO [Operation 58] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 59] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 60] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 61] time(ms)=13
2024/11/23 21:50:56 INFO [Operation 62] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 63] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 64] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 65] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 66] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 67] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 68] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 69] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 70] time(ms)=11
2024/11/23 21:50:56 INFO [Operation 71] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 72] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 73] time(ms)=10
2024/11/23 21:50:56 INFO [Operation 74] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 75] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 76] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 77] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 78] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 79] time(ms)=4
2024/11/23 21:50:56 INFO [Operation 80] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 81] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 82] time(ms)=11
2024/11/23 21:50:56 INFO [Operation 83] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 84] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 85] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 86] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 87] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 88] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 89] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 90] time(ms)=11
2024/11/23 21:50:56 INFO [Operation 91] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 92] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 93] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 94] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 95] time(ms)=7
2024/11/23 21:50:56 INFO [Operation 96] time(ms)=9
2024/11/23 21:50:56 INFO [Operation 97] time(ms)=8
2024/11/23 21:50:56 INFO [Operation 98] time(ms)=6
2024/11/23 21:50:56 INFO [Operation 99] time(ms)=8
2024/11/23 21:50:56 INFO Completed
```
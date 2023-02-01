# zwzmicro

### macOS安装Go

https://dl.google.com/go/go1.15.1.darwin-amd64.pkg

### api语法

https://go-zero.dev/cn/docs/design/grammar?_highlight=api

### proto语法

https://protobuf.dev/programming-guides/proto3/

### yaml语法

https://go-zero.dev/cn/docs/configuration/api

### 宿主机器环境

```bash
go env -w GO111MODULE="on"
go env -w GOPROXY=https://goproxy.cn
go env -w GOMODCACHE=$GOPATH/pkg/mod

# Go 1.15 及之前版本
#GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest
# Go 1.16 及以后版本
go install github.com/zeromicro/go-zero/tools/goctl@latest

go install github.com/zeromicro/goctl-swagger@latest

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

### 容器环境

```bash
docker-compose -f docker-compose.yml -f docker-compose-env.yml up -d

```

### 生成代码

```bash
source deploy/sql/*.sql
sh deploy/script/*/*.sh
vim deploy/nginx/conf.d/looklook-gateway.conf
vim deploy/prometheus/server/prometheus.yml
```

### 启动项目

```bash
sh main.sh

```

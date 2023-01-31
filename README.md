### macOS安装Go

https://dl.google.com/go/go1.15.1.darwin-amd64.pkg

### api语法介绍

https://go-zero.dev/cn/docs/design/grammar?_highlight=api

### proto语法介绍

https://protobuf.dev/programming-guides/proto3/

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

### 运行容器

```bash
/usr/local/bin/docker-compose -f /Users/zhouwenzhe/src/go-zero-looklook/docker-compose.yml -f /Users/zhouwenzhe/src/go-zero-looklook/docker-compose-env.yml up -d
```

### 生成代码

```bash
sh genApi.sh
sh genRpc.sh
sh genSwagger.sh
sh genModel.sh

```

```shell
protoc --go_out=./helloworld --go-grpc_out=./helloworld ./helloworld/helloworld.proto
go run main.go
```
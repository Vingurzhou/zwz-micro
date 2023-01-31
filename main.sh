go run app/index/cmd/api/index.go -f app/index/cmd/api/etc/index-api.yaml &
go run app/greet/cmd/api/greet.go -f app/greet/cmd/api/etc/greet-api.yaml &
#go run app/usercenter/cmd/rpc/greet.go -f app/usercenter/cmd/rpc/etc/greet-api.yaml &
go run app/usercenter/cmd/rpc/usercenter.go -f app/usercenter/cmd/rpc/etc/usercenter.yaml &

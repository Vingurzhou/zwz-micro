go run app/index/cmd/api/index.go -f app/index/cmd/api/etc/index-api.yaml &
go run app/greet/cmd/api/greet.go -f app/greet/cmd/api/etc/greet-api.yaml &
go run app/usercenter/cmd/api/usercenter.go -f app/usercenter/cmd/api/etc/usercenter.yaml &
go run app/usercenter/cmd/rpc/usercenter.go -f app/usercenter/cmd/rpc/etc/usercenter.yaml &

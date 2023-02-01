service=user
mkdir -p "../../../app/${service}/cmd/rpc/pb"
cd ../../../app/${service}/cmd/rpc/pb
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
sed -i "" 's/,omitempty//g' *.pb.go

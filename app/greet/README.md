```shell
goctl api go -api *.api -dir "./" --style=gozero --home="../../deploy/goctl/1.3.4"
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api *.api -dir "./"
rm Dockerfile
goctl docker -go greet.go
rm k8s.yaml
goctl kube deploy -name greet1 -namespace greet-api -image ./ -o k8s.yaml -port 1001
kubectl create namespace greet-api
kubectl apply -f k8s.yaml
go run greet.go
```
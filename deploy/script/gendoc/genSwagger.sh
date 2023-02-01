service=index
mkdir -p "../../../app/${service}/cmd/api/desc"
cd ../../../app/${service}/cmd/api/desc
goctl api plugin -plugin goctl-swagger="swagger -filename ${service}.json" -api ${service}.api -dir .
#goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.2 -basepath /api" -api user.api -dir .
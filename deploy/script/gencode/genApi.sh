service=greet
mkdir -p "../../../app/${service}/cmd/api/desc"
cd ../../../app/${service}/cmd/api/desc
goctl api go -api *.api -dir ../ --style=gozero --home="../../../deploy/goctl/1.3.4"

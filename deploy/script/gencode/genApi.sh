service=search
mkdir -p "../../../app/${service}/cmd/api/desc"
cd ../../../app/${service}/cmd/api/desc
goctl api go -api *.api -dir ../ --style=goZero

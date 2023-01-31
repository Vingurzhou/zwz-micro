service=usercenter
mkdir -p "../../../app/${service}/model"
tables=user
modeldir=../../../app/${service}/model
host=127.0.0.1
port=3307
dbname=looklook_${service}
username=root
passwd=PXDN93VRKUm8TeE7
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" -cache=true --style=goZero

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	// 微信小程序配置
	WxMiniConf struct {
		AppId  string `json:"AppId"`  //微信appId
		Secret string `json:"Secret"` //微信secret
	}
	UsercenterRpcConf zrpc.RpcClientConf
}

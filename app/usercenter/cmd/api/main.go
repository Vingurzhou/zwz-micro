package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os/exec"
)

type Response struct {
	Code int `json:"code"`
	//Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	CodeSuccess      = 20000 // 请求成功
	CodeInvalidParam = 1000  // 请求参数错误
	CodeServerBusy   = 9999  // 服务器繁忙
)

var (
	MsgSuccess      = "请求成功"
	MsgInvalidParam = "请求参数错误"
	MsgServerBusy   = "服务器繁忙"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type InfoResp struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
type LoginResp struct {
	Token string `json:"token"`
}

func main() {

	r := gin.Default()
	r.POST("/vue-admin-template/user/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Code: CodeSuccess,
			Data: &LoginResp{"admin-token"},
		})
	})

	r.GET("/vue-admin-template/user/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Code: CodeSuccess,
			Data: &InfoResp{
				[]string{"admin"},
				"I am a super administrator",
				"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
				"Super Admin",
			},
		})
	})

	r.POST("/vue-admin-template/user/logout", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Code: CodeSuccess,
			Data: "success",
		})
	})
	//
	cmd := exec.Command("go", "run", "/Users/zhouwenzhe/src/zwz-micro/app/usercenter/cmd/script/main.go")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			var buf [1024]byte
			n, err := stdout.Read(buf[:])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			//fmt.Print(string(buf[:n]))
			err = conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	})
	r.Run(":1002") // listen and serve on 0.0.0.0:8080
}

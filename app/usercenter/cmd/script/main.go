package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Unix())
		time.Sleep(time.Second)
	}
}

package main

import (
	"fiber/utils"
	"fmt"
)

func main() {
	utils.Init()
	fmt.Println(utils.Env.Test)
	// go server.Server()
	// webview.Webview()
}

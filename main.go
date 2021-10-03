package main

import (
	"fiber/server"
	"fiber/utils"
	"fiber/webview"
)

func main() {
	utils.Init()
	go server.Server()
	webview.Webview()
}

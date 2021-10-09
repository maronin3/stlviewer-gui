package main

import (
	"fiber/server"
	"fiber/utils"
	"fiber/webview"
)

func main() {
	// 0. application path
	utils.Path()
	// 1. system logger
	utils.Logger()
	// 2. initalize enviorment
	utils.Init()
	// 3. webserver
	go server.Server()
	// 4. webview
	webview.Webview()
}

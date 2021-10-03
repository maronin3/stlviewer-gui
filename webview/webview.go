package webview

import (
	"fiber/utils"

	"github.com/webview/webview"
)

var (
	width  = utils.Env.Width
	height = utils.Env.Height
)

func Webview() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(width, height, webview.HintNone)
	w.Navigate("http://localhost:3000")
	w.Run()
}

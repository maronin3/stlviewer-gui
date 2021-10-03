package utils

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

// Default Config
type Config struct {
	Width  int `default:"800"`
	Height int `default:"600"`
}

var Env Config

// InItialize
func Init() {
	dirpath := "/mnt/e/Dev/Projects/STLViewer/stlviewer-onpremiss"
	// dirpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cfg, err := ini.Load(fmt.Sprintf("%v/server.ini", dirpath))
	if err != nil {
		log.Println(err)
	}

	// cfg.Section("Test").Key("VALUE").SetValue("world")
	// cfg.Section("WebView").Key("Heigh").Int()
	// result, _ :=cfg.Section("WebView").Key("Width").Int()

	Env = Config{
		Width:  800,
		Height: 600,
	}

	cfg.SaveTo(fmt.Sprintf("%v/server.ini", dirpath))
}

package utils

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

// Default Config
type Config struct {
	Test string `default:"hello"`
}

var Env Config

// InItialize
func Init() {
	dirpath := "/home/s-kim/dev/Template/golang/fiber"
	// dirpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cfg, err := ini.Load(fmt.Sprintf("%v/server.ini", dirpath))
	if err != nil {
		log.Println(err)
	}

	cfg.Section("Test").Key("VALUE").SetValue("world")

	Env = Config{
		Test: cfg.Section("Test").Key("VALUE").String(),
	}

	cfg.SaveTo(fmt.Sprintf("%v/server.ini", dirpath))
}

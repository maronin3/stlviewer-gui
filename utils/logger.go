package utils

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func Logger() {
	f, err := os.OpenFile(fmt.Sprintf("%v/system.log", DirPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	logger = log.New(f, "", log.LstdFlags)
}

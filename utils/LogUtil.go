package utils

import (
	"fmt"
	"log"
	"time"
)

func Error(msg string) {
	log.Fatal("Error: " + msg + " --- " + time.Now().String())
}

func Info(msg string) {
	fmt.Println("Info: " + msg + " --- " + time.Now().String())
}

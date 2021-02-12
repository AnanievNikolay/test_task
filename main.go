package main

import (
	"log"
	"os"

	"github.com/AnanievNikolay/test_task/app/file"
	"github.com/AnanievNikolay/test_task/app/service"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	initLogWrapper()
	serviceInstanse := service.New()
	serviceInstanse.Start()
}

func initLogWrapper() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(&lumberjack.Logger{
		Filename:   file.NewPath(os.Args[0], "/logs/console.log").Abs(),
		MaxSize:    500, // megabytes
		MaxBackups: 10,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	})
}

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/AnanievNikolay/test_task/app/servers/websocket"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	InitLogWrapper()
	pool := websocket.NewPool()
	server := websocket.New("", 80, pool.ConnectChannel(), pool.DisconnectChannel())
	go pool.Listen()
	server.Run()
}

//InitLogWrapper init log wrapper
func InitLogWrapper() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(&lumberjack.Logger{
		Filename:   GetAbsFilePath(os.Args[0], "/logs/console.log"),
		MaxSize:    500, // megabytes
		MaxBackups: 10,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	})
}

//GetAbsFilePath returns absolute file path
func GetAbsFilePath(dir, fileDest string) string {
	dir = filepath.Dir(dir)
	absPath, _ := filepath.Abs(dir + fileDest)
	return absPath
}

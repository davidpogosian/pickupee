package logger

import (
	"log"
	"os"
)

var std = log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)

func Info(msg string, v ...any) {
	std.Printf("INFO: "+msg, v...)
}

func Error(msg string, v ...any) {
	std.Printf("ERROR: "+msg, v...)
}

func Panicf(msg string, v ...any) {
	std.Panicf("PANIC: "+msg, v...)
}

func Fatalf(msg string, v ...any) {
	std.Fatalf("FATAL: "+msg, v...)
}

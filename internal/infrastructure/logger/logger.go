package logger

import "log"

type Logger struct{}

func New() *Logger { return &Logger{} }

func (l *Logger) Info(msg string, kv ...any) {
	log.Println(append([]any{"INFO:", msg}, kv...)...)
}

func (l *Logger) Error(msg string, kv ...any) {
	log.Println(append([]any{"ERROR:", msg}, kv...)...)
}

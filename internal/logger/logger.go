package logger

import (
	"log/slog"
	"os"
)

const (
	lvlLocal = "local"
	lvlDev   = "dev"
	lvlProd  = "prod"
)

//type LoggerWrapper struct{ *slog.Logger }
//
//func (l *LoggerWrapper) Info(msg string, fields ...slog.Attr) {
//	l.Logger.Info(msg, fields...)
//}
//
//func (l *LoggerWrapper) Debug(msg string, fields ...slog.Attr) {
//	l.Logger.Debug(msg, fields...)
//}
//
//func (l *LoggerWrapper) Error(msg string, fields ...slog.Attr) {
//	l.Logger.Error(msg, fields...)
//}
//
//func (l *LoggerWrapper) Warn(msg string, fields ...slog.Attr) {
//	l.Logger.Warn(msg, fields...)
//}
//
//func (l *LoggerWrapper) Fatal(msg string, fields ...slog.Attr) {
//	l.Logger.Fatal(msg, fields...)
//}
//
//var log *slog.Logger

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case lvlLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case lvlDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case lvlProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

//func MustInit() {
//	log = New(os.Getenv("LOG_LEVEL"))
//}

/*
Допилить до паттерна синглтон
*/

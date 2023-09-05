package logging

import (
	"fmt"
	"log/slog"
	"os"
)

type Enricher struct {
	ClientId string
}

type logger interface {
	Info(l slog.Logger, msg string, args ...any)
}

type Logger struct {
	l *slog.Logger
}

func (h Logger) Info(msg string, args ...any) {

	//for _, a := range args {
	//	msg = fmt.Sprintf(msg, a)
	//}

	//values := []interface{}{args}

	//msg = fmt.Sprintf(msg, args...)
	//fmt.Println(fmt.Sprintf(msg, args))

	//h.l.LogAttrs(nil, slog.LevelInfo, msg)

	//h.l.Info(msg)

	h.print(msg, args...)
}

func (h Logger) print(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	h.l.LogAttrs(nil, slog.LevelInfo, msg)

}

func NewLogger(args ...any) *Logger {

	handler := slog.NewJSONHandler(os.Stdout, nil)

	l := slog.New(handler)
	return &Logger{l}
}

package log

import (
	"fmt"
	"io"
)

type Logger interface {
	Write(msg string)
}

var log Logger = nil

type logger struct {
	w io.Writer
}

func Init(w io.Writer) {
	log = &logger{w: w}
}

func (l *logger) Write(msg string) {
	fmt.Fprintf(l.w, msg)
}

func Write(msg string) {
	log.Write(msg)
}

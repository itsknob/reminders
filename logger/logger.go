package logger

import (
	"fmt"
)

type Logger struct {
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	fmt.Printf("ERROR: "+f, v...)
}

func (l *Logger) Infof(f string, v ...interface{}) {
	fmt.Printf("INFO: "+f, v...)
}

func (l *Logger) Warningf(f string, v ...interface{}) {
	fmt.Printf("WARNING: "+f, v...)
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	fmt.Printf("DEBUG: "+f, v...)
}

package logs

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
}

var defaultWriter io.Writer = os.Stdout

func NewLogger() *Logger {
	return &Logger{
		InfoLog: log.New(defaultWriter, "INFO\t", log.Ldate|log.Ltime),
		WarningLog: log.New(defaultWriter, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(defaultWriter, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

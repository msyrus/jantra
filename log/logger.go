package log

import (
	"log"
	"os"
)

// Logger interface is the logger interface for the server
type Logger interface {
	Print(...interface{})
	Println(...interface{})
	Printf(string, ...interface{})
}

// DefaultLogger logs to stdout
var DefaultLogger Logger = log.New(os.Stdout, "", log.LstdFlags)

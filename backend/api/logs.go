package api

import (
	"log"
)

type logger struct {
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
}

var Logger = logger{}

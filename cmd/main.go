package main

import (
	"github.com/sgoodliff/hello/pkg/hello"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	goodbye := hello.Going()
	hello := hello.Whatsup()

	log.WithFields(log.Fields{
		"Hello":   hello,
		"Goodbye": goodbye,
	}).Debug("A walrus appears")
}

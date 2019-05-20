package main

import (
	"github.com/sgoodliff/hello/pkg/hello"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	hello := hello.Hello()
	goodbye := hello.Bye()

	log.WithFields(log.Fields{
		"Hello":   hello,
		"Goodbye": goodbye,
	}).Debug("A walrus appears")
}

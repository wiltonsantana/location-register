package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wiltonsantana/hpe-location/hpe-location-register/internal/config"
	"github.com/wiltonsantana/hpe-location/hpe-location-register/pkg/logging"
)

func monitorSignals(sigs chan os.Signal, quit chan bool, logger logging.Logger) {
	signal := <-sigs
	logger.Infof("signal %s received", signal)
	quit <- true
}

func main() {
	config := config.Load()
	logrus := logging.NewLogrus(config.Logger.Level, config.Logger.Syslog)

	logger := logrus.Get("Main")
	logger.Info("starting HPE Location Service")

	sigs := make(chan os.Signal, 1)
	quit := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go monitorSignals(sigs, quit, logger)
}

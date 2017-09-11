package logs

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	logger, err := seelog.LoggerFromConfigAsFile("log.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

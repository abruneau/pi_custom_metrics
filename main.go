package main

import (
	"flag"
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/abruneau/pi_custom_metrics/config"
	log "github.com/abruneau/pi_custom_metrics/logwrapper"
	"github.com/abruneau/pi_custom_metrics/system"
)

var (
	confFilePath *string
)

func init() {
	confFilePath = flag.String("conf", "", "path to directory containing config.yaml")
	flag.Parse()
}

func main() {

	err := config.SetupConfig(*confFilePath, "config", false)

	if err != nil {
		fmt.Println(err)
	}

	var standardLogger = log.NewLogger()

	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		standardLogger.Error(err)
	}

	standardLogger.Info("StatsD client started")
	system.CollectAllMetrics(c, standardLogger)

}

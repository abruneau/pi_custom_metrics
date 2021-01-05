package system

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	li "github.com/abruneau/pi_custom_metrics/logwrapper"
)

type metric struct {
	name, cmd string
}

func collectMetric(m metric, client *statsd.Client, log *li.StandardLogger) error {
	cmd := exec.Command("bash", "-c", m.cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%v: %s", err, stderr.String())
	}

	value, err := strconv.ParseFloat(strings.TrimSuffix(out.String(), "\n"), 64)
	if err != nil {
		return err
	}
	log.Debugf("%s: %v", m.name, value)
	return client.Gauge(fmt.Sprintf("system.%s", m.name), value, nil, 1)
}

// CollectAllMetrics collects system metrics every 15 seconds and send them throught DogStatsD
func CollectAllMetrics(client *statsd.Client, log *li.StandardLogger) {

	var metrics []metric
	metrics = append(metrics, metric{name: "gpu.temperature", cmd: "vcgencmd measure_temp | egrep -o '[0-9]*\\.[0-9]*'"})
	metrics = append(metrics, metric{name: "cpu.temperature", cmd: "cat /sys/class/thermal/thermal_zone0/temp | awk 'END {print $1/1000}'"})
	metrics = append(metrics, metric{name: "threads", cmd: "ps -eo nlwp | tail -n +2 | awk '{ num_threads += $1 } END { print num_threads }'"})
	metrics = append(metrics, metric{name: "processes", cmd: "ps axu | wc -l"})

	for range time.Tick(15 * time.Second) {
		log.Info("Starting metric collection")
		for _, m := range metrics {
			err := collectMetric(m, client, log)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

package metrics

import (
	"os"
	"path"

	prom "github.com/prometheus/client_golang/prometheus"
)

var defaultMetrics *LogMetrics

func init() {
	defaultMetrics = newLogMetrics(path.Base(os.Args[0]))
}

// GetLogMetrics returns the default LogMetrics.
func GetLogMetrics() *LogMetrics {
	return defaultMetrics
}

// LogMetrics --
type LogMetrics struct {
	entryCounter *prom.CounterVec
	procName     string
}

// NewLogMetrics creates a LogMetrics.
func newLogMetrics(proc string) *LogMetrics {
	c := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "_log_entry_total",
			Help: "Total number of output entries, regardless of level",
		},
		[]string{"proc", "level"},
	)
	prom.MustRegister(c)
	return &LogMetrics{
		entryCounter: c,
		procName:     proc,
	}
}

// LogEntry increase the number of log entries.
func (m *LogMetrics) LogEntry(level string) {
	m.entryCounter.WithLabelValues(m.procName, level).Inc()
}

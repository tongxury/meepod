package metrics

import "testing"

func TestLogMetrics(t *testing.T) {
	m := GetLogMetrics()
	m.LogEntry("info")
}

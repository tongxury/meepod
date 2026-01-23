package syslog

import (
	"log/syslog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPriority(t *testing.T) {
	testCases := []struct {
		facility string
		priority syslog.Priority
	}{
		{facility: "", priority: syslog.LOG_LOCAL5},
		{facility: "local5", priority: syslog.LOG_LOCAL5},
		{facility: "local6", priority: syslog.LOG_LOCAL6},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.priority, GetPriority(tc.facility))
	}
}

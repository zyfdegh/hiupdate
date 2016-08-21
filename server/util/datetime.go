package util

import (
	"strings"
	"time"
)

// GetDate returns today's date in "20160822"
func GetDate(t time.Time) string {
	// 2006-01-02T15:04:05Z07:00
	date := t.Format(time.RFC3339)[:10]
	tmp := strings.Split(date, "-")
	return strings.Join(tmp, "")
}

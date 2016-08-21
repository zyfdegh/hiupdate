package util

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

// modify date before testing
func TestGetDate(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z07:00")
	if err != nil {
		t.Fail()
	}
	var cases = []struct {
		t      time.Time
		expect string
	}{
		{t1, "20060102"},
	}

	for _, c := range cases {
		//call
		got := GetDate(c.t)
		//assert
		assert.Equal(t, c.expect, got)
	}
}

// modify date before testing
func TestGetToday(t *testing.T) {
	var cases = []struct {
		expect string
	}{
		{"20160822"},
	}

	for _, c := range cases {
		//call
		got := GetToday()
		//assert
		assert.Equal(t, c.expect, got)
	}
}

// modify date before testing
func TestGetYesterday(t *testing.T) {
	var cases = []struct {
		expect string
	}{
		{"20160821"},
	}

	for _, c := range cases {
		//call
		got := GetYesterday()
		//assert
		assert.Equal(t, c.expect, got)
	}
}

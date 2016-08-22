package util

import (
	"log"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

// modify date before testing
func TestGetDate(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2008-09-08T22:47:31+08:00")
	t2, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05-07:00")
	var cases = []struct {
		t      time.Time
		expect string
	}{
		{t1, "20080908"},
		{t2, "20060102"},
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
		{"20160823"},
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
		{"20160822"},
	}

	for _, c := range cases {
		//call
		got := GetYesterday()
		//assert
		assert.Equal(t, c.expect, got)
	}
}

func TestGetWeekday(t *testing.T) {
	var cases = []struct {
		date   string
		expect string
	}{
		{"20110823", "二"},
		{"20160821", "日"},
		{"20181212", "三"},
	}

	for _, c := range cases {
		//call
		got, err := GetWeekday(c.date)
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		//assert
		assert.Equal(t, c.expect, got)
	}
}

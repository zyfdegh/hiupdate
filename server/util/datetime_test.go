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
		{"20160825"},
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
		{"20160824"},
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

func TestSplitDate(t *testing.T) {
	var cases = []struct {
		date        string
		expect_yyyy string
		expect_mm   string
		expect_dd   string
	}{
		{"20110823", "2011", "08", "23"},
		{"20160821", "2016", "08", "21"},
		{"20181212", "2018", "12", "12"},
	}

	for _, c := range cases {
		//call
		got_yyyy, got_mm, got_dd := SplitDate(c.date)
		//assert
		assert.Equal(t, c.expect_yyyy, got_yyyy)
		assert.Equal(t, c.expect_mm, got_mm)
		assert.Equal(t, c.expect_dd, got_dd)
	}
}

func TestFormatDate(t *testing.T) {
	var cases = []struct {
		date      string
		seperator string
		expect    string
	}{
		{"20110823", "-", "2011-08-23"},
		{"20160821", ",", "2016,08,21"},
		{"20181212", ".", "2018.12.12"},
	}

	for _, c := range cases {
		//call
		got := FormatDate(c.date, c.seperator)
		//assert
		assert.Equal(t, c.expect, got)
	}
}

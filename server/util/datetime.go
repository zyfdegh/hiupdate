package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// GetDate parse date from t and return in "20160822"
func GetDate(t time.Time) string {
	// 2006-01-02T15:04:05Z07:00
	date := t.Format(time.RFC3339)[:10]
	tmp := strings.Split(date, "-")
	return strings.Join(tmp, "")
}

// GetDate returns today's date in "20160822"
func GetToday() string {
	return GetDate(time.Now())
}

// GetDate returns yesterday's date in "20160821"
func GetYesterday() string {
	return GetDate(time.Now().Add(-time.Hour * 24))
}

// GetWeek returns '一' if time is Monday
func GetWeekday(date string) (string, error) {
	// Convert date to RFC3339
	if len(date) != 8 {
		return "", errors.New("invalid date format")
	}
	y, m, d := SplitDate(date)
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	rfc := fmt.Sprintf("%s-%s-%sT08:00:00+08:00", y, m, d)
	t, err := time.Parse(time.RFC3339, rfc)
	if err != nil {
		return "", err
	}
	var weekdays = map[int]string{
		0: "日",
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
		6: "六",
	}
	return weekdays[int(t.Weekday())], nil
}

func SplitDate(date string) (yyyy, mm, dd string) {
	if len(date) != 8 {
		return
	}
	yyyy = string([]byte(date)[0:4])
	mm = string([]byte(date)[4:6])
	dd = string([]byte(date)[6:8])
	return
}

// FormatDate format date with seperator
// 20160902 -> 2016-09-02
func FormatDate(date string, seperator string) string {
	if len(date) != 8 {
		return ""
	}

	y, m, d := SplitDate(date)

	return strings.Join([]string{y, m, d}, seperator)
}

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
		fmt.Println("1")
		return "", errors.New("invalid date format")
	}
	y := []byte(date)[0:4]
	m := []byte(date)[4:6]
	d := []byte(date)[6:8]
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	rfc := fmt.Sprintf("%s-%s-%sT08:00:00+08:00", y, m, d)
	t, err := time.Parse(time.RFC3339, rfc)
	if err != nil {
		fmt.Println("2")
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
	fmt.Println("3")
	return weekdays[int(t.Weekday())], nil
}

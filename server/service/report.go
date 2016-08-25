package service

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/util"
)

// GetReport get report of a day
func GetReport(date string) (report *entity.Report, err error) {
	report = &entity.Report{}
	report.Title = "南京站会纪要"
	report.Date = date
	report.Weekday, err = util.GetWeekday(date)
	if err != nil {
		log.Printf("convert date to weekday error: %v", err)
		return
	}
	records, err := GetRecords(date)
	if err != nil {
		log.Printf("service get records error: %v", err)
		return
	}
	report.Records = records
	return report, nil
}

// GetReportText return report as string
func GetReportText(date string) (report string, err error) {
	r, _ := GetReport(date)
	return ConvertText(*r)
}

// ConvertText convert report as text file
func ConvertText(report entity.Report) (text string, err error) {
	var buf bytes.Buffer
	week, err := util.GetWeekday(report.Date)
	if err != nil {
		log.Printf("get weekday error: %v", err)
		return "", err
	}
	buf.WriteString(fmt.Sprintf("# %s 星期%s %s", report.Date, week, report.Title))
	number := 1
	for _, record := range report.Records {
		buf.WriteString("\n")
		buf.WriteString(fmt.Sprintf("%d. %s\n", number, record.Person.Name))
		buf.WriteString(fmt.Sprintf("昨天: %s\n", record.Content.Done))
		buf.WriteString(fmt.Sprintf("今天: %s\n", record.Content.Todo))
		if len(strings.TrimSpace(record.Content.Issue)) > 0 {
			buf.WriteString(fmt.Sprintf("问题: %s\n", record.Content.Issue))
		}
		number++
	}
	buf.WriteString("\n")
	buf.WriteString("* 以上顺序按照首次汇报日期排序\n")
	return buf.String(), err
}

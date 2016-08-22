package service

import (
	"fmt"
	"log"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/util"
)

func GetReport(date string) (report *entity.Report, err error) {
	report = &entity.Report{}
	report.Title = fmt.Sprintf("# %s 南京站会纪要\n", date)
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

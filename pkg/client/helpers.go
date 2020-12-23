package client

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	startDate = "start_date"
	endDate   = "end_date"
)

func StartEndDate(vars map[string]string) (string, string, error) {
	startDate := vars[startDate]
	endDate := vars[endDate]
	if len(startDate) == 0 && len(endDate) == 0 {
		return "", "", fmt.Errorf(" something went wrong: invalid start or end date. " +
			"Please enter date in YYYY-MM-DD format")
	}
	return startDate, endDate, nil
}

func DateRange(startDate, endDate string) ([]string, error) {
	var dates []string
	startDateSplit := strings.Split(startDate, "-")
	startYear, _ := strconv.Atoi(startDateSplit[0])
	startMonth, _ := strconv.Atoi(startDateSplit[1])
	startDay, _ := strconv.Atoi(startDateSplit[2])

	endDateSplit := strings.Split(endDate, "-")
	endYear, _ := strconv.Atoi(endDateSplit[0])
	endMonth, _ := strconv.Atoi(endDateSplit[1])
	endDay, _ := strconv.Atoi(endDateSplit[2])

	start := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	end := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	for rd := dateBetween(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		dates = append(dates, date.Format("2006-01-02"))
	}

	return dates, nil
}

func dateBetween(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

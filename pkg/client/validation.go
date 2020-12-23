package client

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const invalidFormatMessage = " start or end date has invalid format. Please adapt inputs to format: YYYY-MM-DD"

func ValidParameters(args ...string) error {
	// Splitter chars validation
	if !strings.Contains(args[0], "-") || !strings.Contains(args[1], "-") {
		return fmt.Errorf(invalidFormatMessage)
	}
	if strings.Count(args[0], "-") != 2 || strings.Count(args[1], "-") != 2 {
		return fmt.Errorf(invalidFormatMessage)
	}

	startDateSplit := strings.Split(args[0], "-")
	// Check length of date
	if len(startDateSplit) != 3 {
		return fmt.Errorf(invalidFormatMessage)
	}
	startYear, _ := strconv.Atoi(startDateSplit[0])
	startMonth, _ := strconv.Atoi(startDateSplit[1])
	startDay, _ := strconv.Atoi(startDateSplit[2])

	// Check length of year, month and day
	if len(startDateSplit[0]) != 4 || len(startDateSplit[1]) != 2 || len(startDateSplit[2]) != 2 {
		return fmt.Errorf(invalidFormatMessage)
	}

	endDateSplit := strings.Split(args[1], "-")
	// Check length of date
	if len(endDateSplit) != 3 {
		return fmt.Errorf(invalidFormatMessage)
	}
	endYear, _ := strconv.Atoi(endDateSplit[0])
	endMonth, _ := strconv.Atoi(endDateSplit[1])
	endDay, _ := strconv.Atoi(endDateSplit[2])
	// Check length of year, month and day
	if len(endDateSplit[0]) != 4 || len(endDateSplit[1]) != 2 || len(endDateSplit[2]) != 2 {
		return fmt.Errorf(invalidFormatMessage)
	}

	start := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	end := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)
	if end.Unix() < start.Unix() {
		return fmt.Errorf(" Invalid end date: earlier than start date")
	}

	return nil
}

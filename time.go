package main

import (
	"fmt"
	"time"
)

var layout string = "2006-01-02T15:04:05.000Z" //ISO8601 timestamp format

func main() {

	str := "2022-06-21T11:45:26.371Z"
	start := "2022-06-08T11:45:26.371Z"
	end := "2022-06-20T11:45:26.371Z"
	timeToCheck, err1 := parseTime(str)
	if err1 != nil {
		fmt.Println("Error while parsing date")
	}
	startTimeParsed, err2 := parseTime(start)
	if err2 != nil {
		fmt.Println("Error while parsing date")
	}

	endTimeParsed, err3 := parseTime(end)
	if err3 != nil {
		fmt.Println("Error while parsing date")
	}
	fmt.Println(inTimeSpan(startTimeParsed, endTimeParsed, timeToCheck))

}

//function to parse time with the specified layout
func parseTime(inputTime string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, inputTime)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil

}

//function to check given time is within the start and end dates specified
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

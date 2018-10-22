package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// FILE TIME IP DOM
func main() {

	today := strings.Split(time.Now().String(), " ")
	todaysSeconds := parseDate(today[0], today[1])

	parArgs := os.Args[1:]
	argsLen := len(parArgs)
	minutes := -1
	var err error
	ip := ""
	dns := ""
	fileName := ""

	// code ?
	if argsLen < 2 {
		fmt.Println("Need to specify file and time in minutes")
		os.Exit(1)
	} else {
		fileName = parArgs[0]
		minutes, err = strconv.Atoi(parArgs[1])
		if err != nil {
			fmt.Println("Need to  specify amount of minutes")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if argsLen > 2 {
			ip = parArgs[2]
		}

		if argsLen > 3 {
			dns = parArgs[3]
		}
	}

	// From here

	minutesToSeconds := float64(minutes * 60.0)
	timeLimit := todaysSeconds - minutesToSeconds
	getTMinLog(fileName, ip, dns, timeLimit)
}

func getTMinLog(fname, ip, dns string, time float64) {
	readChunks(fname)
}

func readChunks(fname string) {
	fmt.Println("asdfsdf")
	csvFile, _ := os.Open(fname)
	defer csvFile.Close()
	fi, err := csvFile.Stat()
	check(err)
	csvFileSize := fi.Size()
	csvFile.Seek(csvFileSize/2, 1)
	buf := make([]byte, 4096)
	_, err = io.ReadFull(csvFile, buf)
	check(err)
	fmt.Printf("Data read: %s\n", buf)

}

func binarySearch(buf []byte, timeLimit float64)  int {
	startIndex := 0
	endIndex := len(buf) - 1

	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2
		if recordToSeconds(string(buf[median])) < timeLimit {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}
	}

	if startIndex == len(buf) || recordToSeconds(string(buf[startIndex])) != timeLimit {
		return -1
	} else {
		return startIndex
	}
}

func recordToSeconds(record string) float64 {
	return parseDate(extractDate(record))
}
func extractDate(record string) (string, string) {
	columns := strings.Split(record, " ")
	return columns[0], columns[1]
}

func parseDate(ymd, hms string) float64 {
	year := strings.Split(ymd, "-")
	hours := strings.Split(hms, ":")

	yearToSeconds, err := strconv.ParseFloat(year[0], 64)
	check(err)
	yearToSeconds *= 525600.0 * 60.0

	monthsToSeconds, err := strconv.ParseFloat(year[1], 64)
	check(err)
	monthsToSeconds *= 43800.0 * 60.0

	daysToSeconds, err := strconv.ParseFloat(year[2], 64)
	check(err)

	daysToSeconds *= 1440.0 * 60.0

	hoursToSeconds, err := strconv.ParseFloat(hours[0], 64)
	check(err)

	hoursToSeconds *= 60.0 * 60.0

	minutesToSeconds, err:= strconv.ParseFloat(hours[1], 64)
	check(err)
	minutesToSeconds *= 60.0

	secondsToSeconds, err := strconv.ParseFloat(hours[2], 64)
	check(err)

	return yearToSeconds + monthsToSeconds + daysToSeconds + hoursToSeconds + minutesToSeconds + secondsToSeconds
}

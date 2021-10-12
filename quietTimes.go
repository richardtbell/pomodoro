package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type quietTimes []quietTime

type quietTime struct {
	Start string
	End   string
}

func isQuietTime() bool {
	jsonFile, err := os.Open("quietTimes.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quietTimes quietTimes
	json.Unmarshal(byteValue, &quietTimes)

	ct := time.Now()
	isQuietTime := false
	for _, quietTime := range quietTimes {
		quietTimeStart := getTimeFromString(quietTime.Start)
		quietTimeEnd := getTimeFromString(quietTime.End)
		if ct.After(quietTimeStart) && ct.Before(quietTimeEnd) {
			isQuietTime = true
		}
	}

	return isQuietTime
}

func getTimeFromString(s string) time.Time {
	ct := time.Now()
	t := strings.Split(s, ":")
	hour, _ := strconv.Atoi(t[0])
	min, _ := strconv.Atoi(t[1])
	return time.Date(ct.Year(), ct.Month(), ct.Day(), hour, min, 0, 0, ct.Location())
}

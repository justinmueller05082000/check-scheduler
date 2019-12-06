package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

func checkScheduler() {
	jsonName := flag.String("config", "", "Choose .json file to read the configuration from")
	flag.Parse()

	nameContent, nameError := ioutil.ReadFile(*jsonName)
	if nameError != nil {
		fmt.Println(nameError.Error())
		return
	}

	prints := make(map[string]float64)

	err2 := json.Unmarshal(nameContent, &prints)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	timeAbsolute := make(map[string]time.Time)

	for service, value := range prints {
		timeAbsolute[service] = time.Now().Add(time.Duration(float64(time.Second) * value))
	}

	for {
		lowestValue := time.Unix(1<<63-1, 1<<63-1)

		for _, value := range timeAbsolute {
			if value.Before(lowestValue) {
				lowestValue = value
			}
		}

		time.Sleep(lowestValue.Sub(time.Now()))
		for text, value := range timeAbsolute {
			if !time.Now().Before(value) {
				fmt.Printf("%s %s \n", time.Now().Format("Monday 15:04:05 01-02-2006"), text)
				timeAbsolute[text] = value.Add(time.Duration(prints[text] * float64(time.Second)))
			}
		}
	}
}

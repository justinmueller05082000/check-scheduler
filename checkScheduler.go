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

	for text, pause := range prints {
		go printTest(pause, text)
	}

	for {
		time.Sleep(time.Second)
	}
}

func printTest(pause float64, text string) {
	for {
		time.Sleep(time.Duration(pause * float64(time.Second)))
		fmt.Printf("%s %s \n", time.Now().Format("Monday 15:04:05 01-02-2006"), text)
	}
}

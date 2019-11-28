package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type toPrint struct {
	Text string `json:"text"`
}

func checkScheduler() {
	content, err := ioutil.ReadFile("toPrint.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var prints []toPrint

	err2 := json.Unmarshal(content, &prints)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	var text string

	for _, x := range prints {
		text = x.Text
	}

	for true {
		time.Sleep(30 * time.Second)
		for x := 1; x <= 100; x++ {
			fmt.Printf("%s %s \n", time.Now().Format("Monday 15:04:05 01-02-2006"), text)
		}
	}
}

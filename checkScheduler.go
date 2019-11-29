package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

type toPrint struct {
	Text string `json:"text"`
}

func checkScheduler() {
	jsonname := flag.String("config", "", "Choose .json file to read the configuration from")
	flag.Parse()

	content, err := ioutil.ReadFile(*jsonname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var prints toPrint

	err2 := json.Unmarshal(content, &prints)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	for true {
		time.Sleep(30 * time.Second)
		for x := 1; x <= 100; x++ {
			fmt.Printf("%s %s \n", time.Now().Format("Monday 15:04:05 01-02-2006"), prints.Text)
		}
	}
}

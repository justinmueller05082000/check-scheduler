package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

type toPrint struct {
	Text  string  `json:"text"`  // Defines the check's name
	Times int     `json:"times"` // Defines the repeat of checks
	Pause float64 `json:"pause"` // Defines the break between the checks
}

func checkScheduler() {
	jsonName := flag.String("config", "", "Choose .json file to read the configuration from")
	flag.Parse()

	nameContent, nameError := ioutil.ReadFile(*jsonName)
	if nameError != nil {
		fmt.Println(nameError.Error())
		return
	}

	var prints toPrint

	err2 := json.Unmarshal(nameContent, &prints)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	for true {
		time.Sleep(time.Duration(prints.Pause) * time.Second)
		for x := 1; x <= prints.Times+1; x++ {
			fmt.Printf("%s %s \n", time.Now().Format("Monday 15:04:05 01-02-2006"), prints.Text)
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func checkScheduler() {
	for true {
		time.Sleep(30 * time.Second)
		for x := 1; x <= 100; x++ {
			fmt.Println(time.Now().Format("Monday 15:04:05 01-02-2006"))
			fmt.Println("check_disk")
			fmt.Println(x) // Überprüfung

		}
	}
}

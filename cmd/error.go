package main

import "log"

func recoverFunc() {
	if v := recover(); v != nil {
		log.Fatal(`panic: `, v)
	}
}

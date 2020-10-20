package main

import "log"

func recoverFunc() {
	if v := recover(); v != nil {
		log.Fatal(`panic: `, v)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"time"
)

func notConcurrency() {
	t := time.Now()
	defer func() {
		log.Printf("time of notConcurrency:%v", time.Since(t).Microseconds())
	}()
	var x []int
	for _, n := range generateSlice(300) {
		x = append(x, n)
	}

	var sqX []int
	for _, n := range x {
		sqX = append(sqX, n*n)
	}
	for _, n := range sqX {
		_ = n
	}
}

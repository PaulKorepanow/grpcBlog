package main

import (
	"log"
	"time"
)

func pipeline() {
	t := time.Now()
	done := make(chan struct{})
	defer func() {
		//log.Println("close done channel")
		log.Printf("time of pipeline:%v", time.Since(t).Microseconds())
		close(done)
	}()
	c := gen(generateSlice(300)...)
	out1 := sq(done, c)
	out2 := sq(done, c)

	out := merge(done, out1, out2)
	for n := range out {
		_ = n
	}
	//fmt.Println(<-out, <-out, <-out)
}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

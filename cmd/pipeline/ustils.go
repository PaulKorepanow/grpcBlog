package main

import (
	"log"
	"math/rand"
	"sync"
)

func merge(done chan struct{}, ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := new(sync.WaitGroup)
	output := func(c <-chan int, idx int) {
		for n := range c {
			select {
			case out <- n:
				//log.Println(idx, n)
			case <-done:
				log.Println("done case")
				return
			}
		}
		//log.Println(idx, "Done")
		wg.Done()
	}

	wg.Add(len(ch))
	for idx, c := range ch {
		go output(c, idx)
	}

	go func() {
		wg.Wait()
		//log.Println("close merge channel")
		close(out)
	}()
	return out
}

func generateSlice(nums int) []int {
	res := make([]int, nums)
	for idx := range res {
		res[idx] = rand.Int()
	}
	return res
}

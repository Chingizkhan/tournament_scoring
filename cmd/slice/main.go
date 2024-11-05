package main

import (
	"log"
	"sync"
)

func main() {
	//slice()
	second()
}

func second() {

}

func mergeChannels(channels ...<-chan string) <-chan string {
	res := make(chan string)

	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				res <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func slice() {
	sl := make([]int, 2, 10)
	log.Printf("sl before: %+v", sl)
	change(sl)
	//changeWithPtr(sl)
	log.Printf("sl after: %+v", sl[0:3])
}

func change(sl []int) {
	sl = append(sl, 3)

	//*sl = append(*sl, 3)

	log.Printf("sl inside change: %+v", sl)
}

func changeWithPtr(sl *[]int) {
	*sl = append(*sl, 3)

	log.Printf("sl inside change: %+v", sl)
}

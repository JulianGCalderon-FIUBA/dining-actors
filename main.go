package main

import "time"

const PHILOSOPHERS = 3

func main() {
	newChain(PHILOSOPHERS)

	time.Sleep(10 * time.Second)
}

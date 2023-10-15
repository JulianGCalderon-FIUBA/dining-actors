package main

import "time"

const PHILOSOPHERS = 5

func main() {
	chain := newChain(PHILOSOPHERS)
	chain.start()

	<-time.After(5 * time.Second)
	chain.shutdown()

}

package main

import (
	"dining-actors/philosopher"
	"time"
)

const PHILOSOPHERS = 5

func main() {
	chain := philosopher.NewChain(PHILOSOPHERS)
	chain.Start()

	<-time.After(5 * time.Second)
	chain.Shutdown()
}

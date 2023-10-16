package main

import (
	"dining-actors/internal"
	"time"
)

const PHILOSOPHERS = 5

func main() {
	chain := internal.MakeChain(PHILOSOPHERS)
	chain.Start()

	<-time.After(5 * time.Second)
	chain.Shutdown()
}

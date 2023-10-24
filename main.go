package main

import (
	"dining-actors/internal"
	"os"
	"os/signal"
	"syscall"
)

const PHILOSOPHERS = 5

func main() {
	chain := internal.MakeChain(PHILOSOPHERS)
	chain.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	chain.Shutdown()
}

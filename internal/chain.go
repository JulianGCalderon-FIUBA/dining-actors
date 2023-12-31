package internal

type Chain []Philosopher

func MakeChain(size int) (philosophers Chain) {
	philosophers = make([]Philosopher, size)

	for id := range philosophers {
		p := &philosophers[id]
		p.Init(id)
	}

	for id := range philosophers {
		rightId := (id + 1) % size

		p := &philosophers[id]
		rightP := &philosophers[rightId]

		Link(p, rightP)
	}

	return
}

func (c Chain) Start() {
	for _, p := range c {
		go p.Loop()
	}
}

func (c Chain) Shutdown() {
	ch := make(chan struct{})
	for _, p := range c {
		p.channel <- Shutdown{ch}
	}

	for i := 0; i < len(c); i++ {
		<-ch
	}
}

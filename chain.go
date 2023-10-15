package main

type Chain []Philosopher

func newChain(size int) (philosophers Chain) {
	philosophers = make([]Philosopher, size)

	for i := range philosophers {
		philosophers[i].id = i
		philosophers[i].channel = make(chan Message, 100)

		leftI := (i + size - 1) % size
		rightI := (i + 1) % size

		philosophers[i].rightPhilosopher = &philosophers[rightI]
		philosophers[i].leftPhilosopher = &philosophers[leftI]

		if i < rightI {
			philosophers[i].rightStick = Dirty
		} else {
			philosophers[rightI].leftStick = Dirty
		}
	}

	return
}

func (c Chain) start() {
	for _, p := range c {
		p.Prepare()
		go p.Loop()
	}
}

func (c Chain) shutdown() {
	for _, p := range c {
		p.Send(Shutdown)
	}
}

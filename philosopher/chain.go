package philosopher

type Chain []Philosopher

func NewChain(size int) (philosophers Chain) {
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

func (c Chain) Start() {
	for _, p := range c {
		go p.Loop()
	}
}

func (c Chain) Shutdown() {
	for _, p := range c {
		p.Send(Shutdown)
	}
}

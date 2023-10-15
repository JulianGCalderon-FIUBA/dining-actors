package main

import "fmt"

type Philosopher struct {
	id                  int
	leftStick           Stick
	leftStickRequested  bool
	rightStick          Stick
	rightStickRequested bool
	leftPhilosopher     *Philosopher
	rightPhilosopher    *Philosopher
	channel             chan Message
}

func newChain(size int) (philosophers []Philosopher) {
	philosophers = make([]Philosopher, size)

	for i := range philosophers {
		leftI := (i + size - 1) % size
		rightI := (i + 1) % size

		philosophers[i].id = i
		philosophers[i].leftPhilosopher = &philosophers[leftI]
		philosophers[i].rightPhilosopher = &philosophers[rightI]
		philosophers[i].channel = make(chan Message, 100)

		if i < rightI {
			philosophers[i].rightStick = Dirty
		} else {
			philosophers[rightI].leftStick = Dirty
		}
	}

	for _, p := range philosophers {
		p.send(Start)
		go p.receive()
	}

	return
}

func (p Philosopher) receive() {
	for msg := range p.channel {
		switch msg {
		case RightStickRequest:
			p.handleRightStickRequest()
		case LeftStickRequest:
			p.handleLeftStickRequest()
		case RightStickSend:
			p.handleRightStickSend()
		case LeftStickSend:
			p.handleLeftStickSend()
		case Start:
			p.handleStart()
		}
	}
}

func (p *Philosopher) handleRightStickRequest() {
	p.log("Was requested right stick")

	if p.rightStick == Dirty {
		p.log("Sending right stick")
		p.rightStick = Waiting
		p.rightPhilosopher.send(LeftStickSend)
		p.rightPhilosopher.send(LeftStickRequest)
	} else {
		p.rightStickRequested = true
	}
}

func (p *Philosopher) handleLeftStickRequest() {
	p.log("Was requested left stick")

	if p.leftStick == Dirty {
		p.log("Sending left stick")
		p.leftStick = Waiting
		p.leftPhilosopher.send(RightStickSend)
		p.leftPhilosopher.send(RightStickRequest)
	} else {
		p.leftStickRequested = true
	}
}

func (p *Philosopher) handleRightStickSend() {
	p.log("Was sent right stick")

	p.rightStick = Clean

	if p.leftStick == Clean {
		p.dine()
	}

	if p.leftStick == None {
		p.log("Requesting left stick")
		p.leftStick = Waiting
		p.leftPhilosopher.send(RightStickRequest)
	}
}

func (p *Philosopher) handleLeftStickSend() {
	p.log("Was sent left stick")

	p.leftStick = Clean

	if p.rightStick == Clean {
		p.dine()
	}

	if p.rightStick == None {
		p.log("Requesting right stick")
		p.rightStick = Waiting
		p.rightPhilosopher.send(LeftStickRequest)
	}
}

func (p *Philosopher) handleStart() {
	if p.leftStick == None {
		p.log("Requesting left stick")
		p.leftPhilosopher.send(RightStickRequest)
		p.leftStick = Waiting
	}

	if p.rightStick == None {
		p.log("Requesting right stick")
		p.rightPhilosopher.send(LeftStickRequest)
		p.rightStick = Waiting
	}
}

func (p *Philosopher) dine() {
	fmt.Println("Dining!")

	p.leftStick = Dirty
	p.rightStick = Dirty

	if p.leftStickRequested {
		p.leftStick = None
		p.leftPhilosopher.send(RightStickSend)
	}

	if p.rightStickRequested {
		p.rightStick = None
		p.rightPhilosopher.send(LeftStickSend)
	}

	p.leftStickRequested = false
	p.rightStickRequested = false

	p.send(Start)
}

func (p *Philosopher) log(msg string) {
	fmt.Printf("Philosopher %v: %v\n", p.id, msg)
}

func (p Philosopher) send(msg Message) {
	p.channel <- msg
}

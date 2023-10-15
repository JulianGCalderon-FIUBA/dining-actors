package main

import (
	"fmt"
	// "time"
)

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

type Chain []Philosopher

func newChain(size int) (philosophers Chain) {
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

func (c Chain) close() {
	for _, p := range c {
		p.send(Close)
	}
}

func (p *Philosopher) requestRight() {
	p.log("Requesting right stick")
	p.rightPhilosopher.send(LeftStickRequest)
	p.rightStick = Waiting
}

func (p *Philosopher) requestLeft() {
	p.log("Requesting left stick")
	p.leftPhilosopher.send(RightStickRequest)
	p.leftStick = Waiting
}

func (p *Philosopher) sendRight() {
	p.log("Sending right stick")
	p.rightPhilosopher.send(LeftStickSend)
	p.rightStick = None
}

func (p *Philosopher) sendLeft() {
	p.log("Sending left stick")
	p.leftPhilosopher.send(RightStickSend)
	p.leftStick = None
}

func (p Philosopher) send(msg Message) {
	p.channel <- msg
}

func (p *Philosopher) log(msg string) {
	fmt.Printf("Philosopher %v: %v\n", p.id, msg)
}

func (p Philosopher) receive() {
	for msg := range p.channel {
		switch msg {
		case Start:
			p.start()
		case RightStickRequest:
			p.rightStickRequest()
		case LeftStickRequest:
			p.leftStickRequest()
		case RightStickSend:
			p.rightStickSend()
		case LeftStickSend:
			p.leftStickSend()
		}
	}
}

func (p *Philosopher) start() {
	if p.leftStick == None {
		p.requestLeft()
	}

	if p.rightStick == None {
		p.requestRight()
	}
}

func (p *Philosopher) rightStickSend() {
	p.rightStick = Clean

	if p.leftStick == Clean {
		p.dine()
	} else if p.leftStick == None {
		p.requestLeft()
	}
}

func (p *Philosopher) leftStickSend() {

	p.leftStick = Clean

	if p.rightStick == Clean {
		p.dine()
	} else if p.rightStick == None {
		p.requestRight()
	}
}

func (p *Philosopher) dine() {
	fmt.Println("Dining!")
	// time.Sleep(500 * time.Millisecond)

	p.clean_up()
}

func (p *Philosopher) clean_up() {
	p.leftStick = Dirty
	p.rightStick = Dirty

	if p.leftStickRequested {
		p.sendLeft()
		p.requestLeft()
		p.leftStickRequested = false
	}

	if p.rightStickRequested {
		p.sendRight()
		p.requestRight()
		p.rightStickRequested = false
	}
}

func (p *Philosopher) rightStickRequest() {
	if p.rightStick == Dirty {
		p.sendRight()
		p.requestRight()
	} else {
		p.rightStickRequested = true
	}
}

func (p *Philosopher) leftStickRequest() {
	if p.leftStick == Dirty {
		p.sendLeft()
		p.requestLeft()
	} else {
		p.leftStickRequested = true
	}
}

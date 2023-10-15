package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id               int
	leftStick        Stick
	rightStick       Stick
	leftPhilosopher  *Philosopher
	rightPhilosopher *Philosopher
	channel          chan Message
}

func (p *Philosopher) Send(msg Message) {
	p.channel <- msg
}

func (p *Philosopher) Prepare() {
	if p.leftStick == None {
		p.requestLeft()
	}

	if p.rightStick == None {
		p.requestRight()
	}
}

func (p Philosopher) Loop() {
	for msg := range p.channel {
		switch msg {
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

/* Handlers */

func (p *Philosopher) rightStickSend() {
	p.log("Got right stick")

	p.rightStick = Clean

	if p.leftStick == Clean {
		p.dine()
	}
}

func (p *Philosopher) leftStickSend() {
	p.log("Got left stick")

	p.leftStick = Clean

	if p.rightStick == Clean {
		p.dine()
	}
}

func (p *Philosopher) rightStickRequest() {
	if p.rightStick == Dirty {
		p.sendRight()
		p.requestRight()
	}
}

func (p *Philosopher) leftStickRequest() {
	if p.leftStick == Dirty {
		p.sendLeft()
		p.requestLeft()
	}
}

/* Logic */

func (p *Philosopher) dine() {
	p.scream("Dining!")
	time.Sleep(500 * time.Millisecond)
	p.clean_up()
	time.Sleep(500 * time.Millisecond)
}

func (p *Philosopher) clean_up() {
	p.sendLeft()
	p.sendRight()

	p.requestLeft()
	p.requestRight()
}

/* Helper Functions */

func (p *Philosopher) requestRight() {
	p.log("Requesting right stick")
	p.rightPhilosopher.Send(LeftStickRequest)
}

func (p *Philosopher) requestLeft() {
	p.log("Requesting left stick")
	p.leftPhilosopher.Send(RightStickRequest)
}

func (p *Philosopher) sendRight() {
	p.log("Sending right stick")
	p.rightPhilosopher.Send(LeftStickSend)
	p.rightStick = None
}

func (p *Philosopher) sendLeft() {
	p.log("Sending left stick")
	p.leftPhilosopher.Send(RightStickSend)
	p.leftStick = None
}

func (p *Philosopher) log(msg string) {
	fmt.Printf("Philosopher %v: %v\n", p.id, msg)
}

func (p *Philosopher) scream(msg string) {
	fmt.Printf("\033[31m"+"Philosopher %v, %v\n"+"\033[0m", p.id, msg)
}

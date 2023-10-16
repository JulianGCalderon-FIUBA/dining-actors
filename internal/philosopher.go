package internal

import (
	"fmt"
	"time"
)

const CHANNEL_SIZE = 8

type Philosopher struct {
	id               int
	leftStick        Stick
	rightStick       Stick
	leftPhilosopher  chan Message
	rightPhilosopher chan Message
	channel          chan Message
}

func (p *Philosopher) Init(id int) {
	p.id = id
	p.channel = make(chan Message, CHANNEL_SIZE)
}

func Link(left *Philosopher, right *Philosopher) {
	left.rightPhilosopher = right.channel
	right.leftPhilosopher = left.channel

	if left.id < right.id {
		left.rightStick = Dirty
	} else {
		right.leftStick = Dirty
	}
}

func (p Philosopher) Loop() {
	p.prepare()

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

func (p *Philosopher) prepare() {
	if p.leftStick == None {
		p.requestLeft()
	}

	if p.rightStick == None {
		p.requestRight()
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

func (p *Philosopher) rightStickSend() {
	p.say("Got right stick")

	p.rightStick = Clean

	if p.leftStick == Clean {
		p.dine()
	}
}

func (p *Philosopher) leftStickSend() {
	p.say("Got left stick")

	p.leftStick = Clean

	if p.rightStick == Clean {
		p.dine()
	}
}

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
	p.rightPhilosopher <- LeftStickRequest
}

func (p *Philosopher) requestLeft() {
	p.log("Requesting left stick")
	p.leftPhilosopher <- RightStickRequest
}

func (p *Philosopher) sendRight() {
	p.log("Sending right stick")
	p.rightPhilosopher <- LeftStickSend
	p.rightStick = None
}

func (p *Philosopher) sendLeft() {
	p.log("Sending left stick")
	p.leftPhilosopher <- RightStickSend
	p.leftStick = None
}

func (p *Philosopher) log(msg string) {
	fmt.Printf("Philosopher %v: %v\n", p.id, msg)
}

func (p *Philosopher) say(msg string) {
	fmt.Printf("\033[34m"+"Philosopher %v, %v\n"+"\033[0m", p.id, msg)
}

func (p *Philosopher) scream(msg string) {
	fmt.Printf("\033[31m"+"Philosopher %v, %v\n"+"\033[0m", p.id, msg)
}

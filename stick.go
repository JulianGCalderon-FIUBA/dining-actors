package main

type Stick uint8

const (
	None Stick = iota
	Waiting
	Dirty
	Clean
)

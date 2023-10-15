package main

type Message uint8

const (
	Shutdown Message = iota
	RightStickRequest
	RightStickSend
	LeftStickRequest
	LeftStickSend
)

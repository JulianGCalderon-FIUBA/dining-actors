package main

type Message uint8

const (
	Start Message = iota
	Close
	RightStickRequest
	RightStickSend
	LeftStickRequest
	LeftStickSend
)

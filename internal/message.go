package internal

type Message uint8

const (
	Shutdown Message = iota
	RightStickRequest
	RightStickSend
	LeftStickRequest
	LeftStickSend
)

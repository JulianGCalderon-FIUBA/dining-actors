package philosopher

type Message uint8

const (
	Shutdown Message = iota
	RightStickRequest
	RightStickSend
	LeftStickRequest
	LeftStickSend
)

package internal

type Message interface{}

type RightStickRequest struct{}
type RightStickSend struct{}
type LeftStickRequest struct{}
type LeftStickSend struct{}

type Shutdown struct {
	response chan bool
}

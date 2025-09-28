package channels

import (
	"errors"
)

type MessageChannel struct {
	ch chan []byte
}

func NewMessageChannel(bufferSize int) *MessageChannel{
	return &MessageChannel{ch: make(chan []byte, bufferSize)}
}

func (ch *MessageChannel) Send(msg []byte) error {
	select {
	case ch.ch <- msg:
		return nil
	default:
		return errors.New("Channel is full")
	}
}

func (ch *MessageChannel) Receive() ([]byte, error) {
	select {
	case msg := <-ch.ch:
		return msg, nil
	default:
		return nil, errors.New("Channel is empty")
	}
}
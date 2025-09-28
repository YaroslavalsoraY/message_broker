package interfaces

type MessageBroker interface {
	Send(msg []byte) error
	Receive() ([]byte, error)
}

package interfaces

type MessageBroker interface {
	Send(msg []byte) error
	Get() error
}

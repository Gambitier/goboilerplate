package messageQueue

type MessageHandler func(message []byte)

// MessageQueue defines the interface for a message queue.
type MessageQueue interface {
	Publish(channel string, message []byte) error
	Subscribe(channel string, handler MessageHandler) error
	Close() error
}

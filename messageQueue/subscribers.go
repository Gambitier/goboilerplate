package messageQueue

type Subscriber interface {
	Register(queue MessageQueue)
}

type SubscriberRegistry struct {
	queue MessageQueue
}

func NewSubscriberRegistry(queue MessageQueue) *SubscriberRegistry {
	return &SubscriberRegistry{queue: queue}
}

func (r *SubscriberRegistry) RegisterSubscriber(sub Subscriber) {
	sub.Register(r.queue)
}

func (r *SubscriberRegistry) RegisterSubscribers(subs ...Subscriber) {
	for _, sub := range subs {
		sub.Register(r.queue)
	}
}

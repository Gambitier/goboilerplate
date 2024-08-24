package appContext

import (
	"github.com/gambitier/gocomm/messageQueue"
	queuesubscribers "github.com/gambitier/gocomm/modules/users/queueSubscribers"
)

func (app *AppContext) RegisterMsgQueues() {
	registry := messageQueue.NewSubscriberRegistry(app.MessageQueue)
	registry.RegisterSubscribers(
		// Register messageQueue subscribers here..
		queuesubscribers.NewUploadAvatarSubscriber(app.TempFileStorage),
	)
}

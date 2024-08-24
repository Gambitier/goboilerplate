//go:build wireinject
// +build wireinject

package appContext

import (
	"github.com/gambitier/gocomm/config"
	database "github.com/gambitier/gocomm/db/database"
	"github.com/gambitier/gocomm/messageQueue"
	"github.com/google/wire"
)

// bind interface to implementation
var setMessageQueue = wire.NewSet(
	messageQueue.NewRedisMessageQueue,
	wire.Bind(new(messageQueue.MessageQueue), new(*messageQueue.RedisMessageQueue)),
)

func InitAppContext() (*AppContext, error) {
	wire.Build(
		NewAppContext,
		config.NewConfig,
		database.NewDatabaseRepo,
		setMessageQueue,
	)
	return &AppContext{}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appContext

import (
	"github.com/gambitier/gocomm/config"
	"github.com/gambitier/gocomm/db/database"
	"github.com/gambitier/gocomm/messageQueue"
	"github.com/google/wire"
)

// Injectors from app.initializer.go:

func InitAppContext() (*AppContext, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	databaseRepo, err := databse.NewDatabaseRepo(conf)
	if err != nil {
		return nil, err
	}
	redisMessageQueue, err := messageQueue.NewRedisMessageQueue(conf)
	if err != nil {
		return nil, err
	}
	appContext := NewAppContext(conf, databaseRepo, redisMessageQueue)
	return appContext, nil
}

// app.initializer.go:

// bind interface to implementation
var setMessageQueue = wire.NewSet(messageQueue.NewRedisMessageQueue, wire.Bind(new(messageQueue.MessageQueue), new(*messageQueue.RedisMessageQueue)))

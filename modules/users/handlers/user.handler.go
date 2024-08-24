package handlers

import (
	database "github.com/gambitier/gocomm/db/database"
	"github.com/gambitier/gocomm/messageQueue"
	"github.com/gambitier/gocomm/storage/localfile"
)

type UserHandler struct {
	dbRepo       *database.DatabaseRepo
	FileStorage  localfile.LocalFileStorageImpl
	MessageQueue messageQueue.MessageQueue
}

func NewUserHandler(dbRepo *database.DatabaseRepo, fileSystem localfile.LocalFileStorageImpl, messageQueue messageQueue.MessageQueue) *UserHandler {
	return &UserHandler{
		dbRepo,
		fileSystem,
		messageQueue,
	}
}

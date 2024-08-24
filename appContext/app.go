package appContext

import (
	"github.com/gambitier/gocomm/config"
	database "github.com/gambitier/gocomm/db/database"
	"github.com/gambitier/gocomm/messageQueue"
	"github.com/gambitier/gocomm/storage/localfile"
	"github.com/spf13/afero"
)

type AppContext struct {
	Configs         *config.Conf
	DbRepo          *database.DatabaseRepo
	MessageQueue    messageQueue.MessageQueue
	TempFileStorage localfile.LocalFileStorageImpl // `afero.Fs` is an interface
}

func NewAppContext(configs *config.Conf, dbRepo *database.DatabaseRepo, mq messageQueue.MessageQueue) *AppContext {
	fs := afero.NewBasePathFs(afero.NewOsFs(), configs.TempFileStoragePath)
	fileStorage := localfile.NewLocalFileStorage(fs)

	app := &AppContext{
		Configs:         configs,
		DbRepo:          dbRepo,
		MessageQueue:    mq,
		TempFileStorage: fileStorage,
	}

	return app
}

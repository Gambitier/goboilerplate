package queuesubscribers

import (
	"log"

	"github.com/gambitier/gocomm/imageProcessor"
	"github.com/gambitier/gocomm/messageQueue"
	"github.com/gambitier/gocomm/storage/localfile"
)

const UploadAvatarChannel = "avatarUploads"

type UploadAvatarSubscriber struct {
	fileStorage localfile.LocalFileStorageImpl
}

// `NewUploadAvatarSubscriber` is returning struct that implements `messageQueue.Subscriber` interface with `Register` method
func NewUploadAvatarSubscriber(fileStorage localfile.LocalFileStorageImpl) *UploadAvatarSubscriber {
	return &UploadAvatarSubscriber{fileStorage: fileStorage}
}

func (s *UploadAvatarSubscriber) Register(queue messageQueue.MessageQueue) {
	queue.Subscribe(UploadAvatarChannel, s.handleAvatarUpload)
}

func (s *UploadAvatarSubscriber) handleAvatarUpload(message []byte) {
	filePath := string(message)
	proc := imageProcessor.NewImageProcessor("AvatarPreviews", s.fileStorage)
	_, err := proc.GeneratePreviewImageFromPath(filePath)
	if err != nil {
		log.Printf("Failed to generate preview of file: %v | err: %v", filePath, err)
	}
}

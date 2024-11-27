package webhookController

import (
	"github.com/gin-gonic/gin"
)

type FileService interface {
	RebuildIndex(firstBuild bool) error
}

type Controller struct {
	fileService     FileService
	WebhookEndpoint string
}

func New(fileService FileService, webhookEndpoint string) *Controller {
	return &Controller{fileService: fileService, WebhookEndpoint: webhookEndpoint}
}

func (w *Controller) PostWebhook(_ *gin.Context) {
	go func() {
		_ = w.fileService.RebuildIndex(false)
	}()
}

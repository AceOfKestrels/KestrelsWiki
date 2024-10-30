package webhookController

import (
	"api/logger"
	"github.com/gin-gonic/gin"
)

type SearchService interface {
	UpdateIndex() error
}

type WebhookController struct {
	searchService   SearchService
	WebhookEndpoint string
}

func NewWebhookController(searchService SearchService, webhookEndpoint string) *WebhookController {
	return &WebhookController{searchService: searchService, WebhookEndpoint: webhookEndpoint}
}

func (w *WebhookController) PostWebhook(_ *gin.Context) {
	go func() {
		err := w.searchService.UpdateIndex()
		if err != nil {
			logger.Println(logger.API, err.Error())
		}
	}()
}

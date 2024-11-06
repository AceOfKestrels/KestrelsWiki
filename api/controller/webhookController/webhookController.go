package webhookController

import (
	"github.com/gin-gonic/gin"
)

type SearchService interface {
	RebuildIndex() error
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
		_ = w.searchService.RebuildIndex()
	}()
}

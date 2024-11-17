package ports

import "backend/internal/models"

type ChatPort interface {
    SendMessage(message *models.Message) error
    ReceiveMessages() ([]models.Message, error)
}
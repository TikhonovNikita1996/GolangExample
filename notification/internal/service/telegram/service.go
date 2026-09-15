package telegram

import (
	"bytes"
	"context"
	"embed"
	"text/template"

	"go.uber.org/zap"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/client/http"
	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

const chatID = 451414570

//go:embed templates/assembled_notification.tmpl
//go:embed templates/paid_notification.tmpl
var templateFS embed.FS

type orderPaidTemplateData struct {
	UUID            string
	UserUUID        string
	PaymentMethod   string
	TransactionUuid string
}

type shipAssembledTemplateData struct {
	UUID         string
	UserUUID     string
	BuildTimeSec int64
}

var (
	orderPaidTemplate     = template.Must(template.ParseFS(templateFS, "templates/paid_notification.tmpl"))
	shipAssembledTemplate = template.Must(template.ParseFS(templateFS, "templates/assembled_notification.tmpl"))
)

type service struct {
	telegramClient http.TelegramClient
}

// NewService создает новый Telegram сервис
func NewService(telegramClient http.TelegramClient) *service {
	return &service{
		telegramClient: telegramClient,
	}
}

func (s *service) SendPaidNotification(ctx context.Context, model model.OrderPaid) error {
	message, err := s.buildPaidMessage(model)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}

func (s *service) SendShipAssembledNotification(ctx context.Context, model model.ShipAssembled) error {
	message, err := s.buildAssembledMessage(model)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}

func (s *service) buildPaidMessage(model model.OrderPaid) (string, error) {
	data := orderPaidTemplateData{
		UUID:            model.OrderUuid,
		UserUUID:        model.UserUuid,
		PaymentMethod:   model.PaymentMethod,
		TransactionUuid: model.TransactionUuid,
	}

	var buf bytes.Buffer
	err := orderPaidTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (s *service) buildAssembledMessage(model model.ShipAssembled) (string, error) {
	data := shipAssembledTemplateData{
		UUID:         model.OrderUuid,
		UserUUID:     model.UserUuid,
		BuildTimeSec: model.BuildTimeSec,
	}

	var buf bytes.Buffer
	err := shipAssembledTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

package app

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/config"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/closer"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
)

type App struct {
	diContainer *diContainer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := a.runOrderPaidConsumer(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
		}
	}()

	go func() {
		if err := a.runShipAssembledConsumer(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "Shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "Component crashed, shutting down", zap.Error(err))
		// Триггерим cancel, чтобы остановить второй компонент
		cancel()
		// Дождись завершения всех задач (если есть graceful shutdown внутри)
		<-ctx.Done()
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initTelegramBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initTelegramBot(ctx context.Context) error {
	// Получаем бота из DI контейнера
	telegramBot := a.diContainer.TelegramBot(ctx)

	// Регистрируем обработчик для активации бота
	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		logger.Info(ctx, "chat id", zap.Int64("chat_id", update.Message.Chat.ID))

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "🛸 Notification Bot активирован! Теперь вы будете получать уведомления о оплате частей и сборки корабля.",
		})
		if err != nil {
			logger.Error(ctx, "Failed to send activation message", zap.Error(err))
		}
	})

	// Запускаем бота в фоне
	go func() {
		logger.Info(ctx, "🤖 Telegram bot started...")
		telegramBot.Start(ctx)
	}()

	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) runOrderPaidConsumer(ctx context.Context) error {
	orderConsumerErr := a.diContainer.OrderPaidConsumerService().RunConsumer(ctx)
	if orderConsumerErr != nil {
		return orderConsumerErr
	}
	logger.Info(ctx, "🚀 OrderPaid Kafka consumer running")

	return nil
}

func (a *App) runShipAssembledConsumer(ctx context.Context) error {
	assemblyConsumerErr := a.diContainer.ShipAssembledConsumerService().RunConsumer(ctx)
	if assemblyConsumerErr != nil {
		return assemblyConsumerErr
	}
	logger.Info(ctx, "🚀 ShipAssembled Kafka consumer running")

	return nil
}

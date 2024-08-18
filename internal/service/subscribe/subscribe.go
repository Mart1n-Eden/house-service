package subscribe

import (
	"context"

	"house-service/internal/domain"
)

type SubscribeRepo interface {
	NewSubscription(ctx context.Context, email string, houseId int) error
	GetMessagesForSubscription(ctx context.Context) ([]domain.Message, error)
}

type Sender interface {
	SendEmail(ctx context.Context, recipient string, message string) error
}

type Service struct {
	repo   SubscribeRepo
	sender Sender
}

func New(repo SubscribeRepo, sender Sender) *Service {
	return &Service{
		repo:   repo,
		sender: sender,
	}
}

func (s *Service) NewSubscription(ctx context.Context, email string, houseId int) error {
	return s.repo.NewSubscription(ctx, email, houseId)
}

func (s *Service) GetHouseBySubscription(ctx context.Context) {
	messages, err := s.repo.GetMessagesForSubscription(ctx)
	if err != nil {
		// TODO: log
		return
	}

	for _, message := range messages {
		err = s.sender.SendEmail(ctx, message.Recipient, message.Message)
		if err != nil {
			// TODO: log
			return
		}
	}
}

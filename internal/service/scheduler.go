package service

import (
	"github.com/robfig/cron/v3"
)

// StoreHandler
type SchedulerService struct {
	storage map[string]string
}

// NewStoreHandler
func NewSchedulerService(storage map[string]string) *SchedulerService {
	return &SchedulerService{
		storage: storage,
	}
}

func (s *SchedulerService) CronExpression() error {

	// this scheduler will be write storage to json using goroutine
	// Seconds field, required
	cron.New(cron.WithSeconds())

	// Seconds field, optional
	cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))

	return nil
}

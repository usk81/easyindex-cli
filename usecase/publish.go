package usecase

import (
	"github.com/usk81/easyindex"
	"github.com/usk81/easyindex/coordinator"
	"github.com/usk81/easyindex/logger"
)

type (
	// PublishResult is the result of publish index request
	PublishResult struct {
		Total int
		Count int
		Skips []coordinator.SkipedPublishRequest
	}

	// PublishFunc is defined Publish usecase function
	PublishFunc func(nt easyindex.NotificationType, urls []string, cf string, limit int, skip bool, ignorePrecheck bool) (result *PublishResult, err error)
)

// Publish is usecase of requesting Google Indexing publish API
func Publish(nt easyindex.NotificationType, urls []string, cf string, limit int, skip bool, ignorePrecheck bool) (result *PublishResult, err error) {
	if len(urls) == 0 {
		return
	}
	rs := make([]coordinator.PublishRequest, len(urls))
	for i, v := range urls {
		rs[i] = coordinator.PublishRequest{
			URL:              v,
			NotificationType: nt,
		}
	}
	l, err := logger.New("debug")
	if err != nil {
		return
	}
	s, err := coordinator.New(coordinator.Config{
		CredentialsFile: &cf,
		Logger:          l,
		Skip:            skip,
	})
	if err != nil {
		return
	}
	total, count, _, skips, err := s.Publish(rs, limit)
	if err != nil {
		return
	}
	result = &PublishResult{
		Total: total,
		Count: count,
		Skips: skips,
	}
	return
}

package usecase

import (
	"github.com/usk81/easyindex"
	"github.com/usk81/easyindex/coordinator"
)

type (
	// PublishResult is the result of publish index request
	PublishResult struct {
		Total int
		Count int
		Skips []coordinator.SkipedPublishRequest
	}

	// PublishFunc is defined Publish usecase function
	PublishFunc func(mgr coordinator.Manager, nt easyindex.NotificationType, urls []string, limit int) (result *PublishResult, err error)
)

// Publish is usecase of requesting Google Indexing publish API
func Publish(mgr coordinator.Manager, nt easyindex.NotificationType, urls []string, limit int) (result *PublishResult, err error) {
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
	total, count, _, skips, err := mgr.Publish(rs, limit)
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

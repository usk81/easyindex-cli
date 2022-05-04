package usecase

import (
	"github.com/usk81/easyindex"
	"github.com/usk81/easyindex/coordinator"
)

type (
	PublishRequest struct {
		NotificationType easyindex.NotificationType `json:"notification_type" csv:"notification_type"`
		URL              string                     `json:"url" csv:"url"`
	}

	// PublishBulkFunc is defined Publish usecase function
	PublishBulkFunc func(mgr coordinator.Manager, rq []*PublishRequest, limit int) (result *PublishResult, err error)
)

// PublishBulk is usecase of requesting Google Indexing publish API
func PublishBulk(mgr coordinator.Manager, rq []*PublishRequest, limit int) (result *PublishResult, err error) {
	if len(rq) == 0 {
		return
	}
	rs := []coordinator.PublishRequest{}
	for i, v := range rq {
		if v != nil {
			rs[i] = coordinator.PublishRequest{
				URL:              v.URL,
				NotificationType: v.NotificationType,
			}
		}
	}
	if len(rs) == 0 {
		return
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

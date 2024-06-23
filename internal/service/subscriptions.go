package service

import (
	"log/slog"
	"sync"

	"github.com/kuromii5/posts/internal/models"
)

// Subscription manager is included in Service struct
type SubscriptionManager struct {
	mu          sync.RWMutex
	subscribers map[uint64][]chan *models.Comment
	log         *slog.Logger
}

func NewSubscriptionManager(log *slog.Logger) *SubscriptionManager {
	return &SubscriptionManager{
		log:         log,
		subscribers: make(map[uint64][]chan *models.Comment),
	}
}

func (m *SubscriptionManager) Subscribe(postID uint64) (<-chan *models.Comment, func()) {
	const f = "Service.SubscriptionManager.Subscribe"

	m.mu.Lock()
	defer m.mu.Unlock()

	log := m.log.With(slog.String("func", f), slog.Uint64("postID", postID))
	log.Info("subscribing to comments for post")

	ch := make(chan *models.Comment, 1)
	m.subscribers[postID] = append(m.subscribers[postID], ch)

	unsubscribe := func() {
		m.mu.Lock()
		defer m.mu.Unlock()

		log := m.log.With(slog.String("func", f), slog.Uint64("postID", postID))
		log.Info("unsubscribing from comments for post")

		subs := m.subscribers[postID]
		for i := range subs {
			if subs[i] == ch {
				m.subscribers[postID] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		close(ch)

		log.Info("unsubscribed successfully")
	}

	log.Info("Subscribed successfully")

	return ch, unsubscribe
}

func (m *SubscriptionManager) Publish(postID uint64, comment *models.Comment) {
	const f = "SubscriptionManager.Publish"

	m.mu.RLock()
	defer m.mu.RUnlock()

	log := m.log.With(slog.String("func", f), slog.Uint64("postID", postID))
	log.Info("publishing comment to subscribers")

	for _, ch := range m.subscribers[postID] {
		select {
		case ch <- comment:
			log.Info("comment published to subscriber")
		default:
			log.Warn("subscriber channel full, dropping comment")
		}
	}

	log.Info("Publishing completed")
}

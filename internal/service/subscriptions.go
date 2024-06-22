package service

import (
	"sync"

	"github.com/kuromii5/posts/internal/models"
)

type SubscriptionManager struct {
	mu          sync.RWMutex
	subscribers map[uint64][]chan *models.Comment
}

func NewSubscriptionManager() *SubscriptionManager {
	return &SubscriptionManager{
		subscribers: make(map[uint64][]chan *models.Comment),
	}
}

func (m *SubscriptionManager) Subscribe(postID uint64) (<-chan *models.Comment, func()) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ch := make(chan *models.Comment, 1)
	m.subscribers[postID] = append(m.subscribers[postID], ch)

	unsubscribe := func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		subs := m.subscribers[postID]
		for i := range subs {
			if subs[i] == ch {
				m.subscribers[postID] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		close(ch)
	}
	return ch, unsubscribe
}

func (m *SubscriptionManager) Publish(postID uint64, comment *models.Comment) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, ch := range m.subscribers[postID] {
		select {
		case ch <- comment:
		default:
			// If the channel is full, we can drop the message or handle it otherwise
		}
	}
}

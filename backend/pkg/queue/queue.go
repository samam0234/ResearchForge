package queue

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/researchforge/backend/internal/model"
)

type EventBus interface {
	PublishEvent(ctx context.Context, projectID string, event model.WebSocketEvent) error
	Subscribe(projectID string) (<-chan model.WebSocketEvent, func())
}

type MemoryEventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan model.WebSocketEvent
}

func NewMemoryEventBus() *MemoryEventBus {
	return &MemoryEventBus{
		subscribers: make(map[string][]chan model.WebSocketEvent),
	}
}

func (b *MemoryEventBus) PublishEvent(ctx context.Context, projectID string, event model.WebSocketEvent) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	channels, ok := b.subscribers[projectID]
	if !ok {
		return nil
	}

	for _, ch := range channels {
		select {
		case ch <- event:
		default:
		}
	}
	return nil
}

func (b *MemoryEventBus) Subscribe(projectID string) (<-chan model.WebSocketEvent, func()) {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan model.WebSocketEvent, 50)
	b.subscribers[projectID] = append(b.subscribers[projectID], ch)

	cleanup := func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		chs := b.subscribers[projectID]
		for i, c := range chs {
			if c == ch {
				b.subscribers[projectID] = append(chs[:i], chs[i+1:]...)
				close(ch)
				break
			}
		}
	}

	return ch, cleanup
}

type RedisQueuePayload struct {
	ProjectID string `json:"project_id"`
}

func MarshalPayload(projectID string) ([]byte, error) {
	return json.Marshal(RedisQueuePayload{ProjectID: projectID})
}

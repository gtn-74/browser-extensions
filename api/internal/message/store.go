package message

import (
	"sync"
	"time"
)

// Store はメモリ上のメッセージ置き場（学習用なのでDBなし）。
// 外に公開する必要があるので大文字始まり(Store)。中の mu/messages/nextID は
// 小文字なのでこのパッケージの外からは触れない（カプセル化）。
type Store struct {
	mu       sync.Mutex
	messages []Message
	nextID   int
}

// NewStore は初期化済みの Store を返す。
func NewStore() *Store {
	return &Store{nextID: 1}
}

// List は現在のメッセージ一覧のコピーを返す。
func (s *Store) List() []Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Message, len(s.messages))
	copy(out, s.messages)
	return out
}

// Add は新しいメッセージを追加して返す。
func (s *Store) Add(text string) Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	m := Message{
		ID:        s.nextID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	s.nextID++
	s.messages = append(s.messages, m)
	return m
}

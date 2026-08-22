// Package message はメッセージのドメイン型とその保管庫を提供する。
// この層は他のどのパッケージにも依存しない（一番内側）。
package message

import "time"

// Message は1件のメッセージを表す。
type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

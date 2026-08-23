// Package server は HTTP のルーティングとハンドラを提供する。
// message パッケージに依存する（server → message の一方向）。
package server

import (
	"encoding/json"
	"net/http"
	"time"

	"api/internal/message"
)

// New はルーティング済み・ミドルウェア適用済みの http.Handler を組み立てて返す。
// 依存(store)は外から渡してもらう＝依存性注入。
func NewRouter(store *message.Store) http.Handler {
	mux := http.NewServeMux()
	// note:HandleFunc:シンプル
	// todo:Handle：状態を持たせたいstructにまとめたい？
	// note:Handler：調査用普段使わないらしい

	// ヘルスチェック
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// メッセージ一覧
	mux.HandleFunc("GET /api/messages", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, store.List())
	})

	// メッセージ追加
	mux.HandleFunc("POST /api/messages", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Text == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "text is required"})
			return
		}
		writeJSON(w, http.StatusCreated, store.Add(body.Text))
	})

	return withLog(withCORS(mux))
	// return withLog(mux)
}

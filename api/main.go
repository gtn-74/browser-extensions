package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Message は1件のメッセージを表す。
type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

// store はメモリ上のメッセージ置き場（学習用なのでDBなし）。
type store struct {
	mu       sync.Mutex
	messages []Message
	nextID   int
}

func newStore() *store {
	return &store{nextID: 1}
}

func (s *store) list() []Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	// nil ではなく空スライスを返して JSON を [] にする
	out := make([]Message, len(s.messages))
	copy(out, s.messages)
	return out
}

func (s *store) add(text string) Message {
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

// writeJSON はステータスコードとともに JSON を書き出すヘルパ。
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// withCORS はどのオリジン（拡張機能の chrome-extension:// も含む）からでも
// 呼べるように CORS ヘッダを付与するミドルウェア。
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// プリフライト（OPTIONS）はここで完結させる
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// withLog はアクセスログを出す簡易ミドルウェア。
func withLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	s := newStore()
	mux := http.NewServeMux()

	// ヘルスチェック
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// メッセージ一覧
	mux.HandleFunc("GET /api/messages", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, s.list())
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
		writeJSON(w, http.StatusCreated, s.add(body.Text))
	})

	handler := withLog(withCORS(mux))

	addr := ":8080"
	log.Printf("listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}

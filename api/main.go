package main

import (
	"api/internal/message"
	"api/internal/server"
	"log"
	"net/http"
)

func main() {
	// 1. データ層を用意
	store := message.NewStore()

	// 2. その store を渡して HTTP ハンドラを組み立て（依存性注入）
	handler := server.New(store)

	// 3. サーバー起動
	addr := ":8080"
	log.Printf("listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}

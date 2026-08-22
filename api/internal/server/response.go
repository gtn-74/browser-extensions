package server

import (
	"encoding/json"
	"net/http"
)

// writeJSON はステータスコードとともに JSON を書き出すヘルパ。
// server パッケージ内でしか使わないので小文字（非公開）。
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

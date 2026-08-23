// victim は「ログインが必要な機密サイト」役のデモサーバー（実験用）。
//
// 目的: 拡張機能の広い host_permissions がどれだけ危険かをローカルで再現する。
// このサーバーは *わざと* CORS ヘッダを一切付けない = ごく普通のサイトのふり。
// それでも、別オリジン（拡張）から Cookie 込みで /secret が読めてしまう、
// という事故を体験するためのもの。
//
// 使い方:
//
//	go run .                         で :9090 で起動
//	ブラウザで http://localhost:9090/ を開き「ログイン」を押す（Cookie付与）
//	→ これで「ログイン済みユーザー」の状態を再現できる
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	cookieName = "session"
	validToken = "super-secret-session-token"
	secretData = `{"email":"you@example.com","balance":1234567,"apiKey":"sk-live-DO-NOT-LEAK"}`
)

func main() {
	mux := http.NewServeMux()

	// トップ: 説明とログイン/ログアウトのリンク
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		loggedIn := hasValidSession(r)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<!doctype html><meta charset="utf-8">
<h1>機密サイト (victim :9090)</h1>
<p>ログイン状態: <b>%v</b></p>
<p><a href="/login">ログインする（Cookieを付与）</a> ｜ <a href="/logout">ログアウト</a></p>
<p><a href="/secret">/secret を直接見る</a>（ログイン中だけ中身が返る）</p>
<hr>
<p>実験: 拡張の sidepanel の Console でこれを実行してみて↓</p>
<pre>fetch("http://localhost:9090/secret", { credentials: "include" })
  .then(r =&gt; r.text()).then(console.log)</pre>`, loggedIn)
	})

	// ログイン: セッションCookieを付与する。
	// SameSite=None + Secure にしているのは、別オリジン(拡張)からの
	// credentials:include なリクエストでも Cookie が送られる状況を再現するため。
	// （http://localhost は Chrome が secure context 扱いなので Secure でも動く）
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    validToken,
			Path:     "/",
			HttpOnly: false, // JSからは読めない = 本来は守られてるはずのCookie
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// ログアウト: Cookieを消す
	mux.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: cookieName, Value: "", Path: "/", MaxAge: -1,
			Secure: true, SameSite: http.SameSiteNoneMode,
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// 機密データ: 正しいCookieがある時だけ返す。CORSヘッダは *付けない*。
	mux.HandleFunc("GET /secret", func(w http.ResponseWriter, r *http.Request) {
		if !hasValidSession(r) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"not logged in"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(secretData))
	})

	addr := ":9090"
	log.Printf("victim(機密サイト役) listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

// hasValidSession はリクエストのCookieが正しいセッションか判定する。
func hasValidSession(r *http.Request) bool {
	c, err := r.Cookie(cookieName)
	return err == nil && c.Value == validToken
}

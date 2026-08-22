# sidepanel-api

[`sidepanel-extension`](../sidepanel-extension) から叩かれる Go 製の小さな API。
標準ライブラリ（`net/http`）のみで、メモリ上にメッセージを保持します。

## 開発環境（Nix）

```sh
nix develop        # go / gopls が入ったシェルに入る
go run .            # http://localhost:8080 で起動
```

## エンドポイント

| Method | Path            | 説明                     |
| ------ | --------------- | ------------------------ |
| GET    | `/api/health`   | 稼働確認                 |
| GET    | `/api/messages` | メッセージ一覧           |
| POST   | `/api/messages` | メッセージ追加（`{"text": "..."}`） |

CORS は全オリジン許可（拡張機能の `chrome-extension://` から叩ける）。

## 動作確認

```sh
curl localhost:8080/api/health
curl -X POST localhost:8080/api/messages -d '{"text":"hello"}'
curl localhost:8080/api/messages
```

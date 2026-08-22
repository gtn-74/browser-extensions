# browser-extensions

ブラウザの右サイド **Side Panel** に UI を表示するブラウザ拡張機能と、その連携用 API をまとめたモノレポ。

## 構成

| ディレクトリ | 技術スタック | 説明 |
| ------------ | ------------ | ---- |
| [`client/`](client) | WXT + React + TanStack Router | Side Panel に UI を表示するブラウザ拡張機能（Chrome / Firefox） |
| [`api/`](api) | Go（標準ライブラリ `net/http`） | 拡張から叩かれる小さな API。メモリ上にメッセージを保持 |

`client` から `api`（`http://localhost:8080`）を `fetch` で呼び出して連携します。

## クイックスタート

いずれも Nix（`flake.nix`）で開発環境を用意しています。

### API（先に起動）

```sh
cd api
nix develop        # go / gopls が入ったシェルに入る
go run .            # http://localhost:8080 で起動
```

### 拡張機能

```sh
cd client
nix develop        # node + pnpm が入ったシェルに入る
pnpm install       # 初回のみ
pnpm dev           # Chrome を起動して拡張を読み込む（HMR付き）
```

Chrome のツールバーの拡張アイコンをクリックすると右サイドにパネルが開きます。

## API エンドポイント

| Method | Path            | 説明 |
| ------ | --------------- | ---- |
| GET    | `/api/health`   | 稼働確認 |
| GET    | `/api/messages` | メッセージ一覧 |
| POST   | `/api/messages` | メッセージ追加（`{"text": "..."}`） |

詳細は各ディレクトリの README を参照してください。

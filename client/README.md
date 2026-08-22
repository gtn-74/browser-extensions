# sidepanel-extension

WXT + React で作るブラウザ拡張機能。ブラウザ右サイドの **Side Panel** に UI を表示し、
[`sidepanel-api`](../sidepanel-api)（Go）と連携します。

## 開発環境（Nix）

```sh
nix develop        # node + pnpm が入ったシェルに入る
pnpm install       # 初回のみ
pnpm dev           # Chrome を起動して拡張を読み込む（HMR付き）
```

## 使い方

1. `sidepanel-api` を起動しておく（`http://localhost:8080`）
2. `pnpm dev` で開いた Chrome のツールバーの拡張アイコンをクリック → 右サイドにパネルが開く
3. パネル内で右クリック →「検証」で DevTools を開くと、Network タブに `fetch` が出る

## 構成

- `entrypoints/sidepanel/` … 右サイドの React UI
- `entrypoints/background.ts` … アイコンクリックでパネルを開く設定
- `wxt.config.ts` … 権限（sidePanel / host_permissions）

# Golang 関連作業メモ

## 環境構築

今回の環境構築には Docker を使用して、基本的に開発もその中で行う形をとった。
動作ログなどは Docker Decktop の各コンテナの「Logs」項目を確認する。

### 構築手順
1. `go mod init InvestSuppoter`
2. Docker + Air で開発環境を構築
    1. .air.toml はコンテナ内で `air init` で作成
    2. compose.yaml, app.local.dockerfile は参考 2. などをもとに作成

### 参考

1. [Go 環境セットアップ DevelopersIO](https://dev.classmethod.jp/articles/go-setup-and-sample/)
2. [Air で始める Go 開発](https://zenn.dev/urakawa_jinsei/articles/a5a222f67a4fac)

### コマンド

-   `docker compose up -d` : コンテナを立ち上げる
-   `docker compose exec app sh` : app コンテナに入る

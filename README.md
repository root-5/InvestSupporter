# Golang 関連作業メモ

## 環境構築

今回の環境構築には Docker を使用して、基本的に開発もその中で行う形をとった。
また、Air を使用したコード改変時ホットリロードと、Godoc でのドキュメント生成を行っている。

Air によって出力されるログなどは Docker Decktop の各コンテナの「Logs」項目から閲覧するか、`docker compose logs app -f`

### 構築手順

1. `go mod init InvestSuppoter`
2. Docker + Air で開発環境を構築
    1. .air.toml はコンテナ内で `air init` で作成
    2. compose.yaml, app.local.dockerfile は参考リンクなどをもとに作成

### コマンド
**基本用途**
-   `docker compose up -d` : コンテナを立ち上げる
-   `docker compose down` : コンテナを停止する
-   `docker compose exec app sh` : app コンテナに入る
-   `docker compose logs app -f` : app コンテナのログを表示する
-   `docker compose exec db sh` : db コンテナに入る
    -   `psql -h 127.0.0.1 -p 5432 -U user db` : db に接続する

**作業用**
-   `go mod tidy` : go.mod に記載されているパッケージを整理する（.goファイルで使われていないパッケージの削除）

## ドキュメント

Godoc を採用しているので、 Docker Compose でコンテナを起動していれば、 [Godoc](https://localhost:8080/) にて、ドキュメントを確認できる。

## 参考リンク集

1. [Go 環境セットアップ DevelopersIO](https://dev.classmethod.jp/articles/go-setup-and-sample/)
2. [Air で始める Go 開発](https://zenn.dev/urakawa_jinsei/articles/a5a222f67a4fac)
3. [J-Quants API について](https://jpx.gitbook.io/j-quants-ja)
4. [GitHub リポジトリ](https://github.com/root-5/InvestSupporter)
5. [godoc の記法まとめ](https://zenn.dev/harachan/articles/db3149c1a19c32)

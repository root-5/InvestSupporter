# Golang 関連作業メモ

## 環境構築

今回の環境構築には Docker を使用して、基本的に開発もその中で行う形をとった。
また、Air を使用したコード改変時ホットリロードと、Godoc でのドキュメント生成を行っている。

Air によって出力されるログなどは Docker Decktop の各コンテナの「Logs」項目から閲覧できる。

### 構築手順

1. `go mod init InvestSupporter`
2. Docker + Air で開発環境を構築
   1. .air.toml はコンテナ内で `air init` で作成
   2. compose.yaml, app.local.dockerfile は参考リンクなどをもとに作成

### コマンド

**基本用途**

- `docker-compose -f="compose.local.yaml" up -d` : （ローカル）全てのコンテナを立ち上げる
- `docker-compose -f="compose.local.yaml" down` : 全てのコンテナを停止する
- `docker-compose up -d` : （本番）全てのコンテナを立ち上げる
- `docker-compose down` : 全てのコンテナを停止する
- `docker-compose down app`: app コンテナだけ停止する
- `docker-compose exec app sh` : app コンテナに入る
- `docker-compose logs app -f` : app コンテナのログを表示する
- `docker-compose rm -fsv app` : app コンテナを削除する
- `docker-compose up -d app` : app コンテナを再起動する
- `docker-compose exec db sh` : db コンテナに入る

  - `psql -h 127.0.0.1 -p 5432 -U user financial_data` : db に接続する

- `curl http://127.0.0.1:8080/financial` : 財務データ取得の確認

**作業用**

- `go mod tidy` : go.mod に記載されているパッケージを整理する（.go ファイルで使われていないパッケージの削除）

**DBバックアップとレストア**

1. `docker-compose exec db bash /var/lib/postgresql/backup/backup.sh`
2. `docker-compose exec db bash /var/lib/postgresql/backup/restore.sh`

**ローカル環境完全リセット**

```bash
docker-compose -f="compose.local.yaml" down -v && \
docker system prune -a && \
sudo rm -rf infra/db/data/
```

- `docker-compose -f="compose.local.yaml" down -v` : 全てのコンテナを停止し、ボリュームも削除する
- `docker system prune -a` : イメージ、コンテナ、ネットワークを全て削除する
- `sudo rm -rf infra/db/data/` : DB のデータを削除する

## ドキュメント

[Godoc](http://localhost:8081/)

Godoc を採用しているので、ローカル環境なら上記のリンクからドキュメントを確認できる。ただし、記載されている関数や変数は大文字から始まるもの（プライベートでないもの）のみが表示される。

## 利用ツール

- [GitHub](https://github.com/root-5/InvestSupporter)
- Docker
- TablePlus

## 参考リンク集

1. [Go 環境セットアップ DevelopersIO](https://dev.classmethod.jp/articles/go-setup-and-sample/)
2. [Air で始める Go 開発](https://zenn.dev/urakawa_jinsei/articles/a5a222f67a4fac)
3. [J-Quants API について](https://jpx.gitbook.io/j-quants-ja)
4. [GitHub リポジトリ](https://github.com/root-5/InvestSupporter)
5. [godoc の記法まとめ](https://zenn.dev/harachan/articles/db3149c1a19c32)

# 開発メモ

## 開発開始時の状態

- Go は今回が初めて
- フレームワークなしでの開発も初めて
- アーキテクチャを強く意識した開発も初めて

## 開発の流れ

1. 特に今の自分では最初から完璧なアーキテクチャ、ディレクトリ構成、関数設計を行うことは困難だった
2. 試したかったテスト駆動は先に関数の完成系がイメージできなければ難しいものだったので一旦棚上げ
   1. 今更思ったが、テストが有効なのはプロダクトの全体が理解できていない人間が改修する際や変更による影響範囲が大きくなった際であって、理解しやすいかつ影響の範囲が小さいプロダクトにおいてはあまり意味がないかもしれない
3. 最初はとにかく書いてはリファクタリングを繰り返した
4. ある程度の構成ができたら、テスト駆動開発に移行したい
5. どんな環境でも `git clone` と env ファイルの設定したうえで `docker-compose up -d` だけで動くようにしたい
6. jquantsAPIからのデータを一次データとして、完全な形でデータベースにローカル保存することも考えたが、これはあまりにもDBが重たくなってしまい移行やコピーなどがしづらくなったり、APIが使える限りAPIを一次データとみなせるたりするので断念した
7. 実際にスプレッドシート側でいろいろ使ってみて、自分の分析に役立つようなデータを返せるよう改善

# アイデア

- セキュリティ強化
  - （api_security.go）
  - 不正なURLパスを叩かれた際に IP を記録しておき、一定回数以上アクセスがあった場合はその IP をブロックする
  - 正規の URL であっても、一定回数以上のアクセスがあった場合はメッセージを出してアクセスを制限する
- GoDoc が使えなくなっている
- structToCsv はほとんど AI 任せなので後で再確認
- 冗長性
  - EC2 インスタンスの起動時に docker-compose が自動で走るように設定（ステートレス化）
    - 現在はサーバーが落ちることがなくなってきたため不要になってきた、後回し
  - 本番環境では app コンテナを 2 つビルドし、片方を通常用、もう片方を通常用が落ちた際のスケジューラー維持用として運用する。DB は一つにする代わりに排他ロックが必要

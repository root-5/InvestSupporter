# InvestSupporter

自らの投資を効率化することを目的としたアプリケーション。
具体的には JquantsAPI 等から取得した全銘柄の財務・株価データをデータベースに利用しやすい形に整形・ストックし、 Google スプレッドシート等から IMPORTDATA 関数で呼び出せる API エンドポイントを提供する。
このリポジトリは Golang のアプリコードとアプリ用・データベース用・監視用の 3 つの Docker コンテナの設定方法を管理している。

## 実装済み機能一覧

- JquantsAPI からのデータ取得、整形保存
  - 上場銘柄一覧
  - 財務データ
  - 調整後始値・終値・安値・高値・出来高
- API エンドポイント提供
  - `/howto` - 使い方説明（WEB ブラウザ）
  - `/financial` - 全銘柄基本情報
  - `/statement?code={{銘柄コード}}` - 全期間財務情報（銘柄コード指定）
  - `/price?code={{銘柄コード}}` - 全期間株価情報（銘柄コード指定）
  - `/price?ymd={{日付}}` - 全銘柄株価情報（日付指定）
  - `/price?code={{銘柄コード}}&ymd={{日付}}` - 株価情報（銘柄コード・日付指定）
  - `/closeprice?code={{銘柄コード複数（カンマ区切り）}}` - 株価終値情報（銘柄コード複数）
  - `/closeprice?code={{銘柄コード複数（カンマ区切り）}}&ymd={{日付}}` - 株価終値情報（銘柄コード複数・日付指定）

## ドキュメント

- [基本設計書](./documents/基本設計書.md)
- [テーブル定義書](./documents/テーブル定義書.md)
- [システム構成](./documents/システム構成.md)
- [J-Quants API について](https://jpx.gitbook.io/j-quants-ja)
- [Godoc](http://localhost:8081/) - ローカル環境専用ドキュメント。ただし、関数や変数はプライベートでないもののみ確認可能。

# 作業メモ

## コマンド

- **基本用途**
  - `docker-compose -f="compose.local.yaml" up -d` : （ローカル）全てのコンテナを立ち上げる
  - `docker-compose -f="compose.local.yaml" down` : 全てのコンテナを停止する
  - `docker-compose -f="compose.local.yaml" down -v` : 全てのコンテナを停止し、ボリュームも削除する
  - `docker-compose up -d` : （本番）全てのコンテナを立ち上げる
  - `docker-compose down` : 全てのコンテナを停止する
  - `docker-compose up -d app` : app コンテナを起動する
  - `docker-compose down app`: app コンテナだけ停止する
  - `docker-compose exec app sh` : app コンテナに入る
  - `docker-compose logs app -f` : app コンテナのログを表示する
  - `psql -h 127.0.0.1 -p 5432 -U user financial_data` : db に接続する
  - `curl http://127.0.0.1:8080/financial` : 財務データ取得の確認
- **DB バックアップとレストア**
  - `docker-compose exec db bash /var/lib/postgresql/backup/backup.sh`
  - `docker-compose exec db bash /var/lib/postgresql/backup/restore.sh`
- **ローカル環境完全リセット**
  ```bash
  docker-compose -f="compose.local.yaml" down -v && \
  docker system prune -a && \
  sudo mv infra/db/data/ infra/db/data_backup_$(date +%Y%m%d%H%M%S)/ && \
  sudo rm -rf infra/db/data/
  ```

## アイデア・修正案

- [ ] JquantsAPI を v1 から v2 に移行
- [ ] インフラの Terraform + GCP 移行
- [ ] CI/CD 導入
- [ ] セキュリティ強化
  - [ ] （api_security.go）
  - [ ] 不正な URL パスを叩かれた際に IP を記録しておき、一定回数以上アクセスがあった場合はその IP をブロックする
  - [ ] 正規の URL であっても、一定回数以上のアクセスがあった場合はメッセージを出してアクセスを制限する
- [ ] GoDoc が使えなくなっている
- [ ] structToCsv はほとんど AI 任せなので後で再確認
- [ ] 冗長性
  - [ ] EC2 インスタンスの起動時に docker-compose が自動で走るように設定（ステートレス化）
    - [ ] 現在はサーバーが落ちることがなくなってきたため不要になってきた、後回し
  - [ ] 本番環境では app コンテナを 2 つビルドし、片方を通常用、もう片方を通常用が落ちた際のスケジューラー維持用として運用する。DB は一つにする代わりに排他ロックが必要

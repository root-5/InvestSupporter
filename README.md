# InvestSupporter

自らの投資を効率化することを目的としたアプリケーション。
具体的には JquantsAPI 等から取得した全銘柄の財務・株価データをデータベースに利用しやすい形に整形・ストックし、 Google スプレッドシート等から IMPORTDATA 関数で呼び出せる API エンドポイントを提供する。
このリポジトリは Golang のアプリコードとアプリ用・データベース用・監視用の 3 つの Docker コンテナの設定コード、GCP 用の Terraform コード等を管理している。

## 実装済み機能一覧

- JquantsAPI からのデータ取得、整形保存
  - 上場銘柄一覧
  - 財務データ
  - 調整後始値・終値・安値・高値・出来高
- API エンドポイント提供
  - `/howto` - 使い方説明（WEB ブラウザ）
  - `/financial` - 全銘柄基本情報
  - `/statement?code={{銘柄コード}}` - 全期間財務情報（銘柄コード指定）
  - `/price?code={{銘柄コード}}&ymd={{日付}}` - 株価情報（銘柄コード・日付を指定可能）
  - `/closeprice?code={{銘柄コード複数（カンマ区切り）}}&ymd={{日付}}` - 株価終値情報（銘柄コード複数・日付を指定可能）

## 使用技術

- インフラ: GCP (Compute Engine), Terraform, Docker
- 言語: Golang, SQL, Shell
- データベース: PostgreSQL
- CI/CD: GitHub Actions
- その他ライブラリ・ツール: GitHub Copilot, J-Quants API, J-Quants MCP

## ドキュメント

- [基本設計書](./documents/基本設計書.md)
- [テーブル定義書](./documents/テーブル定義書.md)
- [システム構成](./documents/システム構成.md)
- [作業メモ](./documents/作業メモ.md)
- [J-Quants API ドキュメント](https://jpx-jquants.com/ja/spec)
- [j-quants-doc-mcp](https://github.com/J-Quants/j-quants-doc-mcp) - v2 移行をほぼ完結させるくらいには便利
- [GoDoc](http://localhost:8081/) - ローカル環境専用、関数や変数はプライベートでないもののみ確認可能

## 作業用

### コマンド

- **基本用途**
  - `docker compose -f="compose.local.yaml" up -d` : 全てのコンテナを立ち上げる
  - `docker compose -f="compose.local.yaml" up -d --force-recreate` : 全てのコンテナを強制的に再作成して立ち上げる
  - `docker compose -f="compose.local.yaml" down` : 全てのコンテナを停止する
  - `docker compose -f="compose.local.yaml" down --rmi all` : 全てのコンテナを停止し、イメージも削除する
  - `docker compose -f="compose.local.yaml" down -v` : 全てのコンテナを停止し、ボリュームも削除する
  - `docker compose -f="compose.local.yaml" logs -f` : ログ閲覧
  - `docker compose -f="compose.local.yaml" exec app-local ash` : app コンテナに入る
  - `docker compose up -d` : （本番）全てのコンテナを立ち上げる
  - `docker compose down` : 全てのコンテナを停止する
  - `psql -h 127.0.0.1 -p 5432 -U user financial_data` : db に接続する
  - `curl http://127.0.0.1:8080/financial` : 財務データ取得の確認
- **DB バックアップとレストア**
  - `docker compose exec db bash /var/lib/postgresql/backup/backup.sh`
  - `docker compose exec db bash /var/lib/postgresql/backup/restore.sh`
- **ローカル環境完全リセット**
  - `docker compose -f="compose.local.yaml" down -v && docker system prune -a` : コンテナ・イメージ・ボリュームを全て削除
  - `sudo mv containers/db/data/ containers/db/data_backup_$(date +%Y%m%d%H%M%S)/ && sudo rm -rf containers/db/data/` : db データをバックアップして削除
- **本番環境接続**
  - `gcloud compute ssh invest-supporter-app --zone asia-northeast1-a --tunnel-through-iap`

### アイデア・修正案

- [x] AWS 関連のコードとインフラを完全削除
  - [x] .pem、手順書の一部等
  - [x] AWS WEB コンソール上でのリソース削除 (VPC、EC2、ネットワーク、ユーザーなど)
- [ ] データの整形処理を追加
- [ ] セキュリティ強化
  - [ ] （api_security.go）
  - [ ] 不正な URL パスを叩かれた際に IP を記録しておき、一定回数以上アクセスがあった場合はその IP をブロックする
  - [ ] 正規の URL であっても、一定回数以上のアクセスがあった場合はメッセージを出してアクセスを制限する
- [ ] structToCsv はほとんど AI 任せなので後で再確認
- [ ] ドメイン取得してエンドポイントを独自ドメインで公開する (優先度低、固定 IP で基本は十分)
- [ ] コンテナまではリポジトリ管理下に置くが、Terraform はプライベートの別リポジトリに分離する方が長期的にはいいかも
- [ ] 高速化
  - [ ] メモリを 2~4GB に増やす
  - [ ] HDD から SSD に変更する
  - [ ] Docker network + localhost 問題？
- [ ] 1449 札幌

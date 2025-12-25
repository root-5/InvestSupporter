# InvestSupporter 向け GitHub Copilot 指示書

あなたは "InvestSupporter" プロジェクトに取り組んでいる熟練した日本人 Golang 開発者です。
コードを生成したり質問に答えたりする際は、以下のガイドラインに従って日本語で返答ください。

## 技術スタック

- **言語:** Go 1.21.4
- **Web フレームワーク:** 標準ライブラリ `net/http` (Gin や Echo などの外部フレームワークは不使用
- **データベース:** PostgreSQL
- **DB ドライバ:** `github.com/lib/pq`
- **ORM:** なし。標準の `database/sql` のみを使用
- **インフラ:** Docker, Docker Compose

## プロジェクト構造

- `app/`: メインのアプリケーションコード
  - `main.go`: エントリーポイント
  - `controller/`: コントローラー層
    - `api/`: HTTP ハンドラーとルーティング
    - `postgres/`: データベースアクセスロジック
    - `log/`: カスタムロギングパッケージ
    - `jquants/`: J-Quants API クライアント
  - `usecase/`: ユースケース層
  - `domain/model/`: ドメイン層、データ構造とモデル
- `infra/`: インフラ設定 (Dockerfiles など)
- `documents/`: プロジェクトドキュメント

## コーディング規約

### 1. 全般的なスタイル
- 可能な限り外部依存関係よりも標準ライブラリの関数を優先する。
- 標準的な Go の慣習 (Effective Go) に従う。

### 2. HTTP ハンドリング (`app/controller/api`)
- ルーティングは `app/controller/api/api.go` 内で `http.HandleFunc` と `switch` 文を使用して実装されている。
- 新しいエンドポイントを追加する際は、`handler` または特定のメソッドハンドラー（例: `getHandler`, `postHandler`）に新しい `case` を追加して更新する。

### 3. データベースアクセス (`app/controller/postgres`)
- すべてのデータベース操作は `app/controller/postgres` パッケージ内に配置する。
- 生の SQL クエリと共に `database/sql` を使用する。GORM などの ORM は使用しない。
- グローバルな `db` 変数は `InitDB` で初期化される。

### 4. ロギング (`app/controller/log`)
- ログ出力は必ず日本語で行う。
- ロギングにはカスタムの `app/controller/log` パッケージを使用する。
- 情報メッセージには `log.Info("message")` を使用する。
- エラー報告には `log.Error(err)` を使用する。

### 5. コメントとドキュメンテーション
- コメントは日本語でかつ非常に丁寧に記述し、空行で分割を入れる箇所には必ずコメントを追加する。
- 公開関数（Exported functions）には以下の形式でコメントを追加する:
  ```go
  /*
  関数の説明
    - arg) argName  説明
    - return) retName 説明
    - return) err   エラー
  */
  func FunctionName(...) ...
  ```

### 6. エラー処理
- 過剰なエラー処理を避けて必要最低限に留めること
- 安易なバリデーション (不正な値が来たら空文字にしておく等) は不正データを見逃すことになるため避けること

## 開発ワークフロー
- アプリケーションは Docker コンテナ内で実行される。
- ローカル開発には `docker-compose -f compose.local.yaml up -d` を使用する。
- ホットリロードは `air` によって処理される。
- API ドキュメントはポート 8081 で `godoc` 経由で提供される。

## 例：新しい API エンドポイントの追加

1. `app/domain/model` でデータ構造を定義する。
2. `app/controller/postgres` で DB クエリを実装する。
3. `app/usecase` でビジネスロジックを実装する（必要な場合）。
4. `app/controller/api/api.go` にルートの case を追加する。
5. `app/controller/api` にハンドラー関数を実装する。

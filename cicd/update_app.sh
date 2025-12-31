#!/usr/bin/env bash
set -euo pipefail

# 本番インスタンス上でアプリケーションを最新化するスクリプト
# GitHub Actions からの実行が基本

APP_DIR="$HOME/InvestSupporter"

cd "$APP_DIR"

# アプリをビルド・再起動
docker compose down -v
docker rmi -f investsupporter-app # app のみ更新する場合
# docker system prune -a -f # 全ての未使用データを確認なく削除
docker compose up -d

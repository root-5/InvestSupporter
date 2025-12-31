#!/usr/bin/env bash
set -euo pipefail

APP_DIR="$HOME/InvestSupporter"

cd "$APP_DIR"

# 最新の main を取得
git fetch origin main
git reset --hard origin/main

# テストのためここで終了
exit 0

# アプリをビルド・再起動
docker compose down -v
docker rmi -f investsupporter-app # app のみ更新する場合
# docker system prune -a -f # 全ての未使用データを確認なく削除
docker compose up -d

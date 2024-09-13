#!/bin/bash

# ======================================================
# Postgres のバックアップを取得・管理するスクリプト
# ======================================================

# -e（エラー発生でシェル終了）、-u（未設定の変数をエラーにする）、-o pipefail（パイプラインで一つでも失敗があればエラーにする）
set -euo pipefail

# バックアップの保存期間（日数）
backup_period=7

# ファイルパス
file_path="/var/lib/postgresql/backup"

# 今日の日付をYYYYMMDD形式で取得
today=$(date '+%Y%m%d')

# 現在時刻をHHMMSS形式で取得
now=$(date '+%H%M%S')

# バックアップを作成
echo ""
echo "バックアップを開始します"
pg_dump -Fc -h 127.0.0.1 -p 5432 -U ${POSTGRES_USER} -d ${POSTGRES_DB} > ${file_path}/${POSTGRES_DB}_${today}_${now}.dump
echo "バックアップが完了しました ＞ ${file_path}/${POSTGRES_DB}_${today}_${now}.dump"
echo ""

# 7日以上前のバックアップファイルを削除
echo "7日以上前のバックアップファイルを削除します"
find ${file_path} -name "${POSTGRES_DB}_*.dump" -type f -mtime +${backup_period} -exec rm -f {} \;
echo "7日以上前のバックアップファイルを削除しました"

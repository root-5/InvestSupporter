#!/bin/bash

# ======================================================
# Postgres のバックアップから復元するスクリプト
# ======================================================

# -e（エラー発生でシェル終了）、-u（未設定の変数をエラーにする）、-o pipefail（パイプラインで一つでも失敗があればエラーにする）
set -euo pipefail

# ファイルパス
file_path="/var/lib/postgresql/backup"

# フォルダ内のバックアップファイル（.dump）をリストアップ
echo "バックアップファイルをリストアップします"
ls -l ${file_path}/*.dump
echo ""
echo "復元したいバックアップファイルを入力してください"

# バックアップファイルを選択
read -p "バックアップファイル名（例：sample_20210101_000000.dump）: " backup_file

# バックアップファイルが存在するか確認
if [ ! -f ${file_path}/${backup_file} ]; then
  echo "バックアップファイルが存在しません ＞ ${file_path}/${backup_file}"
  exit 1
fi

# バックアップファイルを復元
echo ""
echo "バックアップファイルを復元します"
pg_restore --clean -h 127.0.0.1 -p 5432 -U ${POSTGRES_USER} -d ${POSTGRES_DB} ${file_path}/${backup_file}

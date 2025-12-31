#!/usr/bin/env bash
set -euo pipefail

# 各コンテナ用の .env と containers/db/backup 配下の .dump を IAP 経由で本番インスタンスの同一パスへ同期するスクリプト
# 実行: `bash cicd/sync_hidden_files.sh PROJECT_ID=your-gcp-project-id`

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
INSTANCE_NAME=${INSTANCE_NAME:-invest-supporter-app}
ZONE=${ZONE:-asia-northeast1-a}
PROJECT_ID=${PROJECT_ID:-}

# 環境変数風の引数 (PROJECT_ID=..., INSTANCE_NAME=..., ZONE=...) を許容
for arg in "$@"; do
	case "$arg" in
		PROJECT_ID=*) PROJECT_ID=${arg#*=} ;;
		INSTANCE_NAME=*) INSTANCE_NAME=${arg#*=} ;;
		ZONE=*) ZONE=${arg#*=} ;;
		*) ;;
	esac
done

if [[ -z "$PROJECT_ID" ]]; then
	echo "PROJECT_ID が設定されていません。" >&2
	exit 1
fi

cd "$ROOT_DIR"

# 対象ファイルのリストを生成
#  - .env / *.env
#  - containers/db/backup 配下の *.dump
mapfile -d '' FILES < <(
	find . -type f \
		\( -name '.env' -o -name '*.env' -o -path './containers/db/backup/*.dump' \) \
		-print0 2>/dev/null
)

if [[ ${#FILES[@]} -eq 0 ]]; then
	echo "同期対象の .env / .dump が見つかりません" >&2
	exit 0
fi

echo "同期対象ファイル:"
for file in "${FILES[@]}"; do
    echo " - $file"
done

# tar で相対パスのまま圧縮し、リモートで展開する
# リモート側は /home/appuser/InvestSupporter 配下に展開
tar --null -czf - --files-from <(printf '%s\0' "${FILES[@]}") |
	gcloud compute ssh "$INSTANCE_NAME" \
		--zone="$ZONE" \
		--project="$PROJECT_ID" \
		--tunnel-through-iap \
		--quiet \
		-- -T "sudo -u appuser bash -c 'cd /home/appuser/InvestSupporter && tar xzf -'"
echo "同期完了"

# dump ファイルがない場合はここで終了
dump_file_name=""
for file in "${FILES[@]}"; do
	if [[ "$file" == ./containers/db/backup/*.dump ]]; then
		dump_file_name=$(basename "$file")
		break
	fi
done
if [[ -z "$dump_file_name" ]]; then
	exit 0
fi

# db コンテナを起動、バックアップ復元スクリプトを実行
echo "DB コンテナを起動し、バックアップ復元スクリプトを実行します"
gcloud compute ssh "$INSTANCE_NAME" \
	--zone="$ZONE" \
	--project="$PROJECT_ID" \
	--tunnel-through-iap \
	--quiet \
	--command="sudo -u appuser bash -c 'cd /home/appuser/InvestSupporter && docker compose up -d db && echo \"$dump_file_name | docker compose exec db bash /var/lib/postgresql/backup/restore.sh\"'"
echo "DB バックアップの復元が完了しました"

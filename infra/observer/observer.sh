# メモリ監視関数
function memory_logger() {
    while true; do
        # メモリ情報の取得
        MEMORY_USAGE_ROW=$(free | grep Mem)

        # メモリ使用率の計算
        MEMORY_TOTAL=$(echo $MEMORY_USAGE_ROW | awk '{print $2}')
        MEMORY_USED=$(echo $MEMORY_USAGE_ROW | awk '{print $3}')
        MEMORY_USAGE=$(echo "scale=2; $MEMORY_USED / $MEMORY_TOTAL * 100" | bc)

        # ログ出力
        MEMORY_USAGE_INT=${MEMORY_USAGE%.*}
        TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')
        echo "$TIMESTAMP - Used: $MEMORY_USED Total:$MEMORY_TOTAL ($MEMORY_USAGE_INT%)" >> $MEMORY_LOGFILE

        # 閾値を超えた場合はアラート
        if [ $MEMORY_USAGE_INT -gt $MEMORY_USAGE_THRESHOLD ]; then
            echo "$TIMESTAMP - ALERT > Memory Usage $MEMORY_USAGE_INT%" >> $LOGFILE
        fi

        # 待機
        sleep $MEMORY_INTERVAL
    done
}

# API 死活監視関数
function api_logger() {
    while true; do
        # APIのステータスコード取得
        STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}" $API_URL)

        # ログ出力
        TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')
        echo "$TIMESTAMP - API Status Code $STATUS_CODE" >> $API_LOGFILE

        # ステータスコードの判定してアラート
        if [ $STATUS_CODE -ne 200 ]; then
            echo "$TIMESTAMP - ALERT > API Status Code $STATUS_CODE" >> $LOGFILE
        fi

        # 待機
        sleep $API_INTERVAL
    done
}

# 一定行数を超えたログレコードの削除
function log_cleaner() {
    while true; do

    # /log ディレクトリのログファイルを取得
    LOG_FILES=$(ls /log)

    for LOG_FILE in $LOG_FILES; do
        # ログファイルの行数を取得
        LOG_ROWS=$(wc -l /log/$LOG_FILE | awk '{print $1}')

        # 行数が閾値を超えた場合は古いログを削除
        if [ $LOG_ROWS -gt $LOG_ROWS_MAX ]; then
            # 超過した行数を計算
            LOG_DELETE_LINES=$(($LOG_ROWS - $LOG_ROWS_MAX))

            # ログファイルの先頭から削除行数分を削除
            sed -i "1,$LOG_DELETE_LINES d" /log/$LOG_FILE
        fi
    done

    # 待機
    sleep $LOG_CLEAN_INTERVAL
    done

}

# API 死活監視関数をバックグラウンドで実行
api_logger &
# ログクリーナー関数をバックグラウンドで実行
log_cleaner &

# メモリ監視関数を実行
memory_logger

FROM postgres:18-alpine

# バックアップ用のディレクトリを作成
RUN mkdir -p /var/lib/postgresql/backup

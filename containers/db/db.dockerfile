FROM postgres:18-alpine

WORKDIR /var/lib/postgresql/data

# バックアップ用のディレクトリを作成
RUN mkdir -p /var/lib/postgresql/backup

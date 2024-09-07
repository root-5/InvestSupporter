FROM postgres:alpine

WORKDIR /var/lib/postgresql/data

# バックアップ用のディレクトリを作成
RUN mkdir -p /var/lib/postgresql/backup

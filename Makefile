SERVICE_NAME=go-moc-app
APP_CONTAINER=go-app-server

all: build up ps
reset: down prune build up ps 

# 未使用のリソースを削除
prune:
	docker system prune -f
# ビルド
build:
	docker-compose build
# 起動
up:
	docker-compose up -d
# 確認
ps:
	docker-compose ps
# 停止
down:
	docker-compose down -v
# アプリコンテナにログイン
login_app:
	docker exec -it ${APP_CONTAINER} /bin/sh
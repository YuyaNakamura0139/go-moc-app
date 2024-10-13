# go-moc-app

[ジブリAPI](https://ghibliapi.vercel.app/#)からデータを取得して表示するアプリを作成する。

## 使い方

### 0. Docker Desktopのインストール

https://www.docker.com/ja-jp/products/docker-desktop/

インストール完了後、以下のコマンドを実行し結果が返ってくればOK

```zsh
nakamura@nakamura go-moc-app % docker --version
Docker version 20.10.12, build e91ed57

nakamura@nakamura go-moc-app % docker compose version 
Docker Compose version v2.2.3
```

### 1. リポジトリをクローン

```zsh: クローン
git clone https://github.com/YuyaNakamura0139/go-moc-app.git
```

### 2. ディレクトリに移動

```zsh: クローンしてきたディレクトリに移動
cd ~/go-moc-app
```

### 2. コンテナを起動

ビルドで時間がかかる可能性あり(10分くらい)

```zsh: コンテナを起動
make all
```

最終的に以下が表示されれば成功

```zsh: ログ
Creating go-app-server ... done
docker-compose ps
    Name        Command   State           Ports         
--------------------------------------------------------
go-app-server   /bin/sh   Up      0.0.0.0:8081->8081/tcp
```

### 3. コンテナにログイン

```zsh: コンテナにログイン
make login_app
```

### 4. シェルに入る

```zsh: 
nakamura@nakamura go-moc-app % make login_app
docker exec -it go-app-server /bin/sh
```

上記実行するとシェルに入れる。試しにgoのバージョンを確認してみる。

```sh: goのバージョンを確認
/go/src/work # go version
go version go1.22.8 linux/amd64
```

あとは自由に開発する。

### 5. コンテナの停止

開発し終わったらコンテナを停止しておく。

```zsh: コンテナの停止
make down
```

## [makeコマンド](https://www.gnu.org/software/make/manual/make.html#Rules)

```txt: makeコマンド
make all: ビルド~コンテナ起動~コンテナの状態の表示
make reset: コンテナの停止~未使用リソースの削除~ビルド~コンテナ起動~コンテナの状態の表示
make prune: 未使用のリソースを削除
make build: ビルド
make up: コンテナ起動
make ps: コンテナの状態の表示
make down: コンテナの停止
make login_app: appコンテナにログイン
make login_db: dbコンテナにログイン
```

[参考](https://qiita.com/mizcii/items/cfbd2aa17f6b7517c37f)
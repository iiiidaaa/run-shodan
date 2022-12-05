# Usage
## Build
```shell
make build
```

上記コマンド成功後、`.env`ファイルが作成されます。  
各サービスのAPIキーを`.env`ファイルに記入してください  
`DB_~~`は修正すると動かなくなると思います

## Exec DB
```shell
docker-compose up -d
```

metabase,db立ち上げ後、`localhost:3000`でmetabaseにアクセスできます。  
初回アクセス時は、metabaseの初期設定が必要です。  
以下の通りに入力してください。
- 言語: お好みで (以下Japaneseを選択した時の表示で説明します)
- 名前、ユーザー情報: お好みで メアドとパスワードはmetabaseのログインに必要になるはず。。
- データの追加
  - MySQLを選択し、以下の通りに入力
  - 表示名: osint お好みで
  - Host: osint-db
  - Port: 3306 未入力も可
  - Database Name: osint
  - username: root
  - password: 未入力
  - その他チェック入れない
- トラッキング: お好みで
上記入力で初期設定完了です。

## Run
上記のmetabase,dbが立ち上がっていることが前提です。(結果の保存先になります)
```shell
docker run --rm --network run-shodan_osint --env-file .env run-shodan {target}
```

## 環境の片付け方
db,metabase
```shell
docker-compose stop
```
データを削除するためには、以下を行う
```shell
docker-compose down
```
この後、docker-compose立ち上げ時に作られたmysqlディレクトリを削除してください。

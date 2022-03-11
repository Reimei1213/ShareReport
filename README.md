# ShareReport

Userサービスのプロトコルバッファからコードを生成するコマンド
```bash
$ cd backend/proto/user
$ ./generate.sh 
```

DockerのMySQLからバックアップするコマンド
```bash
$ docker exec -it コンテナ名 mysqldump データベース名 -uroot -pパスワード > ./dump.sql
```
# server-redis-first

## command

- 前提
    - redis-server --port 7079

### initializing

```bash
mkdir server-redis-first
cd server-redis-first
go mod init linnefromice/server-redis-first
# -> XXX.go ファイルでimportを記載
go build
```

### use method in other file

```bash
# -> main.go で同じパッケージのメソッドを呼び出す
go build
./server-redis-first
```

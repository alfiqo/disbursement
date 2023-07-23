## Getting started


Clone Project
```
$ git clone git@github.com:alfiqo/disbursement.git

$ cd disbursement && go mod tidy

$ cp .env.example .env
```
For running stack
```
$ docker compose up -d
```
Import DB
```
$ docker exec -i disbursement-db-1 mysql -uuser -ppass ewallet < ewallet.sql
```
Run App
```
go run main.go
```
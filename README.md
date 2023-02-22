# daily-attendance

## project

Automatic sign-in daily for multi platforms

### cron

test:
go run cmd/main.go

prod:
DAILY_ENV=prod go run cmd/main.go

### cmd

go run cmd/main.go cmd

## remark

1. go version

2. *Ck exists in the project root directory which has cookie string

## todo

1. mail
2. error handle
3. cron

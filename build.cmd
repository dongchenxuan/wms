@echo off

SET FileName=wms

cd server

echo set go build env GOOS=linux,GOARCH=amd64,GO111MODULE=off
go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64

echo go mod tidy
go mod tidy

echo building
go build -o wms main.go
go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64






REM echo copy to remote host
REM scp -P 22 dly-cmis root@192.168.1.56:/home/update/cmis

REM scp -P 22 cfg.json root@192.168.1.56:/home/cmis

REM scp -P 22 dly-cmis root@192.168.1.56:/home/cmis
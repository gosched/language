# Overview

## Official

- https://golang.org/ref/spec
- https://golang.org/pkg/

## Level

- repository / module / mod
- branch / production
- tag / version
- package
- interface
- struct
- func

### version

- git tag
- git show tag_name

- git tag tag_name commit
- git tag tag_name commit -a -m "message"

- git tag -d tag_name

- go get github.com/username/module@v2.0.1

### module

go.mod

## Object oriented programming

偽繼承 組合結構體 embedding

多態 支持方法重寫（重寫父類的方法）

多態 不支持 Overload (參數不同 返回類型相同或不同)

## Initialize

mod -> package -> var -> init

## Network

- https://golang.org/pkg/net/http/
- https://golang.org/pkg/net/
- https://godoc.org/golang.org/x/net/ipv6
- https://godoc.org/golang.org/x/oauth2

## String

```
strconv.Atoi
strconv.Itoa
strconv.ParseBool
strconv.ParseFloat
strconv.ParseInt
strconv.ParseUint
strconv.FormatBool
strconv.FormatFloat
strconv.FormatInt
strconv.FormatUint
```

## Command

- go mod init
- go run main.go
- go run *.go
- go build .
- go build -o target
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build *.go
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build *.go
- go vet main.go
- go list -json std

## Visual Studio Code

- Command + Shift + P
- Go: Install/Update tools

## Other

https://github.com/golang/go/wiki/Iota

https://github.com/golang/go/wiki/LockOSThread
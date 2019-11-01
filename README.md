# w01fb0ss常规go项目搭建

## 目录结构
```shell
webapi
├── README.md
├── config
│   └── config.go
├── config.yaml
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── jwt
│   │   └── jwt.go
│   └── log
│       └── log.go
├── models
│   ├── auth.go
│   ├── models.go
│   └── test.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   └── util
│       └── jwt.go
├── routers
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       └── test.go
│   └── router.go
└── table.sql
```

## 配置文件说明（config.yaml）
```shell
cat config.yaml
```
```shell
AppName: "common"
RunMode: "debug"  # debug release

App:
  JwtSecret: "WdZuIsaTOMkPdbmcFpjP4SpnpVATdVr7b77JTybbhEw36bOAf" # cat /dev/urandom | tr -dc A-Za-z0-9 | head -c 49;echo

Server:
  Port: 8888

DB:
  User: "root"
  Password: "toor"
  Host: "127.0.0.1:3306" # IP:Port
  Name: "common"
  TablePrefix: "common_"

Log:
  Type: "text" # json text
  Level: "info" # trace debug info warn error fatal panic 默认info

# ...
```

## 参考
[go-gin-example](https://github.com/EDDYCJY/go-gin-example)
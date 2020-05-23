module github.com/xiaobudongzhang/micro-auth

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /wwwroot/microdemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /wwwroot/microdemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /wwwroot/microdemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /wwwroot/microdemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /wwwroot/microdemo/micro-plugins

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.5.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0 // indirect
	github.com/micro/micro/v2 v2.4.0
	github.com/xiaobudongzhang/micro-basic v1.1.2
	github.com/xiaobudongzhang/micro-plugins v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.22.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	honnef.co/go/tools v0.0.1-2019.2.3
)

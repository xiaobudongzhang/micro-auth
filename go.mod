module github.com/xiaobudongzhang/micro-auth

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /data/ndemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /data/ndemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /data/ndemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /data/ndemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /data/ndemo/micro-plugins

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200510134214-a0f89cb675dc // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/micro/v2 v2.4.0
	github.com/xiaobudongzhang/micro-basic v1.1.2
	github.com/xiaobudongzhang/micro-plugins v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.22.0
	honnef.co/go/tools v0.0.1-2019.2.3
)

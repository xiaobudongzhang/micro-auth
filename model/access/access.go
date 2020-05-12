package access

import (
	"fmt"
	"sync"

	r "github.com/go-redis/redis"
	"github.com/micro/xiaobudongzhang/micro-plugins/jwt"
	"github.com/micro/xiaobudongzhang/micro-plugins/redis"
)

var (
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	cfg = &jwt.Jwt{}
)

type service struct {
}

type Service interface {
	MakeAccessToken(subject *Subject) (ret string, err error)
	GetCachedAccessToken(subject *Subject) (ret string, err error)
	DelUserAccessToken(token string) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("getservice 初始化失败")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	ca = redis.GetRedis()

	fmt.Println("ca ", ca)
	s = &service{}
}

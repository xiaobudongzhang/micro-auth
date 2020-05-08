package handler

import (
	"context"

	"strconv"

	"github.com/xiaobudongzhang/micro-auth/model/access"

	"github.com/micro/go-micro/v2/util/log"
	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"
)

var (
	accessService access.Service
)

func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("init handler error %s", err)
		return
	}
}

type Service struct{}

func (s *Service) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Log("create toke")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(req.UserId, 10),
		Name: req.UserName,
	})

	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("token 生成失败 %s", err)
		return err
	}
	rsp.Token = token
	return nil
}

func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Log("清除token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("del token fail %s", err)
		return err
	}
	return nil
}

func (s *Service) GetCachedAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Logf("[GetCachedAccessToken] 获取缓存的token, %d", req.UserId)

	token, err := accessService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatInt(req.UserId, 10),
	})

	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[GetCachedAccessToken] 获取缓存的token失败, err:%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

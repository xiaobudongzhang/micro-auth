package access

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *service) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()

	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "book.micto.mu",
		Subject:   subject.ID,
	}
	return
}

func (s *service) saveTokenToCache(subject *Subject, val string) (err error) {

	err = ca.Set(tokenIDKeyPrefix+subject.ID, val, tokenExpiredDate).Err()

	if err != nil {
		return fmt.Errorf("保存token失败" + err.Error())
	}
	return
}

func (s *service) delTokenFromCache(subject *Subject) (err error) {
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("del token fail" + err.Error())
	}
	return
}

func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.ID).Result()
	if err != nil {
		return token, fmt.Errorf("get token error %s", err)
	}
	return string(tokenCached), nil
}

func (s *service) parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("不合法的token%v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}
		return nil, fmt.Errorf("不合法token %s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("不合法的token")
	}
	return mapClaimsToJwClaim(claims), nil
}

func mapClaimsToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}
	return jC
}

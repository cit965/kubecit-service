package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// secretKey jwt 密钥因为后端开源从环境变量读取
var secretKey = []byte(os.Getenv("KUBECIT_USER_TOKEN_SECRET"))

// 过期时间24小时
const tokenExpireTime = 3600 * 24

type LoginClaims struct {
	RegisteredClaims jwt.RegisteredClaims
	UserID           uint64
	RoleID           uint8
	ExpireTime       int64
}

func (loginClaims *LoginClaims) Valid() error {
	if err := loginClaims.RegisteredClaims.Valid(); err != nil {
		return err
	}
	return nil
}

// GenerateToken 生成jwt token
func GenerateToken(userID uint64, roleID uint8) (string, int64) {
	expireTime := time.Now().Unix() + tokenExpireTime
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": jwt.SigningMethodHS256.Alg(),
		},
		Method: jwt.SigningMethodHS256,
		Claims: &LoginClaims{
			UserID:     userID,
			RoleID:     roleID,
			ExpireTime: expireTime,
		},
	}
	tokenStr, _ := token.SignedString(secretKey)
	return tokenStr, expireTime
}

// VerifyToken 验证jwt token
func VerifyToken(tokenStr string) (*LoginClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&LoginClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	loginClaims, ok := token.Claims.(*LoginClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token valid failed")
	}
	if loginClaims.ExpireTime <= time.Now().Unix() {
		return nil, errors.New("token time failed")
	}

	return loginClaims, nil

}

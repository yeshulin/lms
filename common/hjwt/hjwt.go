package hjwt

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte(beego.AppConfig.String("appKey"))
)

// 产生json web token
func GenToken() string {
	fmt.Println(key)
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "hzwy23",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ss
}

// 校验token是否有效

func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

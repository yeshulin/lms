package hjwt

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte(beego.AppConfig.String("jwtkey"))
)

// 产生json web token
func GenToken(id int, username string, realname string, email string, phone string) string {
	fmt.Println(key)
	//	claims := &jwt.StandardClaims{
	//		NotBefore: int64(time.Now().Unix()),
	//		ExpiresAt: int64(time.Now().Unix() + 14400),
	//		Issuer:    "hzwy23",
	//	}
	now := time.Now()
	exp := now.Add(time.Hour * 24).Unix()

	iat := now.Unix()
	claims := &jwt.MapClaims{
		"sub":      id,
		"iat":      iat,
		"exp":      exp,
		"username": username,
		"email":    email,
		"phone":    phone,
		"realname": realname,
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

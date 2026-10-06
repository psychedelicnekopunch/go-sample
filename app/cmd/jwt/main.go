
package main


import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)



func main() {

	const SECRET_KRY = "SECRET_KRY"
	key := []byte(SECRET_KRY)
	// invalidKEY := []byte("INVALID_SECRET_KRY")


	value := jwt.MapClaims{
		"key": "value",
		"test": "test value",
	}


	fmt.Print("\n== value ==\n", value, "\n=====\n")


	// JWS認証用のトークンを生成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, value)
	token, err := t.SignedString(key)
	if err != nil {
		fmt.Print("\n== EEROR ==\n", err.Error(), "\n=====\n")
		return
	}


	fmt.Print("\n== token ==\n", token, "\n=====\n")


	// JWS認証用のトークンから元データに返還
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }
		// return invalidKEY, nil
		return key, nil
	})
	if err != nil {
		fmt.Print("\n== EEROR ==\n", err.Error(), "\n=====\n")
		return
	}


	fmt.Print("\n== parsedToken ==\n", parsedToken, "\n=====\n")
	fmt.Print("\n== parsedToken.Valid ==\n", parsedToken.Valid, "\n=====\n")
	fmt.Print("\n== parsedToken.Claims ==\n", parsedToken.Claims, "\n=====\n")

	// claims := parsedToken.Claims.(jwt.Claims)
	claims := parsedToken.Claims.(jwt.MapClaims)
	res, ok := claims["test"].(string)
	if !ok {
		return
	}
	// fmt.Print("\n== ok ==\n", ok, "\n=====\n")
	fmt.Print("\n== res ==\n", res, "\n=====\n")


	fmt.Print("\n== key: value ==\n")
	for key, value := range claims {
		fmt.Print(key, ": ", value, "\n")
	}
	fmt.Print("=====\n")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var myKey = []byte("iniadalahfrasarahasiaku")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GeneratorJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)

}

func GeneratorJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jhoni Torina"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(myKey)

	if err != nil {
		fmt.Errorf("terjadi kesalahan: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("client")

	handleRequests()
}

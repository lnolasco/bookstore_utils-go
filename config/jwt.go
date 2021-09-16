package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//SecretKey key used to sign the JWT
	SecretKey []byte

	// ExpiresToken time to expire the JWT token
	ExpiresToken uint64

	// TokenType is the accepted token type
	TokenType = ""
)

func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	string64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println("New Secret:", string64)
}

func ChargeJwt() {
	var err error
	if godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	TokenType = os.Getenv("TOKEN_TYPE")
	ExpiresToken, err = strconv.ParseUint(os.Getenv("EXPIRES_TOKEN"), 10, 64)
	if err != nil {
		ExpiresToken = 3600
	}

}

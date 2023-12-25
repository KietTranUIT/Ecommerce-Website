package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"time"
	"user-service/internal/core/model/request"
)

func DecodeRequest(r io.Reader, object interface{}) any {
	decoder := json.NewDecoder(r)

	switch value := object.(type) {
	case request.SignUpRequest:
		decoder.Decode(&value)
		fmt.Println(value)
		return value
	default:
		return nil
	}
}

func RandomCode() string {
	var code string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func ExpireAt() time.Time {
	now := time.Now()
	return now.Add(time.Minute * 5)
}

func HashPassword(password string) string {
	hasher := sha256.New()

	hasher.Write([]byte(password))

	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

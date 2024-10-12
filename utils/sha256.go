package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SHA256Encode(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	tmpstr := h.Sum(nil)
	return hex.EncodeToString(tmpstr)
}

// 加密操作
func MakePassword(plainpwd, salt string) string {
	//fmt.Println("plainpwd: ", plainpwd)
	//fmt.Println("salt: ", salt)
	return SHA256Encode(plainpwd + salt)
}

func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

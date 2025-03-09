package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
	"time"

	rand2 "math/rand"
)

func GenerateInviteCode() (string, error) {
	// 生成随机前缀 8 或 9
	firstDigit, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return "", err
	}
	prefix := "8"
	if firstDigit.Int64() == 1 {
		prefix = "9"
	}

	// 生成随机后 6 位数
	bytes := make([]byte, 3) // 3 字节 = 6 个十六进制字符
	_, err = rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// 转换成 6 位十六进制字符，并拼接前缀
	code := prefix + hex.EncodeToString(bytes)[:6]

	return code, nil
}

// RandString generate rand string with specified length
func RandString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	data := []byte(str)
	var result []byte
	r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, data[r.Intn(len(data))])
	}
	return string(result)
}

func GenPassword(pass string, salt string) string {
	data := []byte(pass + salt)
	hash := sha3.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

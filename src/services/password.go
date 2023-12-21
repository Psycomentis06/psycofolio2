package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

const saltLength = 8
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HashPassword(pass string) string {
	salt1 := StringWithCharset(saltLength, charset)
	salt2 := StringWithCharset(saltLength, charset)

	for !isSaltsValid(salt1, salt2) {
		salt1 = StringWithCharset(saltLength, charset)
		salt2 = StringWithCharset(saltLength, charset)
	}

	firstSalt, secondSalt := orderSalts(salt1, salt2)

	hasher := sha256.New()
	hasher.Write([]byte(firstSalt + pass))
	hashed := hex.EncodeToString(hasher.Sum(nil))

	hasher.Reset()
	hasher.Write([]byte(secondSalt + hashed))
	hashed = hex.EncodeToString(hasher.Sum(nil))

	return salt1 + "@" + salt2 + "@" + hashed
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func isSaltsValid(salt1, salt2 string) bool {
	for i := len(salt1) - 1; i >= 0; i-- {
		if salt1[i]%2 != salt2[i]%2 {
			return true
		}
	}
	return false
}

func orderSalts(salt1, salt2 string) (string, string) {
	for i := len(salt1) - 1; i >= 0; i-- {
		if salt1[i]%2 == 0 {
			return salt1, salt2
		}
		return salt2, salt1
	}
	return "", ""
}

func VerifyPassword(hashed string, raw string) error {
	saltsAndHash := strings.Split(hashed, "@")
	if len(saltsAndHash) != 3 {
		return errors.New("Malformed password")
	}

	salt1, salt2, correctHash := saltsAndHash[0], saltsAndHash[1], saltsAndHash[2]
	if len(salt1) != saltLength || len(salt2) != saltLength {
		return errors.New(fmt.Sprintf("slat length should be equal to %d", saltLength))
	}

	firstSalt, secondSalt := orderSalts(salt1, salt2)

	hasher := sha256.New()
	hasher.Write([]byte(firstSalt + raw))
	hashedStr := hex.EncodeToString(hasher.Sum(nil))

	hasher.Reset()
	hasher.Write([]byte(secondSalt + hashedStr))
	hashedStr = hex.EncodeToString(hasher.Sum(nil))

	if hashedStr != correctHash {
		return errors.New("Given password doesn not match hashed password")
	}
	return nil
}

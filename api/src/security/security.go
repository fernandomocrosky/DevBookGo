package security

import "golang.org/x/crypto/bcrypt"

func Hash(key string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
}

func CheckHash(hashedKey, key string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(key))
}

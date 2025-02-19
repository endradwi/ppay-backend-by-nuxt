package lib

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pilinux/argon2"
)

func CreateHash(password string) string {
	godotenv.Load()
	hash, _ := argon2.CreateHash(password, os.Getenv("HASH_SECRET"), argon2.DefaultParams)
	return hash
}

func GenerateTokenArgon(password string, hash string) bool {
	godotenv.Load()
	match, _ := argon2.ComparePasswordAndHash(password, os.Getenv("HASH_SECRET"), hash)
	return match
}

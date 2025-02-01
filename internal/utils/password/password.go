package password

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    saltedPassword := addSecretSalt(password)
    
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    
    return string(hashedBytes), nil
}

func VerifyPassword(hashedPwd, plainPwd string) error {
    saltedPassword := addSecretSalt(plainPwd)
    
    err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(saltedPassword))
    if err != nil {
        return fmt.Errorf("invalid password: %w", err)
    }
    
    return nil
}

func addSecretSalt(password string) string {
    h := hmac.New(sha256.New, []byte(os.Getenv("PASSWORD_SECRET")))
    h.Write([]byte(password))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
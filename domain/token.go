package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/techx/portal/config"
)

// GenerateToken generates a JWT token for a user
func GenerateToken(subject string, authConfig config.Auth) (string, error) {
	now := time.Now()
	expirationTime := now.Add(authConfig.AuthSoftExpiryDuration)
	encryptedSubject, err := encrypt(subject, authConfig.CipherKey)
	if err != nil {
		return "", err
	}

	claims := &jwt.StandardClaims{
		Id:        authConfig.AuthIssuerUUID,
		Issuer:    authConfig.AuthIssuer,
		Audience:  authConfig.AuthAudience,
		Subject:   encryptedSubject,
		IssuedAt:  now.Unix(),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(authConfig.AccessTokenSecret))
}

// VerifyToken verifies a JWT token and returns the user's phone number
func VerifyToken(authConfig config.Auth, tokenStr, _ string) (string, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(authConfig.AccessTokenSecret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.NewValidationError("token is invalid", jwt.ValidationErrorSignatureInvalid)
	}

	if !claims.VerifyIssuedAt(time.Now().Add(authConfig.AuthHardExpiryDuration).Unix(), true) {
		return "", jwt.NewValidationError("token is expired", jwt.ValidationErrorClaimsInvalid)
	}

	if !claims.VerifyAudience(authConfig.AuthAudience, true) {
		return "", jwt.NewValidationError("invalid audience", jwt.ValidationErrorClaimsInvalid)
	}

	if !claims.VerifyIssuer(authConfig.AuthIssuer, true) {
		return "", jwt.NewValidationError("invalid issuer name", jwt.ValidationErrorClaimsInvalid)
	}

	if claims.Id != authConfig.AuthIssuerUUID {
		return "", jwt.NewValidationError("invalid issuer uuid", jwt.ValidationErrorClaimsInvalid)
	}

	decryptedSubject, err := decrypt(claims.Subject, authConfig.CipherKey)
	if err != nil {
		return "", err
	}

	return decryptedSubject, nil
}

func encrypt(decryptedString, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(decryptedString), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encryptedString, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedString)
	if len(ciphertext) < 12 {
		return "", jwt.NewValidationError("invalid phone number", jwt.ValidationErrorClaimsInvalid)
	}

	nonce, ciphertext := ciphertext[:12], ciphertext[12:]
	phoneNumber, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(phoneNumber), nil
}

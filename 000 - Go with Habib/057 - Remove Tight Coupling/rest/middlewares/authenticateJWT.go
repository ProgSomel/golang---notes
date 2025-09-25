package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	accessToken := headerArr[1]
	tokenParts := strings.Split(accessToken, ".")
	
	if len(tokenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	jwtSignature := tokenParts[2]

	fmt.Println(jwtHeader)
	fmt.Println(jwtPayload)
	fmt.Println(jwtSignature)

	message := jwtHeader + "." + jwtPayload

	byteArrSecret := []byte(m.cnf.JWTSecrectKey)

	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	hash := h.Sum(nil)

	newSignature := base64URLEncode(hash)

	if newSignature != jwtSignature {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w, r)
	})
	
}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
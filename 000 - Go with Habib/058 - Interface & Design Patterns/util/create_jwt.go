package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub int `json:"sub"` //? sub is user id
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header {
		Alg: "HS256",
		Typ: "JWT",
	}
	//? converting Header to byte
	byteArrHeader, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	headerB64 := base64URLEncode(byteArrHeader)

	//? ocnverting Data to byte
	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	payloadB64 := base64URLEncode(byteArrData)

	message := headerB64 + "." + payloadB64

	byteArrSecret := []byte(secret)

	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)

	signatureB64 := base64URLEncode(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil

}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
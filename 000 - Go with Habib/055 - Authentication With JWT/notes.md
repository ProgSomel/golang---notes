# 055 - Authentication With JWT(Json Web Token)

## Base64
Base64 is a method for encoding binary data into an ASCII string format.
It converts binary data(like images, files, or bytes) into set of 64 bit printable ASCII characters(A-Z, a-z, 0-9, +, /)

Each group of 6 bits in the binary data is represented by one of these character, allowing binary data to be safely transmitted or stored in text-based systems, like email or json.

```go
package main

import (
	"fmt"
)

func main(){
	// cmd.Serve()
	var s string 
	s = "a"
	
	//? Converting to byte
	byteArray := []byte(s)
	fmt.Println(s)
	fmt.Println(byteArray)
}
```
```bash
a
[97]
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){
	// cmd.Serve()
	var s string 
	s = "aa"
	
	//? Converting to byte
	byteArray := []byte(s)
	fmt.Println(s)
	fmt.Println(byteArray)
}
```
```bash
aa
[97 97]
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	// cmd.Serve()
	var s string 
	s = "a"
	
	//? Converting to byte
	byteArray := []byte(s)
	fmt.Println(s)
	fmt.Println(byteArray)

	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)
	base64Str := enc.EncodeToString(byteArray)
	fmt.Println(base64Str)
}
```
```bash
a
[97]
YQ
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	// cmd.Serve()
	var s string 
	s = "a"
	
	//? Converting to byte
	byteArray := []byte(s)
	fmt.Println(s)
	fmt.Println("after byte converting: ", byteArray)

	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)
	base64Str := enc.EncodeToString(byteArray)
	fmt.Println("after base64 Encoding: ", base64Str)

	byteValue, err := enc.DecodeString(base64Str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("after Decoding base64 String: ", byteValue)

	decodedString := string(byteValue)
	fmt.Println("decoded string: ", decodedString)
}
```
```bash
a
after byte converting:  [97]
after base64 Encoding:  YQ
after Decoding base64 String:  [97]
decoded string:  a
```

**-------------------------------------------------------------------------------------------------------------------------**

## SHA - 1
## SHA - 256
## SHA - 512

**SHA** --> Secure Hash Algortihm
```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main(){
	data := []byte("Hello")
	hash := sha256.Sum256(data)
	fmt.Println("Hash after SHA-256: ", hash)
}
```
```bash
Hash after SHA-256:  [24 95 141 179 34 113 254 37 245 97 166 252 147 139 46 38 67 6 236 48 78 218 81 128 7 209 118 72 38 56 25 105]
```

**-------------------------------------------------------------------------------------------------------------------------**

## HMAC --> Hash-based Message Authentication Code --> takes a text and a secret key --> gives Hash
if secret key changes, then Hash output will be different for same text.

**-------------------------------------------------------------------------------------------------------------------------**

## HMAC-SHA-256 --> will take input a text and a secret key
```go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main(){
	secret := []byte("my-secret")
	message := []byte("Hello World")

	h := hmac.New(sha256.New, secret) //? Here hmac will use sha256 algorithm to hash
	h.Write(message)

	text := h.Sum(nil)

	fmt.Println(text)

}
```
```bash
[130 2 84 128 74 151 27 195 128 229 244 212 65 119 253 247 41 199 3 134 153 126 139 144 27 122 131 5 77 64 67 185]
```

**-------------------------------------------------------------------------------------------------------------------------**

# JWT 
Has 3 parts
- Header
The header typically contains **two** pieces of information:
**Token type**: Usually "JWT"
**Signing algorithm**: Like HMAC SHA256 or RSA
```json
{
"alg": "HS256",
"typ": "JWT"
}
```
- Payload
The payload contains the claims - statements about an entity (typically the user) and additional data. There are three types of claims:
- **Registered claims**: Predefined claims like iss (issuer), exp (expiration), sub (subject)
- **Public claims**: Custom claims that should be collision-resistant
- **Private claims**: Custom claims agreed upon by parties using them
```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022,
  "exp": 1516242622
}
```
- Signature
The signature is created by taking the encoded header, encoded payload, a secret, and the algorithm specified in the header. It's used to verify that the sender is who it claims to be and that the message wasn't changed along the way.

The signature is calculated as:
```bash
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret
)
```

**------------------------------------------------------------------------------------------------------------------------**

## creating JWT
```go
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
	Sub string `json:"sub"` //? sub is user id
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

	byteArrSecret := []byte(secret)

	message := headerB64 + "." + payloadB64

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
```

**------------------------------------------------------------------------------------------------------------------------**


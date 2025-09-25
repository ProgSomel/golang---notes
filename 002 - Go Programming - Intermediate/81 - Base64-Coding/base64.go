package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding");
	fmt.Println("Byte Value: ", data)

	urlUnsafeEncoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("URL Unsafe: ", urlUnsafeEncoded)

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded: ", urlSafeEncoded)

}
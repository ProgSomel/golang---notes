# 81 - Base64 Coding
Encoding is essentially the method used to translate data from its original format to a specific format that can be used by other systems, applications, or protocols.

Encoding means is the process of converting data from one format to another, so it can be stored, transmitted, or processed effectively.

Encoding is essential to store data.

Encoding also help to transmitted data. Data sent over networks or between different systems must
often be encoded to ensure that it is transmitted correctly, and can be understood by the receiving system.

Whenever we are coding, we should also think in machine terms, in machine language terms, in binary terms, or in hexadecimal terms, because our application, our program, our API, once it is complete, the data transmission is going to be in different formats over the network or over the internet over a protocol. 

Encoding is also important in terms of data interoperability.

Encoding helps in ensuring that data can be understood and processed by various systems or softwares, regardless their internal data formats.

Some common Examples of Encoding are:
### Text Encoding
- **ASCII** - ASCII is a character encoding standard that uses seven bits to represent text, mainly for English character
- **UTF8** - It is a variable with character encoding used for electronic communication. It can represent any character in the Unicode standard and is backward compatible with ASCII.
- **UTF16** - It is another set of character encoding used to represent text in Unicode. But this time it is using 16bits for each character.

### Data Encoding
- **Base64** - A method for encoding binary data into a text format. It is commonly used in email and URL Encoding.

**Base64** - is a binary to text Encoding scheme that converts binary data into a textual representation using a set of 64 ASCII Characters.

It is commonly used for transmitting binary data over text based protocols, such as email, or storing binary data as text in database or files, so the Encoding process of Base64 involves converting binary data into a textual format using a fixed set of 64 Characters, and these 64 Characters comprise of uppercase letters, lowercase letters, and digits from 0 to 9, plus sign and forward slash and optionally an equal sign as well. 

Equal sign is used for padding at the end of the encoded data.

Base64 enables binary data to be stored as text in database or files that do not support binary formats.

Base64 is used in URLs, especially URL safe variants and in data URLs for embedding small resources directly into HTML or CSS.

Base64 is also used for data transmission, which allows binary data to be transmitted as text, which is usefull for protocol which only support text.

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding");
	fmt.Println(data)

	//? Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data) //? Base64 string //? Everytime same for same input
	fmt.Println(encoded)
}
```
```bash
[72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
SGVsbG8sIEJhc2U2NCBFbmNvZGluZw==
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding");
	fmt.Println("Byte Value: ", data)

	//? Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data) //? Base64 string //? Everytime same for same input
	fmt.Println("Encoding: ", encoded)

	//? Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in Decoding")
	}

	fmt.Println("Decoding: ", decoded)
}
```
```bash
Byte Value:  [72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
Encoding:  SGVsbG8sIEJhc2U2NCBFbmNvZGluZw==
Decoding:  [72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding");
	fmt.Println("Byte Value: ", data)

	//? Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data) //? Base64 string //? Everytime same for same input
	fmt.Println("Encoding: ", encoded)

	//? Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in Decoding")
	}

	fmt.Println("Decoding: ", decoded)
	fmt.Println(string(decoded))
}
```
```bash
Byte Value:  [72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
Encoding:  SGVsbG8sIEJhc2U2NCBFbmNvZGluZw==
Decoding:  [72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
Hello, Base64 Encoding
```

**-----------------------------------------------------------------------------------------------------------------------------**

## URL Safe Encoding --> Avoid '/' and '+'
```go
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
```
```bash
Byte Value:  [72 101 126 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
URL Unsafe:  SGV+bG8sIEJhc2U2NCBFbmNvZGluZw==
URL Safe Encoded:  SGV-bG8sIEJhc2U2NCBFbmNvZGluZw==
```

So, Base64 encoding is useful in embedding small images or files directly into HTML or CSS using Data.

And similarly, we can also store binary data in text based formats such as JSON or XML, and while
using base64 encoding there are some security considerations that need to be careful about.
- First is base64 encoding is not encryption, we were able to decode that very easily. It is a reversal encoding scheme and we need to ensure proper handling of padding, padding is the equal to symbol when we are decoging base64 data.
- If we need standard encoding and we are okay with the special symbols, then we should use the standard encoding. Otherwise if we want an encoding that is URL Safe, then we use URL encoding and it is based on the context that is standard versus URL Safe.
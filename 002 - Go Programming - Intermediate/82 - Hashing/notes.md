# 82 - Hashing
Hashing is a process used in computing to transform data of any size into a fixed size string of characters.

A "fixed size string of characters" means the output always has the **same length**, no matter what you put into the hash function.

## Examples:
**MD5 hash - always 32 characters:**
- Input: "hello" → Output: 5d41402abc4b2a76b9719d911017c592 (32 chars)
- Input: "this is a very long sentence with lots of words" → Output: 7d865e959b2466918c9863afca942d0f (32 chars)

**SHA-256 hash - always 64 characters:**
- Input: "a" → Output: ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb (64 chars)
- Input: entire book → Output: [still exactly 64 characters]

**SHA-1 hash - always 40 characters:**
- Input: anything → Output: always exactly 40 hexadecimal characters

**Key points:**
1. "Fixed size" = always the same length
2. "String of characters" = typically hexadecimal characters (0-9, a-f)
3. No matter the input size = whether you hash 1 byte or 1 gigabyte, output length is identical

## Why fixed size?
```bash
Variable input → Hash Function → Fixed output
"cat" (3 chars) → MD5 → 32 characters
"The quick brown fox..." (100+ chars) → MD5 → 32 characters  
Entire movie file (2GB) → MD5 → 32 characters
```

## This fixed size makes hashes useful for:
- Comparing files quickly
- Storing passwords (same storage space needed)
- Database indexing
- Digital signatures

**The "characters" are usually hexadecimal digits, but they could be represented in other formats too.**

-------------------------------------------------------------------------------------------------------------------------------

This transformation is done using a special algorithm called a hash algorithm, which is Hash Function.

## Some characteristics of Hashing
- Hashing results in fixed sized output, no matter how large or small the input data is, may be
  our password is 5 characters long or 15 characters long, our hashing will be of fixed size.
- Hash output is deterministic in nature, that means the same input will always produce the same hash output. If you hash the same data multiple times, you will always get the same result.
- Hashing results in unique output, that means even a small change in the input will produce a completely different hash. This property is known as avalanche effect.

## Why do we use Hashing
When you are entering your password into login console, the password is not saved directly into the Database, because if there is a breach then your password is compromised,
So, what happens is that the password that you enter is converted to a hash value, then the hash value is stored in Database.
So, if there is a breach and your password is compromised along with your user data, the entity, the person who conducted the breach now has the user info and the password. So now when he tries to log in, he will not be able to log in because when he enters that string, which has a hash output of your password, it will not be acceptable, because when he enters the password as the hash output as the hash string, it will again be converted to another hash and then be matched with the data that is present in the Database. So that is not goning to match because as we discussed earlier even a single alphabet change is going to result in a completely different hash value.


-------------------------------------------------------------------------------------------------------------------------------

## Now let's dig a little deeper on the practical irreversibility of hash functions.
### Sha 256 / Sha 512
Sha 256 produces a 256 bit hash, resulting in two to the power of 256 possible hash values. That is two multiplied by itself 256 times. And to crack a hash by brute force, an attacker would need to try every possible input until they find one that produces the desired hash, and the time required for this increases exponentially with the length of the hash.

-------------------------------------------------------------------------------------------------------------------------------

## Hashing in Go,
In Go, we have crypto Sha 256 and crypto Sha 512 packages that provide implementations of Sha 256 ans Sha 512 hash functions, respectively.
Now, Sha 512 produces a 512 bit hash values, and this provides a higher level of security and is recommended for applications requiring stronger security guarantees.

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "password123"
	hash := sha256.Sum256([]byte(password))
	fmt.Println(password)
	fmt.Println(hash)
}
```
```bash
password123
[239 146 183 120 186 254 119 30 137 36 91 137 236 188 8 164 74 78 22 108 6 101 153 17 136 31 56 61 68 115 233 79]
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "password123"
	hash := sha256.Sum256([]byte(password))
	fmt.Println(password)
	fmt.Println(hash) //? byte slice of the hash
	//? coverting to hex value
	fmt.Printf("SHA-256 Hash Hex value is: %x\n", hash)
}
```
```bash
password123
[239 146 183 120 186 254 119 30 137 36 91 137 236 188 8 164 74 78 22 108 6 101 153 17 136 31 56 61 68 115 233 79]
SHA-256 Hash Hex value is: ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f
```

**-------------------------------------------------------------------------------------------------------------------------------**

## salting
When we talk about hashing passwords the concept of hashing passwords is incomplete wihtout discussing salting.

Salting adds an extra layer of security by combining the password with a unique random value so it can be random. or you can store a string as a salt and use that with every password that you are hashing.

The practice of salting helps protect against dictionary attacks and rainbow table attacks.

Salt is a value added to the password before hashing, and its purpose is that it ensures that even if two users have the same password, their hashed values will be different due to different salts.
```go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

//? Function to hash password
func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func main() {
	password := "password123"
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt: ", err)
		return
	}

	//? Hash the password with salt
	hash := HashPassword(password, salt)

	//? Store the salt and password to the database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Hash: ", hash)
}
```
```bash
Salt:  1dZ1/MYmY7bOyWyBrsM3wg==
Hash:  z+Pkh1apzfZzjNpNUKcx9jCU/XakIubDRTUA9f96bNM=
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

//? Function to hash password
func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func main() {
	password := "password123"
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt: ", err)
		return
	}

	//? Hash the password with salt
	signUpHash := HashPassword(password, salt)

	//? Store the salt and password to the database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Signup Hash: ", signUpHash)

	//! verify
	//? retrieve the saltStr and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt: ", err)
		return
	}
	loginHash := HashPassword(password, decodedSalt)

	//? compare the stored signUp hash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct, You are logged in.")
	}else{
		fmt.Println("Login Failed, Please check user credentials.")
	}
}
```
```bash
Salt:  taS8+xNnQ789vUjPrhl6kA==
Signup Hash:  Z0CW3diD0N+QqGa9czNK++hjCHZbOxF6v2WX4hVs+Jk=
Password is correct, You are logged in.
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

//? Function to hash password
func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func main() {
	password := "password123"
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt: ", err)
		return
	}

	//? Hash the password with salt
	signUpHash := HashPassword(password, salt)

	//? Store the salt and password to the database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Signup Hash: ", signUpHash)

	//! verify
	//? retrieve the saltStr and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt: ", err)
		return
	}
	loginHash := HashPassword("password234", decodedSalt)

	//? compare the stored signUp hash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct, You are logged in.")
	}else{
		fmt.Println("Login Failed, Please check user credentials.")
	}
}
```
```bash
Salt:  uQGmuXZd9WvRQomIEDnnbw==
Signup Hash:  02fvhVtcRGyGioDBt0cgGYmHKuqZg75rDLtdHv8YO4E=
Login Failed, Please check user credentials.
```


**-------------------------------------------------------------------------------------------------------------------------------**

crypto/rand package generates cryptographically secure random numbers, which the math Rand package is Unable to do.
Now, the cryptographically secure random numbers are different from the general random numbers because, these numbers are generated in a way that makes them unpredictable and resistant to reverse engineering.

General random numbers, however, are predictable if enough information about the internal state is known.

Another important factor is that cryptographically secure random numbers are generated using cryptographic algorithms that are designed to be secure against attacks.

**-------------------------------------------------------------------------------------------------------------------------------**

So, overall, Sha 256 and Sha 512 are widely used cryptographic hash functions that provide data integrity and security in various applications.

Go's crypto Sha 256 and crypto Sha 512 packages offer efficient and secure implementations for computing these hashes.
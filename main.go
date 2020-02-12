package main

import (
	"crypto/aes"
	"crypto/sha256"
	"log"
	"os"
)

func main() {
	operation, filename := ParseCommandLineArguments()
	password := ReadPassword()
	key := sha256.Sum256(password)
	cipher, err := aes.NewCipher(key[:])
	if err != nil {
		log.Fatalln("Unable to create AES cipher", err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Unable to open file", filename, err)
	}
	var crypt func([]byte, []byte)
	if operation == Encrypt {
		crypt = cipher.Encrypt
	} else {
		crypt = cipher.Decrypt
	}
	source := make([]byte, cipher.BlockSize())
	for count, err := file.Read(source); count > 0; count, err = file.Read(source) {
		if err != nil {
			file.Close()
			log.Fatalln("Unable to read file", filename, err)
		}
		destination := make([]byte, cipher.BlockSize())
		crypt(destination, source)
		os.Stdout.Write(destination)
	}
	file.Close()
}

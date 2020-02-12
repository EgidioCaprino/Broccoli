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
	source := make([]byte, cipher.BlockSize())
	for count, err := file.Read(source); count > 0; count, err = file.Read(source) {
		if err != nil {
			file.Close()
			log.Fatalln("Unable to read file", filename, err)
		}
		destination := make([]byte, cipher.BlockSize())
		if operation == Encrypt {
			cipher.Encrypt(destination, source)
		} else {
			cipher.Decrypt(destination, source)
		}
		os.Stdout.Write(destination)
	}
	file.Close()
}

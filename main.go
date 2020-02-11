package main

import (
	"bufio"
	"crypto/aes"
	"crypto/sha256"
	"log"
	"os"
)

const usage = "Usage: broccoli file"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalln(usage)
	}
	operation := args[0]
	filename := args[1]

	if operation != "encrypt" && operation != "decrypt" {
		log.Fatalln("Operation should be encrypt or decrypt")
	}

	encrypt := operation == "encrypt"

	reader := bufio.NewReader(os.Stdin)

	log.Println("Password")
	password, err := readPassword(reader, make([]byte, 0))

	if err != nil {
		log.Fatalln("Unable to read password: ", err)
	}

	key := sha256.Sum256(password)
	cipher, err := aes.NewCipher(key[:])

	if err != nil {
		log.Fatalln("Unable to create AES cipher: ", err)
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
		if encrypt {
			cipher.Encrypt(destination, source)
		} else {
			cipher.Decrypt(destination, source)
		}
		os.Stdout.Write(destination)
	}

	file.Close()
}

func encrypt(filename string, password []byte) {

}

func decrypt(filename string, password []byte) {
	key := sha256.Sum256(password)
	cipher, err := aes.NewCipher(key[:])

	if err != nil {
		log.Fatalln("Unable to create AES cipher: ", err)
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
		cipher.Decrypt(destination, source)
		os.Stdout.Write(destination)
	}

	file.Close()
}

func readPassword(reader *bufio.Reader, accumulator []byte) ([]byte, error) {
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	password := append(accumulator, line...)
	if isPrefix {
		return readPassword(reader, password)
	}
	return password, nil
}

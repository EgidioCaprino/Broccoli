package main

import "log"

const encrypt = "encrypt"
const decrypt = "decrypt"

type Operation string

const Encrypt Operation = encrypt
const Decrypt Operation = decrypt

func ToOperation(value string) Operation {
	if !IsValidOperation(value) {
		log.Fatalln("Invalid operation", value)
	}
	if value == encrypt {
		return Encrypt
	}
	return Decrypt
}

func IsValidOperation(value string) bool {
	return value == encrypt || value == decrypt
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseCommandLineArguments() (Operation, string) {
	operation := flag.String("operation", "", fmt.Sprintf("%v or %v", Encrypt, Decrypt))
	filename := flag.String("file", "", "file to encrypt or decrypt")
	flag.Parse()
	return ToOperation(*operation), sanitizeFilename(*filename)
}

func ReadPassword() []byte {
	reader := bufio.NewReader(os.Stdin)
	accumulator := make([]byte, 0)
	log.Print("Password: ")
	return readPassword(reader, accumulator)
}

func sanitizeFilename(filename string) string {
	sanitizedFilename := strings.TrimSpace(filename)
	if len(sanitizedFilename) == 0 {
		log.Fatalln("Invalid filename", filename)
	}
	return sanitizedFilename
}

func readPassword(reader *bufio.Reader, accumulator []byte) []byte {
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		log.Fatalln("Unable to read password", err)
	}
	password := append(accumulator, line...)
	if isPrefix {
		return readPassword(reader, password)
	}
	return password
}

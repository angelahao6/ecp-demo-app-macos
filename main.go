package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"github.com/googleapis/enterprise-certificate-proxy/client"
)

func main() {
	key, err := client.Cred("")
	if err != nil {
		log.Printf("Cred: got %v, want nil err", err)
	}
	mode := flag.String("mode", "", "the flag for action: -e for encryption or -d for decryption")
	msg := flag.String("msg", "", "the message to encrypt or decrypt")
	flag.Parse()

	if *mode == "" || *msg == "" {
		log.Printf("Usage: go run main.go -mode <-e|-d> -msg <msg>")
		return
	}

	if *mode != "-e" && *mode != "-d" {
		log.Printf("Invalid mode: %v", *mode)
		return
	}

	if *mode == "-e" {
		byteSlice := []byte(*msg)
		ciphertext, err := key.Encrypt(byteSlice)
		if err != nil {
			fmt.Printf("Encrypt: got %v, want nil err \n", err)
			return
		}
		str := base64.StdEncoding.EncodeToString(ciphertext)
		fmt.Printf("Encrypted message: %v \n", str)
	} else {
		str, err := base64.StdEncoding.DecodeString(*msg)
		if err != nil {
			fmt.Printf("Base64: got %v, want nil err \n", err)
		}
		plaintext, err := key.Decrypt([]byte(str))
		if err != nil {
			fmt.Printf("Decrypt: got %v, want nil err \n", err)
			return
		}
		plainString := string(plaintext)
		fmt.Printf("Decrypted message: %v \n", plainString)
	}
}

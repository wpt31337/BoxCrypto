package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"wptconsulting.com/utilities"
)

// AESKeySize is the size of the AES key in bytes (256 bits)
const AESKeySize = 32

func main() {

	newKey := flag.Bool("n", false, "Generate new random key")
	encryptFile := flag.String("e", "", "File to encrypt")
	passphrase := flag.String("p", "", "AES 32-byte Passpharse in Base64 encoding")
	decryptFile := flag.String("d", "", "File to decrypt")

	flag.Parse()

	if *newKey == true {

		b, err := utilities.NewAESKey()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		} else {
			fmt.Println(utilities.EncodeToBase64(b))
		}

		return

	} else if *encryptFile != "" && *passphrase == "" {

		log.Fatal("-e file to encrypt, -d file to decrypt, -p passphrase, -n for new key")

	} else if *decryptFile != "" && *passphrase == "" {

		log.Fatal("-e file to encrypt, -d file to decrypt, -p passphrase, -n for new key")

	} else if *encryptFile == "" && *decryptFile == "" {

		log.Fatal("-e file to encrypt, -d file to decrypt, -p passphrase, -n for new key")
	}

	if *encryptFile != "" {

		outfile := *encryptFile + ".enc"
		err := utilities.EncryptToFile(*passphrase, *encryptFile, outfile)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		fmt.Printf("Successfully encrypted to %s\n", outfile)
		return
	}

	if *decryptFile != "" {

		outfile := getDecryptFilename(*decryptFile)

		err := utilities.DecryptToFile(*passphrase, *decryptFile, outfile)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		fmt.Printf("Successfully decrypted to %s\n", outfile)
		return

	}

}

func getDecryptFilename(input string) string {

	lastDotIndex := strings.LastIndex(input, ".")
	if lastDotIndex == -1 {
		return input
	}
	result := input[:lastDotIndex]
	if endsWithKey(result) {
		return result
	}
	return result + ".key"
}

func endsWithKey(input string) bool {
	if len(input) >= 4 {
		return input[len(input)-4:] == ".key"
	}
	return false
}

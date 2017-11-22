// vigenere-encrypt <encipherment key> <plaintext filename>
// vigenere-encrypt DEADBEEF plaintext.txt > ciphertext.txt

package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func showUsage() {
	fmt.Print("\n* Vigenere Decryption Tool *\n\n")
	fmt.Print("A tool for decrypting Vigenere ciphertext files\n")
	fmt.Print("usage: vigenere-decrypt <key> <filename>\n\n")
}

func check(e error){
	if e != nil {
		showUsage()
		panic(e)
	}
}

func sanitize(ip string) string {
	op := []rune{}
	for _, v := range ip {
		if 65 <= v && v <= 90 {
			op = append(op, v)
		} else if 97 <= v && v <= 122 {
			op = append(op, v-32)
		}
	}
	return string(op)
}

func decode(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func decrypt(msg string, key string) string{
	cipherText := make([]rune, 0, len(msg))
	for i, v := range msg {
		cipherText = append(cipherText, decode(v, rune(key[i%len(key)])))
	}
	return string(cipherText)
}

func main(){
	if len(os.Args) < 3 {
		showUsage()
		os.Exit(1)
	}
	key := os.Args[1]
	plainTextFile,err := ioutil.ReadFile(os.Args[2])
	check(err)
	plainText := string(plainTextFile)
	smsg, skey := sanitize(plainText), sanitize(key)
	//fmt.Print("\nKEY: " + skey)
	fmt.Print("\nCiphertext:\t" + smsg)
	fmt.Print("\nPlaintext:\t" + decrypt(smsg, skey) + "\n\n")
}

// vigenere-encrypt <encipherment key> <plaintext filename>
// vigenere-encrypt DEADBEEF plaintext.txt > ciphertext.txt

package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func showUsage() {
	fmt.Print("\n* Vigenere Encryption Tool *\n\n")
	fmt.Print("A tool for encrypting plaintext files using the vigenere cipher\n")
	fmt.Print("usage: vigenere-encrypt <key> <filename>\n\n")
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

func encode(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func encrypt(msg string, key string) string{
	cipherText := make([]rune, 0, len(msg))
	for i, v := range msg {
		cipherText = append(cipherText, encode(v, rune(key[i%len(key)])))
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
	//fmt.Print("\nPlaintext:\t" + smsg)
	fmt.Print(encrypt(smsg, skey))
}

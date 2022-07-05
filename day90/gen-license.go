package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

//nopadding模式
func noPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if len(ciphertext)%blockSize == 0 {
		return ciphertext
	} else {
		return append(ciphertext, bytes.Repeat([]byte{byte(0)}, padding)...)
	}

	//padtext := bytes.Repeat([]byte{byte(0)}, padding)
	//return append(ciphertext, padtext...)
}

//AesEncrypt 加密函数
func aesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = noPadding(plaintext, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("缺少参数,请输入主板uuid(可多个，空格分隔)")
	}

	args = args[1:]
	var args_parse []string
	for _, arg := range args {
		args_parse = append(args_parse, strings.Join([]string{"\"", arg, "\""}, ""))
	}
	arg_str := strings.Join(args_parse, ",")

	//var args_parse string
	//args_parse = strings.Join(args,",")
	plaintextPref := "{\n\t\"expire\":\"20230628\",\n\t\"machineCodes\":["
	plaintextSuffix := "],\n\t\"products\":[\n\t\t{\n\t\t\t\"amount\":1000,\n\t\t\t\"id\":\"1001\",\n\t\t\t\"name\":\"接入量\"\n\t\t},\n\t\t{\n\t\t\t\"amount\":1000,\n\t\t\t\"id\":\"1002\",\n\t\t\t\"name\":\"视频并发量\"\n\t\t},\n\t\t{\n\t\t\t\"amount\":1000,\n\t\t\t\"id\":\"1003\",\n\t\t\t\"name\":\"视频转码并发量\"\n\t\t}\n\t],\n\t\"project\":\"video-media-platform\",\n\t\"version\":\"1.0.0\"\n}"
	plainText := strings.Join([]string{plaintextPref, arg_str, plaintextSuffix}, "")

	fmt.Println(plainText)
	key := []byte("eXVubGl6aGlodWk=")
	//加密
	ciphertext, err := aesEncrypt([]byte(plainText), key, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文：", base64.StdEncoding.EncodeToString(ciphertext))

}

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

//已对齐，填充一个长度为blockSize且每个字节均为blockSize的数据。
//未对齐，需要补充的字节个数为n，则填充一个长度为n且每个字节均为n的数据。
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func getMachineCode() string {
	cmd := exec.Command("bash", "-c", "cat /sys/class/dmi/id/product_uuid")
	cmdOutput, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	content := string(cmdOutput)
	splitReg, err := regexp.Compile("\\s+")
	content = splitReg.ReplaceAllString(content, "")
	fmt.Println("主板uuid:", content)
	return content
}

func encode(key []byte, iv []byte) {
	plaintext := []byte(getMachineCode())
	plaintext = []byte(strings.Join([]string{"video-media-platform", string(plaintext), "20270101"}, "/"))

	//plaintext := []byte("video-media-platform/00000000-0000-0000-0000-AC1F6B947B76/2022-06-29 14:43:19")
	//加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印加密base64后密码
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	//fmt.Println(base64Ciphertext)
	fmt.Println("密文：", base64.StdEncoding.EncodeToString([]byte(base64Ciphertext)))
	//fmt.Println("ajlKWkNnc05zZXIwbTljdnBuNnVEZVlGd3VvNUY1alJsYitrVytheU9hNVJPLzIyenV1ZWRHU1lTLzhYckR4cXI2T21SVVdXTTZ6L3V5blI5dEtsbUhjczIwenFZdXhlMjdZcHVKaWNucm89")

	//解密
	plaintext, err = AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印解密明文
	fmt.Println(string(plaintext))
}

//解密密文
func decode(ciphertext string, key []byte, iv []byte) string {
	//两次base64
	oneTimeDecode, error := base64.StdEncoding.DecodeString(ciphertext)
	if error != nil {
		panic(error)
	}
	towTimeDecode, error := base64.StdEncoding.DecodeString(string(oneTimeDecode))
	if error != nil {
		panic(error)
	}

	plaintext, _ := AesDecrypt([]byte(towTimeDecode), key, iv)

	return string(plaintext)

}

func main3() {
	key := []byte("sN1DEJAVZNf3OdM3")
	iv := []byte("GDHgt7hbKpsIR4b4")

	wordPtr := flag.String("opts", "encode", "请输入操作 encode, decode")
	flag.Parse()
	opts := *wordPtr
	switch opts {
	case "encode":
		encode(key, iv)
	case "decode":
		arg := flag.Args()
		if len(arg) < 1 {
			panic("请输入要解密的密文")
		}
		fmt.Println(decode(arg[0], key, iv))
	default:
		panic("无效参数")
	}

}

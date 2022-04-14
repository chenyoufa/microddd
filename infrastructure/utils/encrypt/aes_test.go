package encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	origData := "Hello World" // []byte() // 待加密的数据
	key := ""                 // []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(string(encrypted), key)
	log.Println("解密结果：", string(decrypted))
	assert.Equal(t, string(origData), string(decrypted))

	log.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(string(encrypted), key)
	log.Println("解密结果：", string(decrypted))
	assert.Equal(t, string(origData), string(decrypted))

	log.Println("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(string(encrypted), key)
	log.Println("解密结果：", string(decrypted))
	assert.Equal(t, string(origData), string(decrypted))
}

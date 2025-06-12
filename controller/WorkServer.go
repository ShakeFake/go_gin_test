package controller

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var (
	defaultSecretKey = "vfw212u9y8d2fwfl"
)

type DecryptReq struct {
	Key string `json:"key" validator:"required"`
}

func DecryptSecretKey(c *gin.Context) {
	var d DecryptReq
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 解谜
	decryptKey := stringDecrypt(d.Key, defaultSecretKey)

	// 直接返回即可
	c.String(http.StatusOK, decryptKey)
	return
}

func stringDecrypt(data, key string) string {
	cryptedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Errorf("DecodeString failed:%v ,data is:%v,key is:%v\n", err, data, key)
		return ""
	}

	originData, err := AesDecrypt(cryptedData, []byte(key))
	if err != nil {
		log.Errorf("AesDecrypt failed:%v ,data is:%v,key is:%v\n", err, data, key)
		return ""
	}
	return string(originData)
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	if len(crypted)%blockSize != 0 {
		return nil, errors.New("input data format error")
	}
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type SignatureReq struct {
	Data string `json:"data" validator:"required"`
}

var (
	mu       sync.Mutex
	Resource map[string]int
)

func init() {
	Resource = make(map[string]int)
}

func GetSignature(c *gin.Context) {
	var s SignatureReq
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hardInfo := strings.Split(s.Data, "/")[0]
	log.Infof(hardInfo)

	if number, ok := Resource[hardInfo]; ok && number > 0 {
		c.String(http.StatusTooManyRequests, "")
		return
	} else {
		mu.Lock()
		Resource[hardInfo]++
		mu.Unlock()
	}

	privateKeyPath := "./resource/rsa-prv.pem"
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Errorf("Failed to read private key file: %v", err)
	}

	// Decode PEM to get the private key
	block, _ := pem.Decode(privateKeyData)
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	// Sign the message
	hash := crypto.SHA256.New()
	hash.Write([]byte(s.Data))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, parsedKey.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		log.Errorf("Failed to sign message: %v", err)
	}

	//// Base64 encode the signature
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	log.Infof("current signature is: %v", signatureBase64)
	// Return the signature
	c.String(http.StatusOK, signatureBase64)
}

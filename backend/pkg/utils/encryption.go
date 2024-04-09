/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:23:42
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:39:08
 * @FilePath: /goMall/backend/pkg/utils/encryption.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

// AES加密算法
type Encryption struct {
	key string
}

var Encrypt *Encryption

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// 加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	// 密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// 去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	// 去掉的长度
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst)
	if err != nil {
		return "0"
	}
	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}

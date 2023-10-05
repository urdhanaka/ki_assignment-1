package utils

import "crypto/des"

func EncryptDES(key []byte, data []byte) ([]byte, error) {
	c, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	c.Encrypt(dst, data)
	return dst, nil
}

func DecryptDES(dst []byte, key []byte, data []byte) error {
	c, err := des.NewCipher(key)
	if err != nil {
		return err
	}
	c.Decrypt(dst, data)
	return nil
}

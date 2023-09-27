package utils

import (
	"crypto/rc4"
)

// RC4 is cryptographically broken and should not be used for secure
// applications.

func EncryptRC4(key []byte, data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)
	return dst, nil
}

func DecryptRC4(dst []byte, key []byte, data []byte) error {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}
	c.XORKeyStream(dst, data)
	return nil
}

// DecryptRC4InPlace decrypts data in place. It is equivalent to DecryptRC4
// but reuses the input slice for the result, which can be useful for reducing
// memory allocations if you don't need to keep the original data.
func DecryptRC4InPlace(key []byte, data []byte) error {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}
	c.XORKeyStream(data, data)
	return nil
}
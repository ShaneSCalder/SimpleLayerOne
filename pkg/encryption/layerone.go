package encryption

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20"
)

// generateKey creates a new random 256-bit key for ChaCha20 encryption.
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32) // 256 bits for ChaCha20
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// generateNonce creates a new random 96-bit nonce for ChaCha20 encryption.
func GenerateNonce() ([]byte, error) {
	nonce := make([]byte, 12) // 96 bits for XChaCha20
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

// encrypt encrypts the given data using ChaCha20 with the provided key and nonce.
func Encrypt(data, key, nonce []byte) ([]byte, error) {
	if len(key) != chacha20.KeySize {
		return nil, errors.New("invalid key size")
	}
	if len(nonce) != chacha20.NonceSize {
		return nil, errors.New("invalid nonce size")
	}

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	return encrypted, nil
}

// decrypt decrypts the given data using ChaCha20 with the provided key and nonce.
func Decrypt(data, key, nonce []byte) ([]byte, error) {
	// Decryption is symmetric in ChaCha20
	return Encrypt(data, key, nonce)
}
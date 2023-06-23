/*
 * File: security.go
 * File Created: Monday, 12th June 2023 2:03:52 pm
 * Last Modified: Friday, 23rd June 2023 12:28:20 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/sha3"
)

// Encrypt encrypts the data using AES encryption with the specified key.
func Encrypt(data []byte, key []byte) ([]byte, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new Galois Counter Mode (GCM) cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate a unique nonce for this encryption operation
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Seal the plaintext using GCM encryption
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// DeriveKey derives a valid AES encryption key from the hardware ID using SHA-256 hashing.
func DeriveKey(hardwareID string) []byte {
	// Create a new SHA-256 hash object
	hasher := sha3.New256()

	// Hash the hardware ID
	hasher.Write([]byte(hardwareID))

	// Return the resulting hash as the derived key
	return hasher.Sum(nil)
}

package permutation

import "fmt"

type PermutationCipher struct {
	key []int
}

func NewPermutationCipher(key []int) *PermutationCipher {
	return &PermutationCipher{key: key}
}

func (c PermutationCipher) encryptBlock(block []byte) []byte {
	result := make([]byte, len(block))
	for i := range block {
		result[i] = block[c.key[i]]
	}
	return result
}

func (c PermutationCipher) Encrypt(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return message, nil
	}
	if len(message)%len(c.key) != 0 {
		return nil, fmt.Errorf("message len mod key len != 0")
	}
	result := make([]byte, 0, len(message))
	i := 0
	for ; (i+1)*len(c.key) < len(message); i++ {
		result = append(result, c.encryptBlock(message[i*len(c.key):(i+1)*len(c.key)])...)
	}
	result = append(result, c.encryptBlock(message[i*len(c.key):])...)
	return result, nil
}

func (c PermutationCipher) decryptBlock(block []byte) []byte {
	result := make([]byte, len(block))
	for i := range block {
		result[c.key[i]] = block[i]
	}
	return result
}

func (c PermutationCipher) Decrypt(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return message, nil
	}
	if len(message)%len(c.key) != 0 {
		return nil, fmt.Errorf("message len mod key len != 0")
	}
	result := make([]byte, 0, len(message))
	i := 0
	for ; (i+1)*len(c.key) < len(message); i++ {
		result = append(result, c.decryptBlock(message[i*len(c.key):(i+1)*len(c.key)])...)
	}
	result = append(result, c.decryptBlock(message[i*len(c.key):])...)
	return result, nil
}

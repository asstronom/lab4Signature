package permutation

import (
	"crypto/rand"
	"math/big"
)

type PermutationCipher struct {
	key []int
}



func GenKey(size int) []int {
	key := make([]int, size)
	//generates key by creating array of given size with this elements [0, 1, 2..., size-1]
	for i := range key {
		key[i] = i
	}
	//shuffles key
	for i := size - 1; i > 0; i-- { // Fisher–Yates shuffle
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		key[i], key[j.Int64()] = key[j.Int64()], key[i]
	}
	return key
}

func NewPermutationCipher(key []int) *PermutationCipher {
	return &PermutationCipher{key: key}
}

//encrypts one block
func (c PermutationCipher) encryptBlock(block []byte) []byte {
	result := make([]byte, len(block))
	j := -1
	for i := 0; i < len(block); i++ {
		j++
		if c.key[j] > len(block)-1 {
			i--
			continue
		} else {
			result[i] = block[c.key[j]]
		}
	}
	return result
}

//encrypts full message
func (c PermutationCipher) Encrypt(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return message, nil
	}
	// if len(message)%len(c.key) != 0 {
	// 	return nil, fmt.Errorf("message len mod key len != 0")
	// }
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
	j := -1
	for i := 0; i < len(block); i++ {
		j++
		if c.key[j] > len(block)-1 {
			i--
			continue
		} else {
			result[c.key[j]] = block[i]
		}
	}
	return result
}

func (c PermutationCipher) Decrypt(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return message, nil
	}
	// if len(message)%len(c.key) != 0 {
	// 	return nil, fmt.Errorf("message len mod key len != 0")
	// }
	result := make([]byte, 0, len(message))
	i := 0
	for ; (i+1)*len(c.key) < len(message); i++ {
		result = append(result, c.decryptBlock(message[i*len(c.key):(i+1)*len(c.key)])...)
	}
	result = append(result, c.decryptBlock(message[i*len(c.key):])...)
	return result, nil
}

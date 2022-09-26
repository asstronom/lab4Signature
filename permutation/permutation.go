package permutation

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

func (c PermutationCipher) Encrypt(message []byte) []byte {
	if len(message) == 0 {
		return message
	}
	result := make([]byte, 0, len(message))
	i := 0
	for ; (i+1)*len(c.key) < len(message); i++ {
		result = append(result, c.encryptBlock(message[i*len(c.key):(i+1)*len(c.key)])...)
	}
	result = append(result, c.encryptBlock(message[i*len(c.key):])...)
	return result
}

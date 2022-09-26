package permutation

type PermutationCipher struct {
	key []int
}

func (c PermutationCipher) encryptBlock(block []byte) []byte {
	result := make([]byte, len(block))
	for i := range block {
		result[i] = block[c.key[i]]
	}
	return result
}
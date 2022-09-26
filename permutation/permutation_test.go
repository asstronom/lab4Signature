package permutation

import "testing"

func TestEncryptBlock(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	block := []byte{1, 2, 3, 4, 5}
	correct := []byte{5, 1, 3, 2, 4}
	cipher := NewPermutationCipher(key)
	block = cipher.encryptBlock(block)
	for i, v := range block {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v", block, correct)
		}
	}
}

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

func TestEncrypt(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	message := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	correct := []byte{5, 1, 3, 1, 4, 10, 6, 8, 7, 9, 15, 11, 13, 12, 14, 16, 17}
	cipher := NewPermutationCipher(key)
	result := cipher.Encrypt(message)
	for i, v := range result {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v", result, correct)
		}
	}
}

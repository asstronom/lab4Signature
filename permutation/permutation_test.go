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

func TestEncryptBlockWithSmallBlock(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	block := []byte{1, 2, 3}
	correct := []byte{1, 3, 2}
	cipher := NewPermutationCipher(key)
	block = cipher.encryptBlock(block)
	for i, v := range block {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v", block, correct)
			break
		}
	}
}

func TestEncrypt(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	message := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	correct := []byte{5, 1, 3, 2, 4, 10, 6, 8, 7, 9, 15, 11, 13, 12, 14, 20, 16, 18, 17, 19}
	cipher := NewPermutationCipher(key)
	result, err := cipher.Encrypt(message)
	if err != nil {
		t.Errorf("error encrypting %s\n", err)
	}
	for i, v := range result {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v\n", result, correct)
		}
	}

	message = []byte{}
	result, err = cipher.Encrypt(message)
	if err != nil {
		t.Errorf("error encrypting %s\n", err)
	}
	if len(result) != 0 {
		t.Errorf("error encrypting, encryption returned not empty slice: %v", result)
	}

	message = []byte{1}
	result, err = cipher.Encrypt(message)
	if err == nil {
		t.Errorf("error encryption, encryption didn't return error after recieving message with wrong length, %d", len(result))
	}
}

func TestDecryptBlock(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	block := []byte{5, 1, 3, 2, 4}
	correct := []byte{1, 2, 3, 4, 5}
	cipher := NewPermutationCipher(key)
	block = cipher.decryptBlock(block)
	for i, v := range block {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v", block, correct)
		}
	}
}

func TestDecryptBlockWithSmallBlock(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	block := []byte{1, 3, 2}
	correct := []byte{1, 2, 3}
	cipher := NewPermutationCipher(key)
	block = cipher.encryptBlock(block)
	for i, v := range block {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v", block, correct)
			break
		}
	}
}

func TestDecrypt(t *testing.T) {
	key := []int{4, 0, 2, 1, 3}
	message := []byte{5, 1, 3, 2, 4, 10, 6, 8, 7, 9, 15, 11, 13, 12, 14, 20, 16, 18, 17, 19}
	correct := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	cipher := NewPermutationCipher(key)
	result, err := cipher.Decrypt(message)
	if err != nil {
		t.Errorf("error encrypting %s\n", err)
	}
	for i, v := range result {
		if v != correct[i] {
			t.Errorf("wrong encryption, %v != %v\n", result, correct)
		}
	}

	message = []byte{}
	result, err = cipher.Decrypt(message)
	if err != nil {
		t.Errorf("error encrypting %s\n", err)
	}
	if len(result) != 0 {
		t.Errorf("error encrypting, encryption returned not empty slice: %v", result)
	}

	message = []byte{1}
	result, err = cipher.Decrypt(message)
	if err == nil {
		t.Errorf("error encryption, encryption didn't return error after recieving message with wrong length, %d", len(result))
	}
}

func TestGenKey(t *testing.T) {
	key := GenKey(10 + (23+92)%7)
	t.Logf("%v\n", key)
	key = GenKey(10 + (23+92)%7)
	t.Logf("%v\n", key)
}

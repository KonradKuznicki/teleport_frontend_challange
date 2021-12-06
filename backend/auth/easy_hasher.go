package auth

type EasyHash struct {
	salt string
}

func (h *EasyHash) Verify(pass string, hash string) (bool, error) {
	newHash, err := h.Hash(pass)
	if err != nil {
		return false, err
	}
	return newHash == hash, nil
}

func (h *EasyHash) Hash(pass string) (string, error) {
	return pass + h.salt, nil
}

func NewEasyHash(salt string) *EasyHash {
	return &EasyHash{salt: salt}
}

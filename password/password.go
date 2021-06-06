package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func EncryptPassword(password string, p *Params) (encryptedPassword string, err error) {
	salt, err := GenerateSalt(p.SaltLength)

	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encryptedPassword = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encryptedPassword, nil
}

func GenerateSalt(length uint32) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

package crypto

import (
	"fmt"

	"github.com/matthewhartstonge/argon2"
	_ "golang.org/x/crypto/argon2"
)

type Argon2Adapter struct{}

func (a *Argon2Adapter) Hash(password string) (string, error) {

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(string(encoded)), nil
}

func (a *Argon2Adapter) Verify(hashedPassword, password string) (bool, error) {

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
	if err != nil {
		return false, nil
	}

	if ok {
		return true, nil
	}
	return false, nil
}

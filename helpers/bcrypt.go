package helpers

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash handles generating password hash
// bcrypt hashes password of type byte
func GeneratePasswordHash(password []byte) string {
	// default cost is 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	// If there was an error panic
	if err != nil {
		panic(err)
	}

	// return stringified password
	return string(hashedPassword)
}

// PasswordCompare handles password hash compare
func PasswordCompare(password []byte, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
}

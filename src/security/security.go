package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compara senha e has e retorna se são iguais
func VerifyPassword(hashPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}

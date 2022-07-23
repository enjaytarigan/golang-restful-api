package security

type PasswordHash interface {
	Hash(password string) (string, error)
	IsMatch(password, hashedPassword string) bool
}

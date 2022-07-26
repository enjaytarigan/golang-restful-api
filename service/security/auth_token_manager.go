package security

type AuthenticationTokenManager interface {
	CreateAccessToken(userId int) (accessToken string, err error)
	VerifyAccessToken(token string) (userId int, err error)
}

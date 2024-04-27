package refreshTokenRepository

type RefreshTokenRepository interface {
	Set(email string) bool
	Find(email, token string) bool
}

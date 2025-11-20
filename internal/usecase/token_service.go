package usecase

type TokenService interface {
	GenerateToken(userId uint) (string, error)
}

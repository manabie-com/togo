package auth

// AuthUsecase is the definition for collection of methods related to the `auths` table use case
type AuthUseCase interface {
	ValidateUser(username string) (bool, error)
	GenerateToken(userID, maxTaskPerday string) (*string, error)
}

type authUseCaseRepository struct {
	repository AuthRepository
}

// NewAuthUsecase returns a AuthUsecase attached with methods related to the `auths` table use case
func NewAuthUseCase(repository AuthRepository) AuthUseCase {
	return &authUseCaseRepository{repository: repository}
}

// ValidateUser returns a AuthUsecase attached with methods related to the `auths` table use case
func (ar *authUseCaseRepository) ValidateUser(username string) (bool, error) {
	return ar.repository.ValidateUser(username)
}

// GenerateToken returns a AuthUsecase attached with methods related to the `auths` table use case
func (ar *authUseCaseRepository) GenerateToken(userID, maxTaskPerday string) (*string, error) {
	return ar.repository.GenerateToken(userID, maxTaskPerday)
}

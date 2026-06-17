package excludepkg

import "errors"

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type User struct {
	Username string
	Password string
}

type UserRepository interface {
	UserOfUsername(username string) (*User, error)
}

type AuthService struct {
	userRepo UserRepository
	logger   Logger
}

func NewAuthService(userRepo UserRepository, logger Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (as *AuthService) Authenticate(username, password string) error {
	as.logger.Info("Authenticate %q", username)

	usr, err := as.userRepo.UserOfUsername(username)
	if err != nil {
		return err
	}

	if usr.Password != password {
		return errors.New("invalid username or password")
	}

	return nil
}

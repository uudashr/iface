package basic

import (
	"errors"
	"time"
)

type Logger interface {
	Debug(msg string, args ...any) // want "^method 'Debug\\(\\)' is declared on interface 'Logger' but not used within the package$"
	Info(msg string, args ...any)

	Warn(msg string, args ...any) //iface:ignore=unusedmethod

	//iface:ignore=unusedmethod
	Error(msg string, args ...any)
}

//iface:ignore=unusedmethod
type Meter interface {
	CountInc(name string, val int, attrs map[string]any)
	Record(name string, val int, attrs map[string]any)
	Set(name string, val int, attrs map[string]any)
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
	meter    Meter
}

func NewAuthService(userRepo UserRepository, logger Logger, meter Meter) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		logger:   logger,
		meter:    meter,
	}
}

func (as *AuthService) Authenticate(username, password string) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		as.meter.Record("authenticate.duration", int(duration.Milliseconds()), nil)
	}()

	as.logger.Info("Authenticate %q", username)

	usr, err := as.userRepo.UserOfUsername(username)
	if err != nil {
		return err
	}

	if usr.Password != password {
		as.meter.CountInc("authenticate.failed", 1, map[string]any{"username": username})
		return errors.New("invalid username or password")
	}

	as.meter.CountInc("authenticate.succeed", 1, map[string]any{"username": username})

	return nil
}

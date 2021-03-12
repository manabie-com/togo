package services_test

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type MockedUserRepo struct {
	Users []entities.User
}

func (ur *MockedUserRepo) GetUserByUserID(ctx context.Context, userID string) (*entities.User, error) {
	for _, user := range ur.Users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, nil
}

func (m *MockedUserRepo) GetUserTaskLimit(ctx context.Context, userID string) (int, error) {
	return 5, nil
}

func newDefaultAuthSvc(initalUsers []entities.User) *services.AuthSvc {
	return services.NewAuthService(services.AuthServiceConfiguration{
		UserRepo:     &MockedUserRepo{Users: initalUsers},
		JWTKey:       services.DefaultJWTKey,
		PwdHashRound: services.DefaultPwdHashRound,
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got error: %v, Want error: %v", got, want)
	}
}

func assertUser(t testing.TB, userGot, userWant *entities.User) {
	t.Helper()
	if match := reflect.DeepEqual(userGot, userWant); !match {
		t.Errorf("Got user is different from user that exist, got: %v, want: %v", userGot, userWant)
	}
}

func getHashedPassword(pwd string) string {
	if pwd == "phuonghau" {
		return "$2y$10$nxL74V8dU1F7DAz9xfaQvuZTJ27ZmUNa4sucT2bV1Vtq.MI8EbX/a"
	}
	log.Fatalf("Hashed string for password: '%q' is unimplemented", pwd)
	return ""
}

func TestFindUserByIDNPwd(t *testing.T) {

	t.Run("FindUserByIDNPwd should return nil, nil if no user presents in database", func(t *testing.T) {
		authSvc := newDefaultAuthSvc([]entities.User{})

		foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), "phuonghau", "phuonghau")
		assertError(t, err, nil)
		assertUser(t, foundUser, nil)
	})

	t.Run("FindUserByIDNPwd username and password", func(t *testing.T) {
		existingUser := entities.User{ID: "phuonghau", Password: getHashedPassword("phuonghau")}
		authSvc := newDefaultAuthSvc([]entities.User{existingUser})
		t.Run("Should be able to find a user that already exists", func(t *testing.T) {
			foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), existingUser.ID, "phuonghau")
			assertError(t, err, nil)
			assertUser(t, foundUser, &existingUser)

		})

		t.Run("Only password matched cannot get a user", func(t *testing.T) {
			foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), "foobarfoo", existingUser.Password)
			assertError(t, err, nil)
			assertUser(t, foundUser, nil)
		})

		t.Run("Only id matched cannot get a user", func(t *testing.T) {
			foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), existingUser.ID, "asdasdsa")
			assertError(t, err, nil)
			assertUser(t, foundUser, nil)
		})
	})

	t.Run("FindUserByIDNPwd with either username or password", func(t *testing.T) {
		authSvc := newDefaultAuthSvc([]entities.User{})
		t.Run("Empty username", func(t *testing.T) {
			foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), "", "phuonghau")
			assertError(t, err, services.ErrServiceUnhandledException)
			assertUser(t, foundUser, nil)
		})
		t.Run("Empty password", func(t *testing.T) {
			foundUser, err := authSvc.FindUserByIDNPwd(context.TODO(), "phuonghau", "")
			assertError(t, err, services.ErrServiceUnhandledException)
			assertUser(t, foundUser, nil)
		})
	})
}

func TestSignJWT(t *testing.T) {
	t.Run("Issued token must contains user_id key as a string", func(t *testing.T) {
		existingUser := entities.User{ID: "phuonghau", Password: getHashedPassword("phuonghau")}
		authSvc := newDefaultAuthSvc([]entities.User{})
		token, err := authSvc.SignJWT(context.TODO(), existingUser.ID)

		claims := make(jwt.MapClaims)
		_, err = jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(services.DefaultJWTKey), nil
		})
		userId, ok := claims["user_id"].(string)
		if !ok {
			t.Fatal("user_id is not visible in payload or either it's not string")
		}
		if userId != existingUser.ID {
			t.Errorf("user_id is not identical to the one that provided, got: %v, want: %v", userId, existingUser.ID)
		}
		assertError(t, err, nil)
	})

	t.Run("Return unhandled exception if user id is null", func(t *testing.T) {
		authSvc := services.NewAuthService(services.AuthServiceConfiguration{
			UserRepo: &MockedUserRepo{},
			JWTKey:   services.DefaultJWTKey,
		})

		token, err := authSvc.SignJWT(context.TODO(), "")
		if len(token) != 0 {
			t.Errorf("Token should not be issued but got: %s", token)
		}
		assertError(t, err, services.ErrServiceUnhandledException)
	})

}

func TestComparePassword(t *testing.T) {
	providedPwd := "phuonghau"

	t.Run("identical passwords (hashed, raw) as arguments must return true", func(t *testing.T) {
		authSvc := newDefaultAuthSvc([]entities.User{})
		want := true

		hashedPwdBytes, _ := bcrypt.GenerateFromPassword([]byte(providedPwd), services.DefaultPwdHashRound)
		got := authSvc.ComparePassword(providedPwd, string(hashedPwdBytes))

		if got != want {
			t.Errorf("Identical passwords (one raw, one hashed with defaultPwdHashRound) but received false")
		}
	})

	t.Run("Non-identical passwords (hashed, raw) as arguments must return false", func(t *testing.T) {
		authSvc := newDefaultAuthSvc([]entities.User{})
		want := false

		hashedPwdBytes, _ := bcrypt.GenerateFromPassword([]byte(providedPwd), services.DefaultPwdHashRound)
		got := authSvc.ComparePassword("changedPassword", string(hashedPwdBytes))

		if got != want {
			t.Errorf("Identical passwords (one raw, one hashed with defaultPwdHashRound) but received false")
		}
	})
}

func TestAuthenticateUser(t *testing.T) {
	t.Run("It should be able to issue token based on provided username and password", func(t *testing.T) {
		existingUser := entities.User{ID: "phuonghau", Password: getHashedPassword("phuonghau")}
		defaultAuthSvc := newDefaultAuthSvc([]entities.User{existingUser})

		token, err := defaultAuthSvc.AuthenticateUser(context.TODO(), existingUser.ID, "phuonghau")
		assertError(t, err, nil)
		if len(token) == 0 {
			t.Errorf("Token should be issued, but none received")
		}
	})
}

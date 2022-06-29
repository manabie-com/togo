package usecase

import (
	"errors"
	"lntvan166/togo/internal/domain"
	e "lntvan166/togo/internal/entities"
	"lntvan166/togo/pkg"
)

type userUsecase struct {
	userRepo domain.UserRepository
	taskRepo domain.TaskRepository
	crypto   domain.AppCrypto
}

func NewUserUsecase(repo domain.UserRepository, taskRepo domain.TaskRepository, crypto domain.AppCrypto) *userUsecase {
	return &userUsecase{
		userRepo: repo,
		taskRepo: taskRepo,
		crypto:   crypto,
	}
}

func (u *userUsecase) Register(user *e.User) error {
	checkUserExist := u.CheckUserExist(user.Username)

	if checkUserExist {
		// pkg.ERROR(w, http.StatusBadRequest, errors.New("user already exist"), "")
		return errors.New("user already exist")
	}

	user.Password = u.crypto.HashPassword(user.Password)

	// err := user.IsValid()
	// if err != nil {
	// 	// pkg.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
	// 	return err
	// }

	// TODO: check valid user

	err := u.userRepo.AddUser(user)
	if err != nil {
		// pkg.ERROR(w, http.StatusInternalServerError, err, "failed to add user!")
		return err
	}

	return nil
}

func (u *userUsecase) Login(user *e.User) (string, error) {
	checkUserExist := u.CheckUserExist(user.Username)

	if !checkUserExist {
		// pkg.ERROR(w, http.StatusBadRequest, errors.New("user not found"), "")
		return "", errors.New("user not found")
	}

	oldUser, err := u.userRepo.GetUserByName(user.Username)
	if err != nil {
		// pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return "", errors.New("failed to find user")
	}
	check := u.crypto.ComparePassword(user.Password, oldUser.Password)
	if !check {
		// pkg.ERROR(w, http.StatusBadRequest, errors.New("password incorrect"), "")
		return "", errors.New("password incorrect")
	}

	token, err := pkg.GenerateToken(user.Username)
	if err != nil {
		// pkg.ERROR(w, http.StatusInternalServerError, err, "failed to generate token!")
		return "", err
	}

	return token, nil
}

func (u *userUsecase) GetAllUsers() ([]*e.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id int) (*e.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *userUsecase) GetUserByName(username string) (*e.User, error) {
	user, err := u.userRepo.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserIDByUsername(username string) (int, error) {
	user, err := u.userRepo.GetUserByName(username)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *userUsecase) GetMaxTaskByUserID(id int) (int, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	return int(user.MaxTodo), nil
}

func (u *userUsecase) GetPlan(username string) (string, error) {
	user, err := u.userRepo.GetUserByName(username)
	if err != nil {
		return "", errors.New("user not found")
	}
	return user.Plan, nil
}

func (u *userUsecase) UpdateUser(user *e.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *userUsecase) UpgradePlan(userID int, plan string, maxTodo int) error {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	if user.Plan == plan {
		return errors.New("plan already upgraded")
	}

	user.Plan = plan
	user.MaxTodo = int64(maxTodo)
	return u.userRepo.UpdateUser(user)
}

func (u *userUsecase) CheckUserExist(username string) bool {
	user, err := u.userRepo.GetUserByName(username)
	if err != nil {
		return false
	}
	if user.ID == 0 {
		return false
	}
	return true
}

func (u *userUsecase) DeleteUserByID(id int) error {
	err := u.taskRepo.DeleteAllTaskOfUser(id)
	if err != nil {
		return err
	}
	return u.userRepo.DeleteUserByID(id)
}

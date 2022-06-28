package controller

import (
	"encoding/json"
	"lntvan166/togo/internal/domain"
	"lntvan166/togo/pkg"
	"net/http"

	e "lntvan166/togo/internal/entities"
)

type AuthController struct {
	UserUsecase domain.UserUsecase
}

func NewAuthController(userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		UserUsecase: userUsecase,
	}
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	user := e.NewUser()
	var err error

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	err = u.UserUsecase.Register(user)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to register user!")
		return
	}

	pkg.JSON(w, http.StatusCreated, "Register Successfully")
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	newUser := e.NewUser()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	token, err := u.UserUsecase.Login(newUser)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to login user!")
		return
	}

	pkg.JSON(w, http.StatusOK, map[string]string{"token": token, "message": "login successfully"})
}

// func (u *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
// 	user, err := u.UserUsecase.GetUserByName(context.Get(r, "username").(string))
// 	if err != nil {
// 		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
// 		return
// 	}

// 	err = json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
// 		return
// 	}

// 	err = user.IsValid()
// 	if err != nil {
// 		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
// 		return
// 	}

// 	user.PreparePassword()

// 	err = u.UserUsecase.UpdateUser(user)
// 	if err != nil {
// 		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to update user!")
// 		return
// 	}

// 	pkg.JSON(w, http.StatusOK, "Update Password Successfully")
// }

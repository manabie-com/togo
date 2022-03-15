package controllers

import (
	"togo-service/app/models"
	requests "togo-service/app/request"
	"togo-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (handler Handler) Login(ctx *fiber.Ctx) error {
	type LoginUser struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	var login LoginUser
	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid credentials",
		})
	}
	var usermodel models.User

	handler.DB.Where("username = ?", login.Username).Find(&usermodel)

	if usermodel.ID == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(login.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid credentials",
		})
	}

	// Generate a new Access token.
	token, err := utils.CreateToken(usermodel.ID, usermodel.Role == "admin")
	if err != nil {
		// Return status 500 and token generation error.
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error":        false,
		"message":      nil,
		"access_token": token,
	})
}

func (handler Handler) Register(ctx *fiber.Ctx) error {
	var userParam requests.CreateUserParam
	ctx.BodyParser(&userParam)

	if err := userParam.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	var usermodel models.User
	handler.DB.Where("username = ?", userParam.Username).Find(&usermodel)

	if usermodel.ID > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Username is used by another user.",
		})
	}

	hash, err := utils.HashPassword(userParam.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	userParam.Password = hash

	usermodel.Name = userParam.Name
	usermodel.Username = userParam.Username
	usermodel.Password = userParam.Password
	usermodel.Role = "user"

	if result := handler.DB.Create(&usermodel); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": result.Error,
		})
	}

	var userSetting models.Setting
	userSetting.UserID = uint64(usermodel.ID)
	userSetting.QuotaPerDay = 10

	handler.DB.Create(&userSetting)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Created user successfully",
	})
}

// UpdateUser Setting
func (handler Handler) UpdateUser(ctx *fiber.Ctx) error {
	//Extract the access token metadata
	tokenAuth, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if tokenAuth.IsAdmin == false {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You don't have permission",
		})
	}

	var settingParm requests.UserSettingPararm
	ctx.BodyParser(&settingParm)

	if err := settingParm.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var user models.User
	err = handler.DB.Preload("Setting").Model(&user).Find(&user, settingParm.UserID).Error

	if user.ID == 0 || err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}
	user.Setting.QuotaPerDay = uint64(settingParm.QuotaPerDay)

	handler.DB.Save(&user.Setting)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Updated",
		"data":    user.Setting,
	})
}

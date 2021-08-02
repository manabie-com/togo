package daos

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/manabie-com/togo/internal/database"
	models "github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/utils"
)

type AccountDAO struct {
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	TableName string
}

func (u *AccountDAO) CreateAccount(account models.Account) (*models.Account, error) {
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	//validate if account already exists
	accountAvailable, _ := accountDAO.FindAccountByUsername(account)
	if accountAvailable.Username != "" {
		return &account, errors.New("account unavailable")
	}
	//Generate username if reqbody doesn't have one
	if account.Username == "" {
		currentTimeMillis := utils.GetCurrentEpochTimeInMiliseconds()
		newUsername := account.Username + strconv.FormatInt(currentTimeMillis, 10)
		account.Username = newUsername
	}
	//if id is not inputted, generate one.
	if account.ID.String() == "00000000-0000-0000-0000-000000000000" {
		account.ID = uuid.New()
	}
	err = db.Debug().Model(&models.Account{}).Create(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (u *AccountDAO) FindAccountByID(id uuid.UUID) (*models.Account, error) {
	accountResult := models.Account{}
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Model(&models.Account{}).Preload("Tasks").
		First(&accountResult, "id=?", fmt.Sprint(id)).Error
	if err != nil {
		return nil, errors.New("account not found based on credentials")
	}
	return &accountResult, err
}

func (u *AccountDAO) FindAccountByUsernameAndPassword(accountInfo models.Account) (*models.Account, error) {
	result := models.Account{}
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Model(&models.Account{}).
		First(&result, "username=?", accountInfo.Username).Error
	if err != nil {
		return nil, errors.New("account not found based on credentials")
	}
	//compare password from db to input
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(accountInfo.Password))
	if err != nil {
		return nil, errors.New("account not found based on credentials")
	}
	return &result, err
}

func (u *AccountDAO) FindAccountByUsername(account models.Account) (*models.Account, error) {
	accountResult := models.Account{}
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Model(&models.Account{}).
		First(&accountResult, "username=?", account.Username).Error
	return &accountResult, err
}

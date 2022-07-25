package model

import (
	"context"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:user"`
	ID            int       `bun:"id,pk,autoincrement"`
	Email         string    `bun:"email"`
	Password      string    `bun:"password"`
	IsActive      bool      `bun:"is_active,default:false"`
	DailyLimit    int       `bun:"daily_limit,default:3"`
	Tasks         []*Task   `bun:"rel:has-many,join:id=created_by"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

// FindOneUserByID /**
func FindOneUserByID(ctx context.Context, id int, db bun.IDB) (User, error) {
	user := User{
		ID: id,
	}

	err := db.NewSelect().Model(&user).WherePK().Scan(ctx)
	return user, err
}

// FindOneUserByEmail /**
func FindOneUserByEmail(ctx context.Context, email string, db bun.IDB) (User, error) {
	user := User{
		Email: email,
	}

	err := db.NewSelect().Model(&user).Where("email=?", user.Email).Scan(ctx)
	return user, err
}

// Create /**
func (user *User) Create(ctx context.Context, db bun.IDB) error {
	// hash password using bcrypt before saving
	if hashedPassword, err := HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
	}
	user.IsActive = true // default is true when user register
	user.DailyLimit = 3  // default daily limit is 3
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// HashPassword /**
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash /**
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsExceededDailyLimit /**
func (user *User) IsExceededDailyLimit(ctx context.Context, db bun.IDB) bool {
	if user.countTasksCurrentDate(ctx, db) < user.DailyLimit {
		return false
	}
	return true
}

// countTasksByDate /**
func (user *User) countTasksByDate(date string, ctx context.Context, db bun.IDB) int {
	count, err := db.NewSelect().
		Model((*Task)(nil)).
		Where("published_date=?", date).
		Where("created_by=?", user.ID).
		Count(ctx)
	if err != nil {
		log.Println(err)
		return 0
	}

	return count
}

// countTasksCurrentDate /**
func (user *User) countTasksCurrentDate(ctx context.Context, db bun.IDB) int {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	return user.countTasksByDate(currentDate, ctx, db)
}

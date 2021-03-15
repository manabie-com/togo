package storages

// import "github.com/jinzhu/gorm"

// Behind gorm.Model
/*
type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`
}
*/

// Task reflects tasks in DB
type Task struct {
	// gorm.Model
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	// gorm.Model
	ID                string `json:"id"`
	Password          string `json:"password"`
	CurrentNumberTask int16  `json:"current_number_task"`
}

// ConfigServer for server
type ConfigServer struct {
	Value int16  `json:"value"`
	Name  string `json:"name"`
}

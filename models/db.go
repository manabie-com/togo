package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Create a Db connection
type DbConn struct {
	DB *sql.DB
}
type BaseHandler struct {
	BaseCtrl *DbConn
}

// newdbConn returns a new DbConn
func NewdbConn(db *sql.DB) *DbConn {
	return &DbConn{
		DB: db,
	}
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(BC *DbConn) *BaseHandler {
	return &BaseHandler{
		BaseCtrl: BC,
	}
}

// connect to database
func Connect(DB_URI string) *DbConn{ 	
	db, err := sql.Open("postgres", DB_URI)
	dbconn := NewdbConn(db)
	if err != nil {
		log.Fatal("Connect to database failed, err: "+err.Error())
	}
	return dbconn
}

// Hash password into a crypt text
func Hash(password string) (string, error) { 
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

 // check a crypted text and a password user enter
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// interface implement
// func (r *Repository)Close() {
// 	r.DB.Close()
// }
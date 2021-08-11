package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

const ContextKeyTransaction = "ctx_transaction"

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		tx := db.Begin()
		log.Println("Begin transaction")
		context.Set(ContextKeyTransaction, tx)
		context.Next()
		if context.Writer.Status() == http.StatusOK {
			if err := tx.Commit().Error; err != nil {
				log.Println("Commit error: ", err)
				return
			}
			log.Println("Commit transaction")
			return
		}
		log.Println("Rollback transaction")
		tx.Rollback()
	}
}

func DbFromContext(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(ContextKeyTransaction).(*gorm.DB)
	return db
}

package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id uint			`json:"id" gorm:"primary_key;column:user_id"`
	Email string	`json:"email"`
	Password string	`json:"-"`
	JoinedAt string	`json:"joinedAt" gorm:"column:created_at;<-:false"`
}

func main () {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Mana xato: %v", err)
	}

	PORT := ":4000"

	r := gin.Default()

	r.GET("/users", func (ctx *gin.Context) {

		var users []User

		db.Find(&users)

		ctx.IndentedJSON(http.StatusOK, users)
	})

	fmt.Printf("Server ready at %s\n", PORT)

	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}
}

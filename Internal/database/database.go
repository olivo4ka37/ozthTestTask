package database

import (
	"PostCommentService/Internal/graph/model"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Store interface {
	GetPosts() ([]*model.Post, error)
	GetPost(id int) (*model.Post, error)
	GetComments(postID int) ([]*model.Comment, error)
	GetComment(id int) (*model.Comment, error)
	CreatePost(title, content, author string) (*model.Post, error)
	CreateComment(postID int, author, content string, parentId *int) (*model.Comment, error)
	UpdatePost(id int, title, content string) (*model.Post, error)
	UpdateComment(id int, content string) (*model.Comment, error)
	DisableComments(postID int) error
}

func NewStore(useMemory bool) Store {
	if useMemory {
		return NewMemoryStore()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected!")
	return NewPostgresStore(db)
}

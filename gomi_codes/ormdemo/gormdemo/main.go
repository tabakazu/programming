package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" form:"title" json:"title"`
	Content string `gorm:"not null" form:"content" json:"content" sql:"type:text;"`
}

func CreateDb() error {
	db, err := sql.Open("mysql", "root:@/")
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS gormdemo_development")
	if err != nil {
		return err
	}
	return nil
}

func MigrateDbSchema(db *gorm.DB) error {
	if err := db.AutoMigrate(&Post{}).Error; err != nil {
		return err
	}
	return nil
}

func init() {
	if err := CreateDb(); err != nil {
		panic(err.Error())
	}
}

type PostRepo struct{}

func newPostRepo() *PostRepo {
	return new(PostRepo)
}

func (r PostRepo) CreatePost(post Post, db *gorm.DB) error {
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (r PostRepo) Count(db *gorm.DB) (int, error) {
	var count int
	if err := db.Table("posts").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	db, err := gorm.Open("mysql", "root:@/gormdemo_development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := MigrateDbSchema(db); err != nil {
		panic(err.Error())
	}

	repo := newPostRepo()
	post := Post{Title: "Post 1", Content: "Body 1"}
	if err := repo.CreatePost(post, db); err != nil {
		panic(err.Error())
	}

	count, err := repo.Count(db)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("- Post count: %v", count)
}

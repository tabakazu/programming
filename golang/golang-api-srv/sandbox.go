package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	// curl -X GET http://localhost:3000/posts -H 'content-type: application/json'
	r.GET("/posts", getPostListHandler)
	// curl -X GET http://localhost:3000/posts/1 -H 'content-type: application/json'
	r.GET("/posts/:id", getPostHandler)
	// curl -X POST http://localhost:3000/posts -H 'content-type: application/json' -d '{ "title": "test title", "body": "test body" }'
	r.POST("/posts", createPostHandler)
	// curl -X PUT http://localhost:3000/posts/1 -H 'content-type: application/json' -d '{ "title": "new test title", "body": "new test body" }'
	r.PUT("/posts/:id", updatePostHandler)
	// curl -X DELETE http://localhost:3000/posts/1 -H 'content-type: application/json'
	r.DELETE("/posts/:id", destroyPostHandler)

	r.Run(":3000")
}

//
// entity
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
type Post struct {
	Model
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

//
// http handler
// Index
func getPostListHandler(c *gin.Context) {
	var posts []Post

	db, _ := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Show
func getPostHandler(c *gin.Context) {
	var p Post

	db, _ := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err := db.First(&p, c.Params.ByName("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// Create
func createPostHandler(c *gin.Context) {
	var p Post

	db, _ := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Update
func updatePostHandler(c *gin.Context) {
	var p Post
	var newp Post

	db, _ := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err := c.ShouldBindJSON(&newp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&p, c.Params.ByName("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&p).Updates(newp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Delete
func destroyPostHandler(c *gin.Context) {
	db, _ := gorm.Open("mysql", "root:@/golang_api_srv_dev?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err := db.Where("id = ?", c.Params.ByName("id")).Delete(Post{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

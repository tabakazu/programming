package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/posts", PostsIndex)
	// r.GET("/posts/:id", PostsShow)
	// r.POST("/posts", PostsCreate)
	// r.PUT("/posts/:id", PostsUpdate)
	// r.DELETE("/posts/:id", PostsDestroy)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/posts", ApiPostsIndex)
		}
	}

	r.Run(":3000")
}

// http://localhost:3000/posts
func PostsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post List",
	})
}

// http://localhost:3000/api/v1/posts
func ApiPostsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Api::V1::Post List",
	})
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/MilesChou/go-post/entity"
	"strconv"
	"net/http"
)

func CreateServer() *gin.Engine {
	server := gin.Default()
	server.GET(`/ping`, func(c *gin.Context) {
		s := make([]string, 1)
		s[0] = `pong`

		c.JSON(200, map[string]string{
			"result": "pong",
		})
	})

	server.POST("/posts", func(c *gin.Context) {
		name := c.PostForm("name")
		content := c.PostForm("content")

		post := entity.CreatePost(name, content)

		c.JSON(http.StatusOK, gin.H{
			"id": post.ID,
			"createdAt": post.CreatedAt,
			"updatedAt": post.UpdatedAt,
			"deletedAt": post.DeletedAt,
			"name": post.Name,
			"content": post.Content,
		})
	})

	server.GET("/posts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			panic("failed parse int")
		}

		post := entity.ReadPost(uint(id))

		if post.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Resource is not found",
			})
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": post.ID,
			"createdAt": post.CreatedAt,
			"updatedAt": post.UpdatedAt,
			"deletedAt": post.DeletedAt,
			"name": post.Name,
			"content": post.Content,
		})
	})

	return server
}

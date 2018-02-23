package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/MilesChou/go-post/entity"
	"strconv"
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

		c.JSON(200, post)
	})

	server.GET("/posts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			panic("failed parse int")
		}

		post := entity.ReadPost(uint(id))

		c.JSON(200, post)
	})

	return server
}

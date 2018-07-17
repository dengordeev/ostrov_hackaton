package main

import "github.com/gin-gonic/gin"

//User - структура, содержащая данные о пользователе
type User struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	Lastname   string     `json:"lastname"`
	Firstname  string     `json:"firsname"`
	About      string     `json:"about"`
	MyProjects []*Project `json:"myprojects"`
	Projects   []*Project `json:"projects"`
}

//Project - структура, содержащая данные о конкретном проекте
type Project struct {
	ID     int      `json:"id"`
	IsOpen bool     `json:"isopen"`
	Name   string   `json:"name"`
	About  string   `json:"about"`
	Part   []*Users `json:"party"`
}

type Users []*User
type Projects []*Project

var (
	users    = new(Users)
	projects = new(Projects)
)

func main() {
	r := gin.Default()

	r.Static("/app", "./static")
	/*r.POST("/teamleaders", func(c *gin.Context) {

	})
	r.POST("/users", func(c *gin.Context) {

	})
	*/
	r.Any("/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"projects": projects,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

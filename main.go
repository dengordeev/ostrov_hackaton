package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	ID       int     `json:"id"`
	IsOpen   bool    `json:"isopen"`
	Name     string  `json:"name"`
	About    string  `json:"about"`
	Part     []*User `json:"party"`
	Template `json:"template"`
}


var (
	templates = []*Template
	users    = []*User{}
	projects = []*Project{}
)

func findTeamLeaders(users []*User) []*User {
	result := []*User{}
	for _, u := range users {
		if len(u.MyProjects) != 0 {
			result = append(result, u)
		}
	}

	return result
}

func listTemplates() []*Template {

}
func main() {
	r := gin.Default()

	r.Static("/app", "./static")

	projects = append(projects, p1)
	users = append(users, u1, u2)
	r.GET("/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"projects": projects,
		})
	})
	r.PUT("/projects", func(c *gin.Context) {
		project := &Project{}
		err := c.BindJSON(&project)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, nil)
		}
		project.ID = len(projects) + 1
		projects = append(projects, project)

		c.JSON(http.StatusOK, nil)
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	r.PUT("/users", func(c *gin.Context) {
		user := &User{}
		err := c.BindJSON(&user)
		log.Println(user)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		user.ID = len(users) + 1
		users = append(users, user)

		c.JSON(http.StatusOK, nil)
	})
	r.GET("/teamleaders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"teamleaders": findTeamLeaders(users),
		})
	})

	r.PUT("/templates", func(c *gin.Context) {
		template := &Template{}

		err := c.BindJSON(&template)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		template.ID = len(templates) + 1

		templates = append(templates, template)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	templateStore = "./db/template.json"
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

//Role - структура, показывающая данные о ролях в проекте
type Role struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

//Template - структура шаблона
type Template struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	AuthorID      int      `json:"authorId"`
	AuthorInfo    string   `json:"authorInfo"`
	VideoList     []string `json:"videoList"`
	CourseLink    string   `json:"courseLink"`
	Roles         []*Role  `json:"roles"`
	Additional    []string `json:"additional"`   //материалы
	Resources     []string `json:"resources"`    //ресурсы
	ProjectsCount int      `json:"projectCount"` //количество успешных проектов
}

func save() {
	str, err := json.Marshal(&templates)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(templateStore, str, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

var (
	users     = []*User{}
	projects  = []*Project{}
	templates = []*Template{}
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

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Static("/app", "./static")

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
		user.ID = users[len(users)-1].ID + 1
		users = append(users, user)

		c.JSON(http.StatusOK, nil)
	})
	r.GET("/teamleaders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"teamleaders": findTeamLeaders(users),
		})
	})

	r.POST("/templates/add", func(c *gin.Context) {
		template := &Template{}

		log.Println(template)
		err := c.BindJSON(&template)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		if len(templates) == 0 {
			template.ID = 1
		} else {
			template.ID = templates[len(templates)-1].ID + 1
		}

		templates = append(templates, template)
	})

	r.POST("/templates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"templates": templates,
		})
	})

	r.GET("/template/:id", func(c *gin.Context) {
		tid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		for _, t := range templates {
			if t.ID == tid {
				c.JSON(http.StatusOK, gin.H{
					"template": t,
				})
				return
			}
		}

		c.JSON(http.StatusOK, nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

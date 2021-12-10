package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/internal/http"
	"github.com/manabie/project/model"
	"github.com/manabie/project/pkg/snowflake"
	"strconv"
	"time"
)

type router struct {
	http 		http.Http
	snowflake 	snowflake.SnowflakeData
}

type Router interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	TaskAll(c *gin.Context)
	TaskById(c *gin.Context)
}

func NewRouter(http http.Http, snowflake snowflake.SnowflakeData) Router {
	return &router{
		http: 		http,
		snowflake:	snowflake,
	}
}

func(r *router) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed " + err.Error(),
			"date": time.Now(),
		})
		return
	}
	token, err := r.http.Login(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message" : "create token failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token" : token,
		"message": "user login success",
		"time": time.Now(),
	})
}

func(r *router) SignUp(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userData := model.User{
		Id:       r.snowflake.GearedID(),
		Username: username,
		Password: password,
		MaxTodo:  0,
	}
	if err := r.http.SignUp(userData); err != nil {
		c.JSON(500, gin.H{
			"message": "signup user failed " + err.Error(),
			"date": time.Now(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "create user success",
		"time": time.Now(),
	})
}

func(r *router) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed " + err.Error(),
			"date": time.Now(),
		})
		return
	}
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	taskTodo := model.Task{
		Id:          r.snowflake.GearedID(),
		Content:     task.Content,
		UserId:      userId,
		CreatedDate: time.Now(),
	}
	if err := r.http.CreateTask(taskTodo, userId); err != nil {
		c.JSON(500, gin.H{
			"message": "create task failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "create task success",
		"user_id": userId,
		"date": time.Now(),
	})
}

func(r *router) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	var task model.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	if err := r.http.UpdateTask(id, task); err != nil {
		c.JSON(500, gin.H{
			"message": "update data failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update data success",
		"id": id,
		"date": time.Now(),
	})
}

func(r *router) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	if err := r.http.DeleteTask(id); err != nil {
		c.JSON(500, gin.H{
			"message": "delete data task failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "delete data task success",
		"id": id,
		"date": time.Now(),
	})
}

func(r *router) TaskAll(c *gin.Context) {
	tasks, err := r.http.TaskAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": tasks,
		"date": time.Now(),
	})
}

func(r *router) TaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "get data param failed "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	task, err := r.http.TaskById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "get list task no data "+ err.Error(),
			"date": time.Now(),
		})
		return
	}
	c.JSON(200, gin.H{
		"id": id,
		"data":task,
		"date": time.Now(),
	})
}
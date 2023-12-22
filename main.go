package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	//	"github.com/teris-io/shortid"
)

type Task struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

var tasks = make(map[string]Task)
var conn *pgx.Conn

// logfile, err := os.Create("app.log")
func initDB() {
	var err error
	connectionString := "postgresql://postgres:Golang@localhost:5433/TaskScheduler"
	conn, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
func CreateTask(c *gin.Context) {
	requestBody, _ := io.ReadAll(c.Request.Body)
	fmt.Printf("Received: %s\n", requestBody)
	// Reset the request body so it can be read again
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate an ID for the task (for simplicity, we use a timestamp).

	//id, err := shortid.Generate()
	initDB()
	err := conn.QueryRow(context.Background(), `
        INSERT INTO tasks (title, description, due_date)
        VALUES ($1, $2, $3)
        RETURNING id;
    `, task.Title, task.Description, task.DueDate).Scan(&task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate task ID"})
		return
	}
	//task.ID = id

	//tasks[task.ID] = task
	c.JSON(http.StatusCreated, task)
	fmt.Println(c.Request.Body)
	log.Println(task)
}
func GetAllTasks(c *gin.Context) {
	var tasks []Task
	initDB()
	rows, err := conn.Query(context.Background(), "SELECT ID, title, description,due_date from tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query tasks from database"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan task from row"})
			return
		}
		tasks = append(tasks, task)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error iterating over tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func GetTask(c *gin.Context) {
	var task Task
	// Retrieve the task ID from the route parameter
	id := c.Param("id")
	initDB()
	err := conn.QueryRow(context.Background(), "SELECT ID, title, description,due_date from tasks where ID = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate)
	log.Println(id)
	if err == pgx.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query tasks from database"})
		return
	}
	c.JSON(http.StatusOK, task)
}
func main() {
	r := gin.Default()

	/*r.GET("/ping", func(c *gin.Context) {
	    c.JSON(200, gin.H{
	        "message": "pong",
	    })
	})*/
	r.POST("/tasks", CreateTask)
	r.GET("/tasks", GetAllTasks)
	r.GET("/tasks/:id", GetTask)
	/*  r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)*/
	r.Run() // By default it serves on :8080
}

****TaskScheduler App****

**Overview**
TaskScheduler is a RESTful web application that allows users to create, retrieve, and manage tasks. The application is built using the Go programming language, the Gin framework for HTTP routing, and PostgreSQL as the database.

**Features**
Create Task: Users can create new tasks with a title, description, and due date.
Retrieve All Tasks: Users can retrieve a list of all tasks stored in the database.
Retrieve Single Task: Users can retrieve details of a specific task by its ID.

**Components**

**Gin Framework**
The Gin framework is a lightweight and fast HTTP web framework for Go. It is used for routing and handling HTTP requests in the TaskScheduler app.

**PostgreSQL**
PostgreSQL is an open-source relational database system used to store and manage task data. The pgx package is used to interact with the PostgreSQL database.

**Task Struct**
The Task struct represents the structure of a task with the following fields:

**ID** (int32): The unique identifier of the task.
**Title** (string): The title of the task.
**Description** (string): A brief description of the task.
**DueDate **(time.Time): The due date of the task.
**Routes**
**POST /tasks**: Creates a new task.
**GET /tasks**: Retrieves all tasks.
**GET /tasks/:id**: Retrieves a task by its ID.
**Installation and Setup**
**Clone the repository:**

sh
Copy code
git clone <repository-url>
cd TaskScheduler
**Install dependencies:**
Make sure you have Go and PostgreSQL installed on your machine.

**Set up the database:**
Create a PostgreSQL database and update the connection string in the initDB function:

go
Copy code
connectionString := "postgresql://<username>:<password>@localhost:<port>/<database>"
**Run the application:**

sh
Copy code
go run main.go
Usage
**Create a new task:**
Send a POST request to http://localhost:8080/tasks with the task details in the request body.

**Retrieve all tasks:**
Send a GET request to http://localhost:8080/tasks.

**Retrieve a specific task by ID:**
Send a GET request to http://localhost:8080/tasks/{id}, replacing {id} with the task ID.

**Logging**
The application logs important events and errors to help with debugging and monitoring.

**Contribution**
Feel free to fork this repository, make your changes, and create a pull request. Contributions are welcome!
**
License**
This project is licensed under the MIT License. See the LICENSE file for details.

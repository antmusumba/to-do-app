## To-Do App

A simple to-do application with a CLI and an API.
## Description

This project is a to-do application that allows users to manage their tasks. It provides both a CLI (Command Line Interface) for interacting with the application from the command line and an API for programmatic access. The application is built using Go, with tasks stored in a file-based system.
## 🚀 Features
## CLI

The CLI is located in the cmd/cli directory. It provides the following commands:

    add: Adds a new task.

        Example:

    go run cmd/cli/main.go add "Learn Go CLI"

    Adds the task with the given title.

list: Lists all tasks.

    Example:

    go run cmd/cli/main.go list

    Displays a list of all tasks, including their ID, title, completion status, and deadline.

complete: Marks a task as complete.

    Example:

    go run cmd/cli/main.go complete 1

    Marks the task with the given ID as completed.

delete: Deletes a task.

    Example:

        go run cmd/cli/main.go delete 1

        Deletes the task with the specified ID.

## API

The API is located in the cmd/api directory. It provides the following endpoints:

    GET /tasks: Lists all tasks.

        Example request:

    curl http://localhost:8080/tasks

    Returns a JSON array of all tasks stored in the system.

POST /tasks: Adds a new task.

    Example request:

    curl -X POST -H "Content-Type: application/json" \
      -d '{"title": "Learn Go REST API"}' \
      http://localhost:8080/tasks

    Adds a task with the title provided in the request body.

PUT /tasks/:id: Updates a task (e.g., mark it as completed).

    Example request:

    curl -X PUT -H "Content-Type: application/json" \
      -d '{"completed": true}' \
      http://localhost:8080/tasks/1

    Marks the task with the given ID as completed.

DELETE /tasks/:id: Deletes a task.

    Example request:

        curl -X DELETE http://localhost:8080/tasks/1

        Deletes the task with the given ID.

## Internal Packages

The project uses the following internal packages:

    internal/task: Defines the Task type and provides functions for working with tasks, including creating, updating, and deleting tasks.

    internal/storage: Provides an interface for storing tasks, using a simple file-based storage mechanism (currently tasks.json).

    internal/handler: Provides handlers for the API endpoints, handling HTTP requests and performing operations like creating, updating, or listing tasks.

## 🛠️ Running the Project

    Clone the repository:

git clone <repo-url>
cd to-do-app

Run the server (for the API):

go run cmd/api/main.go

The server will be running at http://localhost:8080.

Test the API:
Use curl commands to interact with the API as described in the API section above.

Use the CLI:

    Run the following commands to interact with the to-do application:

go run cmd/cli/main.go <command> <arguments>

For example:

    To add a task:

go run cmd/cli/main.go add "New Task"

To list tasks:

go run cmd/cli/main.go list

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

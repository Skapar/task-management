# Task Management API

This is a RESTful API for managing users, tasks, and projects. It supports CRUD operations and is built with Go, using Gorilla Mux for routing and PostgreSQL for the database. The application is containerized using Docker and Docker Compose.

## Table of Contents

- [Task Management API](#task-management-api)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
    - [Using Docker](#using-docker)
    - [Without Docker](#without-docker)
  - [API Endpoints](#api-endpoints)
    - [Users](#users)
    - [Tasks](#tasks)
    - [Projects](#projects)
  - [Environment Variables](#environment-variables)
  - [Database Migrations](#database-migrations)
  - [Docker](#docker)
    - [Building and Running the Application](#building-and-running-the-application)

## Prerequisites

- Go 1.26 or higher
- Docker
- Docker Compose
- Postman (for testing)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/task-management.git
    cd task-management
    ```

2. Install Go dependencies:

    ```sh
    go mod download
    ```

3. Create a `.env` file in the root directory and add the following environment variables:

    ```env
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=password
    POSTGRES_DB=task_management
    DATABASE_URL=postgres://postgres:password@db:5432/task_management?sslmode=disable
    ```

## Running the Application

### Using Docker

1. Build and run the Docker containers:

    ```sh
    make build
    make run
    ```

2. The API will be available at `http://localhost:8000`.

### Without Docker

1. Start PostgreSQL and create the `task_management` database.

2. Run database migrations:

    ```sh
    go run cmd/server/main.go
    ```

3. Run the application:

    ```sh
    go run cmd/server/main.go
    ```

4. The API will be available at `http://localhost:8000`.

## API Endpoints

### Users

- **Create User**
  - Method: `POST`
  - URL: `/users`
  - Body:

    ```json
    {
      "name": "John Doe",
      "email": "john@example.com",
      "role": "admin"
    }
    ```

  - Response: `201 Created`

- **Get All Users**
  - Method: `GET`
  - URL: `/users`
  - Response: `200 OK`

- **Get User by ID**
  - Method: `GET`
  - URL: `/users/{id}`
  - Response: `200 OK`

- **Update User**
  - Method: `PUT`
  - URL: `/users/{id}`
  - Body:
  
    ```json
    {
      "name": "John Smith",
      "email": "john.smith@example.com",
      "role": "admin"
    }
    ```

  - Response: `200 OK`

- **Delete User**
  - Method: `DELETE`
  - URL: `/users/{id}`
  - Response: `200 OK`

### Tasks

- **Create Task**
  - Method: `POST`
  - URL: `/tasks`
  - Body:
  
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "priority": "high",
      "status": "new",
      "assignee_id": 1,
      "project_id": 1
    }
    ```

  - Response: `201 Created`

- **Get All Tasks**
  - Method: `GET`
  - URL: `/tasks`
  - Response: `200 OK`

- **Get Task by ID**
  - Method: `GET`
  - URL: `/tasks/{id}`
  - Response: `200 OK`

- **Update Task**
  - Method: `PUT`
  - URL: `/tasks/{id}`
  - Body:

    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "priority": "medium",
      "status": "in-progress"
    }
    ```

  - Response: `200 OK`

- **Delete Task**
  - Method: `DELETE`
  - URL: `/tasks/{id}`
  - Response: `200 OK`

### Projects

- **Create Project**
  - Method: `POST`
  - URL: `/projects`
  - Body:

    ```json
    {
      "title": "Project Title",
      "description": "Project Description",
      "start_date": "2023-01-01T00:00:00Z",
      "end_date": "2023-12-31T23:59:59Z",
      "manager_id": 1
    }
    ```

  - Response: `201 Created`

- **Get All Projects**
  - Method: `GET`
  - URL: `/projects`
  - Response: `200 OK`

- **Get Project by ID**
  - Method: `GET`
  - URL: `/projects/{id}`
  - Response: `200 OK`

- **Update Project**
  - Method: `PUT`
  - URL: `/projects/{id}`
  - Body:

    ```json
    {
      "title": "Updated Project Title",
      "description": "Updated Project Description",
      "start_date": "2023-01-01T00:00:00Z",
      "end_date": "2023-12-31T23:59:59Z",
      "manager_id": 1
    }
    ```

  - Response: `200 OK`

- **Delete Project**
  - Method: `DELETE`
  - URL: `/projects/{id}`
  - Response: `200 OK`

## Environment Variables

- `POSTGRES_USER`: The PostgreSQL user.
- `POSTGRES_PASSWORD`: The PostgreSQL password.
- `POSTGRES_DB`: The PostgreSQL database name.
- `DATABASE_URL`: The connection string for the PostgreSQL database.

## Database Migrations

Migrations are located in the `migrations` directory. They are run automatically when the application starts.

## Docker

### Building and Running the Application

To build and run the application using Docker, use the following commands:

```sh
make build
make run

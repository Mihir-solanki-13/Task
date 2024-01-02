# Go Backend Project: Simple REST API with Database Integration

This is a basic Go backend project that implements a simple REST API with database integration. The project uses the Go programming language, the Gin web framework for building APIs, and Gorm as the Object Relational Mapper (ORM) for database interactions.

## Project Structure

- **`main.go`**: Entry point of the application.
- **`internal/database`**: Database initialization and connection setup.
- **`internal/handlers`**: HTTP request handlers for CRUD operations.
- **`internal/models`**: Definition of the data model (`Task` struct).

# API Endpoints

## 1. Create a Task

**Endpoint:** `POST /tasks`

## 2. Get All Tasks

**Endpoint:** `GET /tasks`

## 3. Get Task by ID

**Endpoint:** `GET /tasks/:id`

## 4. Update a Task

**Endpoint:** `PUT /tasks/:id`

## 5. Delete a Task

**Endpoint:** `DELETE /tasks/:id`

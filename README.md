# Trivia REST API

A **Dockerized REST API** built using **Go Fiber**, **PostgreSQL**, and **Docker** for managing trivia questions and answers. This project includes a reliable and scalable architecture using Go Fiber, manages structured data with PostgreSQL, and is deployed in a containerized environment using Docker.

![thumbnail](https://github.com/tunde262/trivia_rest_api_docker_postgres_go/blob/main/assets/thumbnail.png?raw=true)

## Project Goal

The goal of this project was to implement a scalable and containerized **REST API Solution** with a structured database schema. The API solves a real-world problem by efficiently managing trivia question and answer data.

## Features

- **Go Fiber Framework**: High-performance web framework for building RESTful APIs.
- **PostgreSQL**: Structured data storage with GORM for efficient ORM capabilities.
- **Dockerized Deployment**: Easily deploy and manage containers with `docker-compose`.
- **Trivia Management**: Create and retrieve trivia facts (questions and answers).
- **Environment Variables**: Secure and configurable database credentials using `.env` file.

## Technologies

- **Programming Language**: Go (Golang)
- **Web Framework**: Go Fiber
- **Database**: PostgreSQL
- **Containerization**: Docker, Docker Compose
- **Tools**: Postman for API testing

## Installation

### Prerequisites

- Docker & Docker Compose installed on your system.
- Go installed (for local development without Docker).

### Steps

1. **Clone the Repository**:
   
   ```bash
   git clone https://github.com/tunde262/trivia_rest_api_docker_postgres_go.git
   cd trivia_rest_api_docker_postgres_go

2. **Create an `.env` File**: Define the following environment variables in a `.env` file:
   
   ```bash
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name

3. **Start Services**: Use Docker Compose to build and start the services:
   
   ```bash
   docker-compose up --build

4. **Access the API:y**:
   
- The API will run on `http://localhost:3000`.
- Use tools like Postman or `curl` to interact with the endpoints.

## API Endpoints
- **GET /**: List all trivia facts.
- **POST /fact**: Create a new trivia fact.

## Example Payload for `POST /fact`

  ```bash
  {
    "question": "What is the capital of France?",
    "answer": "Paris"
  }
  ```

## How It Works

1. Database Connection:
   
  - A PostgreSQL database is initialized using credentials from the .env file.
  - GORM is used for migrations and ORM capabilities.

2. REST API:
   
  - The `ListFacts` handler retrieves all trivia facts from the database.
  - The `CreateFact` handler allows users to add new trivia facts.

3. Docker:

  - `web` service runs the Go Fiber app.
  - `db` service sets up a PostgreSQL database.
  - Both services are connected through Docker Compose.

## Project Structure

  ```bash
  trivia_rest_api_docker_postgres_go/
  ├── cmd/
  │   └── main.go         # Entry point of the application
  │   └── routes.go       # API route definitions
  ├── database/
  │   └── database.go     # Database connection and migrations
  ├── handlers/
  │   ├── facts.go        # Handlers for API routes
  ├── models/
  │   └── models.go       # Data models (e.g., Fact)
  ├── docker-compose.yml  # Docker Compose configuration
  ├── Dockerfile          # Docker image setup
  └── .env.example        # Example .env file
```

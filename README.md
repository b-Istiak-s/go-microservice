# Microservices Project

## Purpose Of The Repository

This repository demonstrates a **modular microservices architecture** using Go (Golang), Docker, and Kubernetes. It follows clean architecture principles and applies **DDD-inspired concepts** where relevant. 
A video of API test is provided [here](static/go-microservice-insomnia-screenrecord.webm).

It contains two independent services:

- **Auth Service (`auth-service`)**  
  Handles user registration, login, and authorization. Uses **PostgreSQL**.

- **CRUD Service (`crud-service`)**  
  Manages notes (Create, Read, Update, Delete). Uses **MySQL**.  
  Also interacts with `auth-service` to verify users.

Both services are containerized with Docker and orchestrated using Kubernetes. Each service runs two replicas and connects to its respective local database.

---

## Features

- Clean separation of services with dedicated databases.
- Follows clean architecture and DDD-inspired structuring.
- REST-based HTTP APIs.
- JWT authentication and route-level authorization.
- Structured request validation and response formatting.
- Kubernetes support with service and ingress routing.
- Custom middleware and error handling.

---

## Tech Stack

| Component         | Technology    |
|------------------|---------------|
| Language         | Go (Golang)   |
| API Protocol     | REST (HTTP)   |
| Auth DB          | PostgreSQL    |
| Notes DB         | MySQL         |
| Containers       | Docker        |
| Orchestration    | Kubernetes    |

---

## Services Overview

### 1. Auth Service (`auth-service`)

- **Purpose:** User authentication and token handling.
- **Stack:** Go, Gin, PostgreSQL
- **Endpoints:**
  - `POST /register` – Register new users
  - `POST /login` – User login and JWT generation
  - `POST /verify` – Internal verification (used by `crud-service`)

### 2. CRUD Service (`crud-service`)

- **Purpose:** CRUD operations for user notes.
- **Stack:** Go, Gin, MySQL
- **Endpoints:**
  - `POST /notes` – Create a note
  - `GET /notes` – List all notes
  - `PUT /notes/:id` – Update a note
  - `DELETE /notes/:id` – Delete a note

---

## Kubernetes Configuration

### Files

- `auth-service.yaml` – Deployment and service for Auth
- `crud-service.yaml` – Deployment and service for CRUD
- `secret.yaml` – Secrets for DB credentials
- `ingress.yaml` – Ingress configuration with path-based routing

### Ingress Routing

- `/api/auth/...` → routes to `auth-service` (port 8080)
- `/api/notes/...` → routes to `crud-service` (port 8081)

---

## Repository Structure

<details>
<summary>Click to expand</summary>

```plaintext
.
├── auth-service
│   ├── cmd
│   │   └── myapp
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── auth
│   │   │   └── auth.go
│   │   ├── controller
│   │   │   └── auth
│   │   │       └── auth_controller.go
│   │   ├── db
│   │   │   └── postgresql.go
│   │   ├── middleware
│   │   │   └── authentication.go
│   │   ├── model
│   │   │   └── user.go
│   │   ├── repository
│   │   │   └── auth
│   │   │       └── user_repository.go
│   │   ├── route
│   │   │   └── auth_route.go
│   │   ├── service
│   │   │   └── auth
│   │   │       └── auth_service.go
│   │   ├── util
│   │   │   └── response
│   │   │       └── response.go
│   │   └── validator
│   │       └── auth
│   │           ├── login_request.go
│   │           └── register_request.go
│   └── tmp
│       └── build-errors.log
├── crud-service
│   ├── cmd
│   │   └── myapp
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── auth
│       │   └── auth.go
│       ├── controller
│       │   └── note_controller.go
│       ├── db
│       │   └── mysql.go
│       ├── middleware
│       │   └── authentication.go
│       ├── model
│       │   └── note.go
│       ├── repository
│       │   └── note_repository.go
│       ├── route
│       │   └── note_route.go
│       ├── util
│       │   └── response
│       │       └── response.go
│       └── validator
│           └── note
│               ├── create_request.go
│               ├── update_request.go
│               └── validate_note.go
├── k8s-configs
│   ├── auth-service.yaml
│   ├── crud-service.yaml
│   ├── ingress.yaml
│   └── db-secret.example.yaml
├── LICENSE
├── README.md
└── static
    ├── go-microservice-insomnia-screenrecord.webm
    └── vs-code-go-profile.code-profile
```

</details>

---

## Deployment Notes

- Services use external databases accessible via `192.168.0.130`.
- Each service has 2 replicas for availability.
- All sensitive data is injected using Kubernetes secrets.
- LoadBalancer or Ingress exposes services through a unified API endpoint.
- `crud-service` internally calls `auth-service` to validate user identity.

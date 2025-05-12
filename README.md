# Microservices Project

## Purpose Of The Repository

This project demonstrates a **simple, modular microservices architecture** using Go (Golang), Docker, and Kubernetes, following clean architecture and applying **Domain-Driven Design (DDD) principles**.

It consists of two independent services:

- **Auth Service (`auth-service`)**  
  Responsible for user authentication and authorization. Uses **PostgreSQL** as its persistence layer.

- **CRUD Service (`crud-service`)**  
  Responsible for managing notes with basic Create, Read, Update, and Delete (CRUD) operations. Uses **MySQL** as its persistence layer.

Both services are containerized using Docker and orchestrated via Kubernetes configurations.

---

## Features

- Clean microservice separation with dedicated databases.
- Implementation of clean architecture and DDD principles.
- HTTP API endpoints exposed via REST.
- JWT-based authentication and authorization.
- Validation layer for request/response.
- Kubernetes-ready with YAML configurations.
- Custom middlewares, error handling, and response formatting.

---

## Tech Stack

| Component         | Technology    |
|--------------------|--------------|
| Language            | Go (Golang)  |
| API Protocol        | REST (HTTP)  |
| Auth Service DB     | PostgreSQL   |
| CRUD Service DB     | MySQL        |
| Containers          | Docker       |
| Orchestration       | Kubernetes   |

---

## Services Overview

### 1. Auth Service (`auth-service`)

- **Purpose:** Handles user registration, login, and JWT issuance.
- **Tech Stack:** Go, Gin, PostgreSQL
- **Endpoints:**
  - `POST /register`
  - `POST /login`
  - `POST /verify` (Protected) (crud-service calls the route to verify the logged in user exists)

### 2. CRUD Service (`crud-service`)

- **Purpose:** Manages user notes.
- **Tech Stack:** Go, Gin, MySQL
- **Endpoints:**
  - `POST /notes`
  - `GET /notes`
  - `PUT /notes/:id`
  - `DELETE /notes/:id`

---

## Kubernetes Configurations

- `auth-service.yaml`
- `crud-service.yaml`
- `mysql.yaml`
- `postgres.yaml`
- `ingress.yaml`
- `secret.yaml`

---

## Repository Structure

<details>
<summary>Structure</summary>

```plaintext
.
├── auth-service
│   ├── cmd
│   │   └── myapp
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── auth
│   │   │   └── auth.go
│   │   ├── controller
│   │   │   └── auth
│   │   │       └── auth_controller.go
│   │   ├── db
│   │   │   └── postgresql.go
│   │   ├── middleware
│   │   │   └── authentication.go
│   │   ├── model
│   │   │   └── user.go
│   │   ├── repository
│   │   │   └── auth
│   │   │       └── user_repository.go
│   │   ├── route
│   │   │   └── auth_route.go
│   │   ├── service
│   │   │   └── auth
│   │   │       └── auth_service.go
│   │   ├── util
│   │   │   └── response
│   │   │       └── response.go
│   │   └── validator
│   │       └── auth
│   │           ├── login_request.go
│   │           └── register_request.go
│   └── tmp
│       └── build-errors.log
├── crud-service
│   ├── cmd
│   │   └── myapp
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── auth
│       │   └── auth.go
│       ├── controller
│       │   └── note_controller.go
│       ├── db
│       │   └── mysql.go
│       ├── middleware
│       │   └── authentication.go
│       ├── model
│       │   └── note.go
│       ├── repository
│       │   └── note_repository.go
│       ├── route
│       │   └── note_route.go
│       ├── util
│       │   └── response
│       │       └── response.go
│       └── validator
│           └── note
│               ├── create_request.go
│               ├── update_request.go
│               └── validate_note.go
├── k8s-configs
│   ├── auth-service.yaml
│   ├── crud-service.yaml
│   ├── ingress.yaml
│   ├── mysql.yaml
│   ├── postgres.yaml
│   └── secret.yaml
└── README.md

37 directories, 38 files
```
</details>


## Issues

1. **Kubernetes Deployment Issues**
   - Both services (`auth-service` and `crud-service`) run successfully in local Docker environments.
   - However, when deployed in Kubernetes, both services face **database connection issues**.
   - Possible reasons:
     - Database services (`PostgreSQL` and `MySQL`) are not fully initialized when the microservices try to connect.
     - Missing or incorrect Kubernetes `Service`, `Secret`, or `ConfigMap` configurations for DB connections.
     - Network policies or incorrect `Service` discovery within the cluster.
   - **Next Steps to fix:**
     - Ensure proper `initContainers` or startup probes to verify DB readiness before service starts.
     - Validate connection strings and ensure secrets are properly mounted.
     - Check Kubernetes logs (`kubectl logs`) and events (`kubectl describe`) to pinpoint the issue.
     - Use Kubernetes `liveness` and `readiness` probes effectively.

2. **Secrets & Environment Configuration**
   - Missing or incorrect environment variables (DB host, port, username, password) could cause connection failures.
   - Ensure all sensitive data is securely configured using Kubernetes `Secrets` and not hardcoded.

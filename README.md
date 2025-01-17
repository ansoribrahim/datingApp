# Dating App Backend

This is a backend implementation for a dating app similar to Tinder or Bumble. It provides core features like user sign-up, login, swiping, and purchasing premium packages. The backend is developed in **Golang** and uses **PostgreSQL** for data storage and **Redis** for caching. It is containerized using Docker Compose for easy setup.

---

## Features

### Core Features:
1. **User Authentication**:
    - Sign-up
    - Login
2. **Swiping**:
    - Swipe left (pass)
    - Swipe right (like)
    - Limit of 10 swipes per day for non-premium users
3. **Premium Features**:
    - Remove swipe quota
    - Add a "Verified" label to user profiles

---

## Tech Stack

- **Golang**: Backend development
- **PostgreSQL**: Relational database
- **Redis**: In-memory caching
- **Docker Compose**: Container orchestration
- **Gin**: HTTP web framework for Golang

---

## Prerequisites

- Docker and Docker Compose installed
- Golang installed (1.19 or later)

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/dating-app-backend.git
cd dating-app-backend
```

### 2. Set Up Infrastructure with Docker Compose
Start the database and Redis services:
```bash
docker-compose up -d
```

### 3. Run Database Migrations
Ensure the database schema is created:
````bash
docker exec -i dating_app_db psql -U user -d dating_app < db/migrations.sql
````

### 4. Run the Application
Start the backend server:
```bash
go run main.go
```
The server will run on http://localhost:8080.
---

License
This project is licensed under the MIT License.


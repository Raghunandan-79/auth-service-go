# Auth Service (Go + Gin + JWT + Refresh Tokens)

A **production-inspired authentication service** built using **Go**, **Gin**, **GORM**, **MySQL**, **JWT access tokens**, and **database-backed refresh tokens**.

This project is **not a toy tutorial** — it demonstrates **real-world authentication concepts**, correct database constraints, and secure session handling.

---

## Why this Auth Service is Different

Most authentication tutorials:
- Store everything in memory
- Use only JWT (no logout support)
- Don’t handle refresh tokens correctly
- Allow duplicate users due to missing DB constraints
- Ignore real database behavior

### This project does it RIGHT:
- **JWT Access Tokens** (short-lived)
- **Refresh Tokens stored in DB** (logout supported)
- **Database-level UNIQUE constraint on email**
- **bcrypt password hashing**
- **Session-based refresh token model**
- **No duplicate users possible**
- Designed to be tested with Postman / curl
- Clean project structure (controllers, models, middleware, utils)

This mirrors how **real backend systems** handle authentication.

---

## Authentication Flow
```bash
Register -> Login -> Access Protected Routes
-> Refresh Token -> New Access Token
Logout -> Refresh Token Deleted
```

## Project Structure
```bash
auth-service/
├── controllers/ # Business logic (register, login, refresh, logout)
├── database/ # Database connection
├── middleware/ # JWT authentication middleware
├── models/ # GORM models (User, RefreshToken)
├── routes/ # API routes
├── utils/ # JWT & token utilities
├── main.go # Application entry point
├── .env # Environment variables
├── README.md
└── LICENSE
```

---

## Tech Stack

- **Go**
- **Gin** – HTTP framework
- **GORM** – ORM
- **MySQL**
- **JWT** – Access tokens
- **bcrypt** – Password hashing

---

## Environment Setup

### 1 Install Dependencies

```bash
go mod tidy
```

### 2 Create .env File

```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=auth_service

JWT_SECRET=your_super_secret_key
ACCESS_TOKEN_MINUTES=15
REFRESH_TOKEN_DAYS=30
```

### 3 Create Database

```sql
CREATE DATABASE auth_service;
```


## Running the Server

```bash
go run main.go
```

- You should see:

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

2026/01/30 17:27:50 MySQL connected
[GIN-debug] POST   /auth/api/v1/register     --> github.com/Raghunandan-79/auth-service/controllers.Register (3 handlers)
[GIN-debug] POST   /auth/api/v1/login        --> github.com/Raghunandan-79/auth-service/controllers.Login (3 handlers)
[GIN-debug] POST   /auth/api/v1/refresh      --> github.com/Raghunandan-79/auth-service/controllers.Refresh (3 handlers)
[GIN-debug] POST   /auth/api/v1/logout       --> github.com/Raghunandan-79/auth-service/controllers.Logout (3 handlers)
[GIN-debug] GET    /auth/api/v1/me           --> github.com/Raghunandan-79/auth-service/controllers.Me (4 handlers)
2026/01/30 17:27:50 Server started on localhost:8082
```

## API Endpoints

- Base URL:
```bash
http://localhost:8082
```

### Register (one-time per email)

- POST

```bash
/auth/api/v1/register
```
- Body

```json
{
  "name": "yourname",
  "email": "youremail@example.com",
  "password": "yourpassword"
}
```

- Response

```json
{
  "message": "registered"
}
```

### Login

- POST

```bash
/auth/api/v1/login
```

- Body

```json
{
  "email": "youremail@example.com",
  "password": "yourpassword"
}
```

- Response

```json
{
  "access_token": "JWT_ACCESS_TOKEN"
}
```

>Note: A refresh token is automatically stored as an HttpOnly cookie.

### Protected Route

- GET

```bash
/auth/api/v1/me
```

- Header

```bash
Authorization: Bearer <ACCESS_TOKEN>
```

- Response

```json
{
  "user_id": 1
}
```

### Refresh Access Token

- POST

```bash
/auth/api/v1/refresh
```

- No body
- Refresh token sent via cookie automatically

- Response

```json
{
  "access_token": "NEW_ACCESS_TOKEN"
}
```


### Logout

- POST

```bash
/auth/api/v1/logout
```

- Deletes refresh token from DB
- Clears cookie

- Response

```json
{
  "message": "logged out"
}
```

## Security Features 
- Passwords hashed with bcrypt
- JWT signed with secret key
- Refresh tokens stored in database
- Logout supported (session invalidation)
- Database-level UNIQUE email constraint
- No in-memory auth tricks

## Notes
- This project is secure for learning and internal use
- For production, consider:
    - Hashing refresh tokens
    - Refresh-token rotation
    - Rate limiting
    - HTTPS + Secure cookies
    - Database migrations (golang-migrate)

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author
Built by **Raghunandan Sharma**. Learning focused, backend-first, and security aware

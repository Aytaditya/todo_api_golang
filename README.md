
# Todo & Notes API 

A **RESTful API** built in **Go** that allows users to create, read, update, and delete **Todos/Notes**. The project uses **SQLite** for storage and implements **JWT-based authentication** for secure, protected routes.  

---

## Features

- User registration with **hashed passwords** (bcrypt)
- User login with **JWT token issuance**
- Protected routes using **JWT middleware**
- CRUD operations for **Todos / Notes**
- SQLite database integration
- Clean and modular Go code structure

---

## Tech Stack

- **Backend:** Go (net/http)
- **Database:** SQLite
- **Authentication:** JWT (github.com/golang-jwt/jwt/v5)
- **Password Hashing:** bcrypt (golang.org/x/crypto/bcrypt)

---

## API Endpoints

### 1. Register User

```

POST /register
Body: {
"username": "yourname",
"email": "[you@example.com](mailto:you@example.com)",
"password": "yourpassword"
}
Response: {
"message": "User created successfully",
"user_id": 1,
"token": "<jwt_token>"
}

```

---

### 2. Login User

```

POST /login
Body: {
"email": "[you@example.com](mailto:you@example.com)",
"password": "yourpassword"
}
Response: {
"message": "User logged in successfully",
"user_id": 1,
"token": "<jwt_token>"
}

```

---

### 3. Create Todo (Protected)

```

POST /todos
Headers: Authorization: Bearer <jwt_token>
Body: {
"title": "My Todo",
"content": "Todo details",
"tag": "Todo-tag"
}
Response: {
"message": "Todo created successfully",
"todo_id": 1
}

```

---

### 4. Get Todos (Protected)

```

GET /todos
Headers: Authorization: Bearer <jwt_token>
Response: [
{
"id": 1,
"title": "My Todo",
"content": "Todo details",
"tag": "todo-tag",
"user_id": 1
},
...
]

```

---

### 5. Delete Todo (Protected)

```

DELETE /todos/{id}
Headers: Authorization: Bearer <jwt_token>
Response: {
"message": "Todo deleted successfully"
}

````

---

## How it Works

1. **Registration**
   - User sends username, email, password.
   - Password is **hashed using bcrypt** before storing.
   - JWT token is issued after successful registration.

2. **Login**
   - User sends email and password.
   - Password is verified using bcrypt.
   - JWT token is issued for authenticated requests.

3. **Protected Routes**
   - JWT token is sent in the **Authorization header**.
   - Middleware extracts and validates the token.
   - If valid, request proceeds; otherwise, 401 Unauthorized is returned.

4. **Database**
   - SQLite is used for simplicity.
   - `users` table stores user info and hashed passwords.
   - `todos/notes` table stores todos linked to users.

---

### **Config & Environment**

This project uses a `config/local.yaml` file to store configuration values such as:

```yaml
env: "local"
storage_path: "./data/sqlite.db"
http_server:
  addr: ":8080"
jwt_secret: "SUPER_SECRET_KEY_CHANGE_THIS"
```

* `env` → environment type (local, dev, prod)
* `storage_path` → path to SQLite database file
* `http_server.addr` → server address and port
* `jwt_secret` → secret key used to sign JWT tokens

The project uses the [cleanenv](https://github.com/ilyakaznacheev/cleanenv) Go package to **load and validate configuration** at runtime.

---

### **Setup & Run (with config)**

1. Make sure `config/local.yaml` exists and contains all required keys.
2. Clone the repository:

```bash
git clone <your-repo-url>
cd project-root
```

3. Install dependencies:

```bash
go mod tidy
```

4. Run the server using the configuration:

```bash
go run main.go --config config/local.yaml
```

* The server will read all required settings from `local.yaml`
* JWT secret, database path, and server port will be loaded automatically

---

### ⚡ Tip:

For **production**, you can override values using **environment variables**. For example:

```bash
export ENV=prod
export JWT_SECRET="SUPER_SECURE_KEY"
go run main.go --config config/prod.yaml
```

---


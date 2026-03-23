# Chautari

A real-time chat application API built with **Go**, **Chi Router**, **Websockets**, and **PostgreSQL**.

---

## Environment Variables

| Variable       | Description                       | Example                       |
| -------------- | --------------------------------- | ----------------------------- |
| `ENVIRONMENT`  | App environment                   | `development` or `production` |
| `PORT`         | Server port                       | `8080`                        |
| `DATABASE_URL` | Neon PostgreSQL connection string | `postgresql://...`            |
| `JWT_SECRET`   | Secret key for signing JWT tokens | `your-secret-key`             |

---

## API Endpoints

### Base URL

```
.../chautari/api/v1
```

## Endpoints

### POST `/register`

**Request:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "Password123!"
}
```

```
Password Validation
- Minimum 8 characters
- At least one letter
- At least one number
- At least one special character

```

**Response `201 Created`:**

```json
{
  "success": true,
  "message": "User registered successfully"
}
```

**Error Responses:**

| Status | Message                                   |
| ------ | ----------------------------------------- |
| `400`  | Invalid JSON                              |
| `400`  | Name cannot be empty                      |
| `400`  | Invalid email                             |
| `400`  | Password must be at least 8 characters... |
| `409`  | Email already in use                      |
| `409`  | Error generating unique username          |
| `500`  | Something went wrong                      |

---

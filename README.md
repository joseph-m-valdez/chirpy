# Chirpy

Chirpy is a lightweight HTTP server written in Go that provides authenticated user management and microblog-style messaging ("chirps").
It uses PostgreSQL for data storage and offers RESTful API endpoints for user and chirp operations, along with webhook handling and basic server metrics.

## Features

- User account creation and authentication with JWT support
- CRUD operations for "chirps" (short messages)
- Token refresh and revocation endpoints
- Webhook endpoint integration (e.g., Polka webhooks)
- File serving with metrics tracking for requests
- Health check endpoint for server status
- Admin endpoints for metrics and reset operations

## Environment Variables

The application expects the following environment variables:

- `DB_URL`: PostgreSQL connection string (required)
- `JWT_SEKRET`: Secret key for JWT token generation
- `PLATFORM`: Platform identifier (optional)
- `POLKA_KEY`: Key for webhook authentication (optional)

Ensure a `.env` file or environment setup provides these variables before running the server.

## Building and Running

1. Clone the repository.
2. Set up your PostgreSQL database and update the `DB_URL`.
3. Create a `.env` file with the necessary environment variables.
4. Build and run the application:
```bash
go build -o chirpy
./chirpy
```


The server listens on port `8080` by default.

## API Endpoints

### User Endpoints

- `POST /api/users`  
  Create a new user.

- `PUT /api/users`  
  Update user authorization details.

- `POST /api/login`  
  User login to retrieve JWT token.

- `POST /api/refresh`  
  Refresh JWT token.

- `POST /api/revoke`  
  Revoke JWT token.

### Chirp Endpoints

- `POST /api/chirps`  
  Create a new chirp.

- `GET /api/chirps`  
  Get a list of chirps.

- `GET /api/chirps/{chirpID}`  
  Get details of a single chirp.

- `DELETE /api/chirps/{chirpID}`  
  Delete a chirp.

### Webhook Endpoint

- `POST /api/polka/webhooks`  
  Handle Polka webhook callbacks.

### Admin Endpoints

- `GET /admin/metrics`  
  Retrieve server metrics.

- `POST /admin/reset`  
  Reset server state or metrics.

### Miscellaneous

- `GET /api/healthz`  
  Health check for the server.

## Project Structure

- `main.go`: Server setup and HTTP routing.
- `internal/api`: API handlers and middleware.
- `internal/database`: Database query implementations.
- `internal/config`: Configuration structures.


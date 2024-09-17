# go-rate-limiter
Rate limiter implementation in GO.

## Overview

This Go application implements middleware for rate limiting using the Token Bucket algorithm. It supports varying rate limits for different users and endpoints, manages concurrent requests efficiently, and includes basic logging and metrics for monitoring.

## Running the Service

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dev-pramodThapa/go-rate-limiter.git
   ```
   ```bash
   cd go-rate-limiter
   ```

2. **Build the application:**

   ```bash
   go build -o go-rate-limiter
   ```

3. **Run the server:**

   ```bash
   ./go-rate-limiter
   ```

4. **Access the endpoints:**
   - `GET /user/:id/data`
   - `GET /admin/:id/dashboard`
   - `GET /public/info`
   - `GET /metrics`
   - `POST /update-rate-limit`

     - Sample request body:
     ```json
     "user_type": "admin",
     "id": "1",
     "max_tokens": 5,
     "refill_rate": 10
     ```

## Testing the System

To run the test follow the steps:

```bash
	cd tests
```
```bash
   	go test ./
```
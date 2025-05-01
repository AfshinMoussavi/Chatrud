## Running the Project with Docker

This project provides a ready-to-use Docker setup for local development and testing. The configuration uses Docker Compose to orchestrate the Go application and a PostgreSQL database.

### Project-Specific Requirements
- **Go Version:** The Dockerfile builds the app using Go `1.24.2-alpine`.
- **PostgreSQL:** Uses the official `postgres:latest` image.

### Environment Variables
- The PostgreSQL service is configured with the following environment variables (set in `docker-compose.yml`):
  - `POSTGRES_USER=postgres`
  - `POSTGRES_PASSWORD=postgres`
  - `POSTGRES_DB=appdb`
- No additional environment variables are required for the Go app by default. If you need to provide custom environment variables, you can uncomment and use the `env_file` section in the `docker-compose.yml`.

### Build and Run Instructions
1. **Build and start all services:**
   ```sh
   docker compose up --build
   ```
   This will build the Go application and start both the app and the PostgreSQL database.

2. **Stopping services:**
   ```sh
   docker compose down
   ```

### Configuration
- The Go application uses configuration files from the `/app/config` directory. By default, `config-docker.yml` is included in the image.
- Logs are stored in `/app/logs` inside the container. To persist logs on the host, uncomment the `volumes` section for `go-app` in `docker-compose.yml`.
- The application runs as a non-root user for improved security.

### Exposed Ports
- **Go Application:**
  - Host: `8080` → Container: `8080`
- **PostgreSQL:**
  - Host: `5432` → Container: `5432`

### Additional Notes
- The Go app depends on the PostgreSQL service and will wait for it to be healthy before starting.
- The Docker Compose network is named `app-network` (bridge driver).
- Database data is persisted in a Docker volume named `pgdata`.

---

_If you need to customize configuration or environment variables, edit the `config-docker.yml` or provide a `.env` file as needed._
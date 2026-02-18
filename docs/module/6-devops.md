# 6. DevOps & Infrastructure

Our infrastructure is designed for 99.9% uptime and easy scalability.

## Containerization
The entire stack is Dockerized. We use a multi-container setup orchestrated by Docker Compose.

* **API (Go)**: Compiled in a lightweight alpine image.
* **Web (Next.js)**: Optimized production build served via Nginx.
* **Reverse Proxy (Nginx)**: Handles SSL termination and load balancing.

## CI/CD Pipeline (GitHub Actions)
Every code push undergoes rigorous checks:
1. **Linting**: Ensures code quality and consistency.
2. **Unit Testing**: 
   - `go test ./...` for backend logic.
   - `npm test` for frontend components.
3. **Automated Deployment**: Merges to `main` trigger a webhook to the production server.

## Monitoring
We use **Prometheus** to track real-time analytics:
- Request latency.
- Error rates (4xx and 5xx).
- Active database connections.

---

### Deployment Command
```bash
# To deploy the entire system
./deploy.sh
```

> [!NOTE]
> The `deploy.sh` script automatically handles database migrations and container restarts with zero downtime (blue-green transition where possible).

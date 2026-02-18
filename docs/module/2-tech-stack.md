# 2. Technical Stack

Our technology choices are driven by performance, type safety, and developer productivity.

## Frontend
- **Framework**: [Next.js 15 (App Router)](https://nextjs.org/)
- **Logic**: React 18 with Custom Hooks & Context API
- **Data Fetching**: [SWR](https://swr.vercel.app/) for stale-while-revalidate caching.
- **Styling**: Tailwind CSS & Shadcn UI
- **Animations**: Framer Motion

## Backend
- **Language**: [Go 1.24](https://go.dev/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/) with PostgreSQL
- **Security**: JWT with HttpOnly Cookies & Google OAuth2

## Infrastructure
- **Media**: [Cloudinary](https://cloudinary.com/) for cloud-based image management.
- **Monitoring**: Prometheus & Gin Metrics
- **Runtime**: Docker & Nginx Reverse Proxy
- **CI/CD**: GitHub Actions

---

### Why this stack?
| Choice | Reason |
| :--- | :--- |
| **Go** | Extremely fast execution and low memory footprint for high-concurrency booking tasks. |
| **SWR** | Provides an "instant" feel UI by caching remote data locally. |
| **HttpOnly Cookies** | Mitigates XSS attacks by keeping tokens invisible to JavaScript. |
| **Cloudinary** | Offloads image processing and storage, keeping the server lightweight. |

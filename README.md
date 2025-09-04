# Backend - Unit Management API

## 🚀 Features
- RESTful API with endpoints:
  - `GET /api/v1/units` → Fetch all units (supports `?status=Available` filter).
  - `POST /api/v1/units` → Create a new unit.
  - `PUT /api/v1/units/:id` → Update unit status.
  - `GET /api/v1/units/:id` → Get details of a single unit.
- PostgreSQL database integration.
- Business rule validation:
  - A unit cannot change directly from `Occupied → Available`.
  - Must transition via `Cleaning In Progress` or `Maintenance Needed`.

---

## 🛠️ Tech Stack
- Golang (`net/http`, `github.com/go-chi/chi`)
- PostgreSQL
- `github.com/jmoiron/sqlx` for DB access
- Migration handled with **Makefile**

---

## ⚙️ Setup & Installation

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- Make

### Steps
```bash
# Clone repository
git clone https://github.com/your-repo.git
cd your-repo/backend

# Copy env
cp .env.example .env
# Update .env with your DB connection

# Run migration
make migrate-up

# Start server
make run

# cms-be

✨ Backend service for the new portfolio CMS.

Built with **Go**, **Fiber**, **Gorm**, and **PostgreSQL**, this repository is designed to become the API backbone for two separate clients:

- 🌐 `cms-fe` — public-facing portfolio frontend
- 🛠️ `cms-adm` — admin panel for managing content

The backend owns the data model and exposes HTTP APIs. Both clients should communicate with this service only — **no direct database access**.

---

## 🚀 Stack

- **Go**
- **Fiber**
- **Gorm**
- **PostgreSQL**
- **Clean Architecture-inspired structure**

---

## 📌 Current Status

This project is in active early-stage development.

### ✅ Already available
- application bootstrap
- environment-based configuration
- PostgreSQL connection setup
- automatic entity migration on startup
- real healthcheck endpoint with database ping
- initial domain entities
- initial `Profile` foundation *(work in progress)*
- local uploads directory foundation

### 🧩 In progress
- repository/usecase layering completion
- admin authentication
- full CRUD endpoints for content modules
- validation, authorization, and production hardening

---

## 🔗 API Base Path

```text
/api/v1
```

### ❤️ Healthcheck

```http
GET /api/v1/health
```

Example response:

```json
{
  "message": "ok",
  "data": {
    "status": "healthy",
    "database": "up"
  }
}
```

---

## 🗂️ Project Structure

```text
cms-be/
├── cmd/
│   └── api/                  # application entrypoint
├── internal/
│   ├── app/                  # bootstrap/container
│   ├── config/               # environment config loader
│   ├── delivery/http/        # Fiber app, routes, handlers
│   ├── domain/entity/        # core domain entities
│   ├── domain/repository/    # repository contracts
│   ├── infrastructure/       # external infra adapters
│   ├── repository/postgres/  # PostgreSQL repository impls
│   └── usecase/              # business usecases
├── migrations/               # database migration assets
├── pkg/response/             # shared HTTP response helpers
└── storage/uploads/          # local file upload storage
```

---

## ⚙️ Environment Variables

Copy the example file first:

```bash
cp .env.example .env
```

Main configuration:

```env
APP_HOST=127.0.0.1
APP_PORT=8080

DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=cms_be_user
DB_PASSWORD=your_password
DB_NAME=cms_be
DB_SSLMODE=disable

JWT_SECRET=your_secret_key
UPLOAD_DIR=storage/uploads
```

---

## ▶️ Running Locally

Make sure PostgreSQL is running and the configured database is accessible.

Then start the API:

```bash
go run ./cmd/api
```

If everything is healthy, the service should be available at:

```text
http://127.0.0.1:8080
```

---

## 🧱 Initial Domain Modules

Planned core modules:
- 👤 `Profile`
- 💼 `Project`
- 🧠 `Skill`
- 🧾 `Experience`
- 🔗 `SocialLink`
- 🔐 `AdminUser`

These modules are intended to support a single portfolio owner with content managed through the admin panel.

---

## 📝 Development Notes

- This repository currently uses `AutoMigrate` during bootstrap for early development.
- Public and admin clients will communicate with this backend through HTTP APIs.
- Uploads currently target local storage first; external object storage can be added later.

---

## 🛣️ Roadmap

Short-term priorities:
- finish `Profile` endpoints
- add proper request validation
- implement admin authentication
- complete CRUD for portfolio modules
- prepare stable API contracts for `cms-fe` and `cms-adm`

---

## 🔒 License

Private project for internal development.

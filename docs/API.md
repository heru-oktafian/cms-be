# CMS BE API Reference

Base URL (local):

```text
http://127.0.0.1:8080/api/v1
```

## Conventions

- Content-Type for JSON requests: `application/json`
- Admin endpoints may require:

```http
Authorization: Bearer <JWT_TOKEN>
```

---

## Health

### GET `/health`
Check backend and database status.

**Response**
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

## Admin Auth

### POST `/admin/auth/login`
Login using admin email and password.

**Request body**
```json
{
  "email": "admin@cms.local",
  "password": "your_password"
}
```

**Success response**
```json
{
  "message": "login success",
  "data": {
    "token": "JWT_TOKEN",
    "user": {
      "id": 1,
      "name": "Admin",
      "email": "admin@cms.local",
      "is_active": true,
      "created_at": "2026-04-29T00:00:00Z",
      "updated_at": "2026-04-29T00:00:00Z"
    }
  }
}
```

### POST `/admin/auth/logout`
Logical logout endpoint.

**Headers**
```http
Authorization: Bearer <JWT_TOKEN>
```

**Success response**
```json
{
  "message": "logout success",
  "data": null
}
```

---

## Profile

### GET `/public/profile`
Fetch public profile.

### GET `/admin/profile`
Fetch admin profile.

### PUT `/admin/profile`
Create or update profile.

**Request body**
```json
{
  "full_name": "Heru Oktafian",
  "headline": "Backend Engineer",
  "sub_headline": "Building clean systems",
  "bio": "Sample profile",
  "email": "hello@example.com",
  "phone": "08123456789",
  "location": "Indonesia",
  "avatar_path": "/uploads/avatar.png",
  "resume_path": "/uploads/resume.pdf"
}
```

**Required fields**
- `full_name`

---

## Projects

### GET `/public/projects`
List public projects.

### GET `/admin/projects`
List admin projects.

### GET `/admin/projects/:id`
Get project by id.

### POST `/admin/projects`
Create project.

### PUT `/admin/projects/:id`
Update project.

### DELETE `/admin/projects/:id`
Delete project.

**Request body**
```json
{
  "title": "CMS Portfolio",
  "slug": "cms-portfolio",
  "summary": "Portfolio CMS project",
  "description": "Backend API for portfolio CMS",
  "thumbnail_path": "/uploads/project.png",
  "project_url": "https://example.com",
  "repo_url": "https://github.com/heru-oktafian/cms-be",
  "is_featured": true,
  "sort_order": 1
}
```

**Required fields**
- `title`
- `slug`

---

## Skills

### GET `/public/skills`
List public skills.

### GET `/admin/skills`
List admin skills.

### GET `/admin/skills/:id`
Get skill by id.

### POST `/admin/skills`
Create skill.

### PUT `/admin/skills/:id`
Update skill.

### DELETE `/admin/skills/:id`
Delete skill.

**Request body**
```json
{
  "name": "Go",
  "level": "Advanced",
  "icon_path": "/uploads/go.png",
  "sort_order": 1
}
```

**Required fields**
- `name`

---

## Experiences

### GET `/public/experiences`
List public experiences.

### GET `/admin/experiences`
List admin experiences.

### GET `/admin/experiences/:id`
Get experience by id.

### POST `/admin/experiences`
Create experience.

### PUT `/admin/experiences/:id`
Update experience.

### DELETE `/admin/experiences/:id`
Delete experience.

**Request body**
```json
{
  "company": "PT Example",
  "position": "Backend Engineer",
  "description": "Handle backend architecture and API development",
  "start_date": "2023-01",
  "end_date": "2024-12",
  "is_current": false,
  "sort_order": 1
}
```

**Required fields**
- `company`
- `position`

---

## Social Links

### GET `/public/social-links`
List public social links.

### GET `/admin/social-links`
List admin social links.

### GET `/admin/social-links/:id`
Get social link by id.

### POST `/admin/social-links`
Create social link.

### PUT `/admin/social-links/:id`
Update social link.

### DELETE `/admin/social-links/:id`
Delete social link.

**Request body**
```json
{
  "platform": "GitHub",
  "label": "heru-oktafian",
  "url": "https://github.com/heru-oktafian",
  "icon_path": "/uploads/github.png",
  "sort_order": 1
}
```

**Required fields**
- `platform`
- `url`

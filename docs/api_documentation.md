# Portfolio REST API Documentation

This document describes the REST API endpoints provided by the backend server built with Golang Fiber, GORM, and PostgreSQL.

* **Base URL:** `http://localhost:3000/api/v1`
* **Content-Type:** `application/json`

---

## Response Envelope Structure

All endpoints return responses wrapped in a consistent JSON structure:

### Success Response
```json
{
  "success": true,
  "message": "Retrieval successful",
  "data": ...
}
```

### Error Response
```json
{
  "success": false,
  "message": "Description of error",
  "error": "Detailed error string (optional)"
}
```

---

## Authentication

Protected endpoints require a JSON Web Token (JWT) provided in the `Authorization` request header:

```http
Authorization: Bearer <JWT_TOKEN>
```

---

## 1. Authentication Endpoints

### Admin Login
* **Method:** `POST`
* **Path:** `/auth/login`
* **Auth Required:** No
* **Request Body:**
  ```json
  {
    "username": "admin",
    "password": "admin123"
  }
  ```
* **Success Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Login successful",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "admin": {
        "id": "d748f3b8-6202-4b2a-89a1-bf8713a2862a",
        "name": "Portfolio Admin",
        "username": "admin",
        "email": "admin@example.com",
        "created_at": "2026-05-28T23:42:00Z",
        "updated_at": "2026-05-28T23:42:00Z"
      }
    }
  }
  ```

---

## 2. Projects Endpoints

### Get All Projects
* **Method:** `GET`
* **Path:** `/projects`
* **Auth Required:** No
* **Sorting:** Sorted by `order` asc, then `created_at` desc.
* **Success Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Projects retrieved successfully",
    "data": [
      {
        "id": "e4811b70-7619-482a-bc91-2384a92c81da",
        "title": "YOLOv8 Palm Oil Leaf Disease Detection",
        "description": "Applied deep learning segmentation...",
        "thumbnail": "/uploads/projects/sample-yolo.webp",
        "github_url": "https://github.com/...",
        "demo_url": "",
        "tech_stack": "Python,YOLOv8,PyTorch,OpenCV",
        "order": 1,
        "created_at": "2026-05-28T23:42:00Z"
      }
    ]
  }
  ```

### Get Project by ID
* **Method:** `GET`
* **Path:** `/projects/:id`
* **Auth Required:** No

### Create Project
* **Method:** `POST`
* **Path:** `/projects`
* **Auth Required:** Yes
* **Request Body:**
  ```json
  {
    "title": "New System Project",
    "description": "Detailed description of the project.",
    "thumbnail": "/uploads/projects/thumbnail.png",
    "github_url": "https://github.com/...",
    "demo_url": "https://...",
    "tech_stack": "React,Golang,PostgreSQL",
    "order": 3
  }
  ```

### Update Project
* **Method:** `PUT`
* **Path:** `/projects/:id`
* **Auth Required:** Yes
* **Request Body:** Same fields as Create Project (pass updated fields).

### Delete Project
* **Method:** `DELETE`
* **Path:** `/projects/:id`
* **Auth Required:** Yes

---

## 3. Experiences Endpoints

### Get All Experiences
* **Method:** `GET`
* **Path:** `/experiences`
* **Auth Required:** No
* **Sorting:** Sorted by `order` asc, then `created_at` desc.
* **Success Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Experiences retrieved successfully",
    "data": [
      {
        "id": "b3e34b1a-8c90-4823-92aa-bf610c3e1a8a",
        "company": "Tech Corp",
        "role": "Backend Engineer",
        "location": "Jakarta, ID",
        "start_date": "July 2023",
        "end_date": "Present",
        "current_job": true,
        "description": "- Integrated new third-party logistics APIs...\n- Maintained GORM databases.",
        "order": 1,
        "created_at": "2026-05-28T23:42:00Z"
      }
    ]
  }
  ```

### Create / Update / Delete Experience
* **Endpoints:** `POST /experiences`, `PUT /experiences/:id`, `DELETE /experiences/:id`
* **Auth Required:** Yes
* **Request Body Example:**
  ```json
  {
    "company": "Innovations Ltd",
    "role": "AI Engineer",
    "location": "Remote",
    "start_date": "Jan 2022",
    "end_date": "June 2023",
    "current_job": false,
    "description": "Built neural forecasting models...",
    "order": 2
  }
  ```

---

## 4. Certificates Endpoints

### Get All Certificates
* **Method:** `GET`
* **Path:** `/certificates`
* **Auth Required:** No

### Create / Update / Delete Certificate
* **Endpoints:** `POST /certificates`, `PUT /certificates/:id`, `DELETE /certificates/:id`
* **Auth Required:** Yes
* **Request Body Example:**
  ```json
  {
    "title": "AWS Certified Solutions Architect",
    "issuer": "Amazon Web Services",
    "issue_date": "March 2025",
    "credential_url": "https://aws.amazon.com/...",
    "thumbnail": "/uploads/certificates/aws.png",
    "order": 1
  }
  ```

---

## 5. Skills Endpoints

### Get All Skills
* **Method:** `GET`
* **Path:** `/skills`
* **Auth Required:** No
* **Success Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Skills retrieved successfully",
    "data": [
      {
        "id": "e98e8211-1a3b-41ca-ab1c-32bdf91e1d0f",
        "name": "Golang",
        "category": "Backend",
        "order": 1,
        "created_at": "2026-05-28T23:42:00Z"
      }
    ]
  }
  ```

### Create / Update / Delete Skill
* **Endpoints:** `POST /skills`, `PUT /skills/:id`, `DELETE /skills/:id`
* **Auth Required:** Yes
* **Request Body Example:**
  ```json
  {
    "name": "Python",
    "category": "Data Science & AI",
    "order": 2
  }
  ```

---

## 6. File Upload Endpoint

### Upload Image
* **Method:** `POST`
* **Path:** `/upload/:type`
* **Path Variable `:type`:** Either `project` or `certificate`
* **Headers:** `Content-Type: multipart/form-data`
* **Form Field Name:** `file` (or `image` as fallback)
* **Max Size:** 2MB
* **Allowed Types:** `jpg`, `jpeg`, `png`, `webp`, `gif`
* **Auth Required:** Yes
* **Success Response (201 Created):**
  ```json
  {
    "success": true,
    "message": "File uploaded successfully",
    "data": {
      "filename": "d38b1f24-ca82-41f2-95f2-51bc8c8d8b9c-1716943200.png",
      "url": "/uploads/projects/d38b1f24-ca82-41f2-95f2-51bc8c8d8b9c-1716943200.png",
      "size": 182390,
      "type": "image/png"
    }
  }
  ```

---

## 7. Analytics Endpoint

### Get Dashboard Summary
* **Method:** `GET`
* **Path:** `/analytics/summary`
* **Auth Required:** Yes
* **Success Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Analytics summary retrieved successfully",
    "data": {
      "projects": 5,
      "experiences": 0,
      "certificates": 0,
      "skills": 16
    }
  }
  ```

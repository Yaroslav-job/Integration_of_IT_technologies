# ✋ Gesture Recognition Results Microservice

A microservice in GoLang that collects, stores, and retrieves gesture recognition results for users in the **Russian Sign Language (RSL)** learning platform.

---

## 📦 Features

- `POST /recognitions`: Submit a recognized gesture result.
- `GET /recognitions/{userId}`: Fetch all results for a user.
- Uses in-memory DB for rapid prototyping (swappable with PostgreSQL/Redis).
- UUID-based result IDs and UTC timestamps.
- Fully unit-tested.

---

## 📁 Project Structure

gesture-recognition-results/  
├── go.mod  
├── main.go # Service entry point  
├── internal/  
    │ ├── handler/  
    │ │ ├── recognition.go # HTTP handlers  
    │ │ └── recognition_test.go # Unit tests for handlers  
    │ ├── db/  
    │ │ ├── memory.go # In-memory DB store  
    │ │ └── memory_test.go # DB tests  
    │ └── model/  
    │ └── recognition.go # Recognition result struct  
└── README.md  

---

## 🚀 Getting Started

### 1. Requirements

- Go 1.20+
- `go.work` file linking `gesture-recognition-results`

### 2. Setup

```
# bash
cd golang-services/
go work init ./services/gesture-recognition-results
cd services/gesture-recognition-results
go mod tidy
```
### 3. Run

```
# bash
go run main.go
```

---

## 🧪 API Reference

### POST `/recognitions`

Submit a new gesture recognition result.

#### Payload:

```json
{
  "user_id": "user42",
  "gesture": "hello",
  "confidence": 0.95
}
```

- `confidence`: float in range `[0.0 - 1.0]`

#### Response: `201 Created`

```json
{
  "id": "b312e3d4-...",
  "user_id": "user42",
  "gesture": "hello",
  "confidence": 0.95,
  "timestamp": "2025-04-15T15:00:00Z"
}
```

---

### GET `/recognitions/{userId}`

Fetch all gesture results for a given user.

#### Example:

```bash

curl http://localhost:8080/recognitions/user42
```

#### Response: `200 OK`

```json
[
  {
    "id": "...",
    "user_id": "user42",
    "gesture": "hello",
    "confidence": 0.95,
    "timestamp": "..."
  }
]
```

---

## 🧪 Running Tests

Unit tests are located in:
- `internal/db/memory_test.go`
- `internal/handler/recognition_test.go`

Run all tests:
```bash

go test ./internal
```

---

## 🔁 Next Steps

- Swap in-memory storage with PostgreSQL or Redis.
- Add webhook/event bus notification on new recognition.
- Add Swagger/OpenAPI documentation.
- Expose via gRPC for internal service mesh.
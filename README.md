# 📬 Real-Time Messaging System (Golang Microservices)

This project is a real-time, event-driven messaging system built using microservices architecture in Golang. It includes authentication, message ingestion, persistence with PostgreSQL, and Server-Sent Events (SSE) for real-time message delivery.

---

## 🧱 Services

| Service      | Description                                                              |
|--------------|--------------------------------------------------------------------------|
| Auth         | User registration & login with JWT token                                 |
| Ingestion    | (Optional) Receives API `/api/message` and publishes to Kafka            |
| Persistence  | Consumes Kafka messages and stores them in PostgreSQL                    |
| SSE Service  | Subscribes to Kafka and sends real-time messages via Server-Sent Events  |

---

## ⚙️ Tech Stack

- Golang + Gin
- PostgreSQL + GORM
- Apache Kafka
- Kafka-Go library
- Docker (optional)
- Kubernetes manifests (optional)
- Unit Testing

---

## 📂 Environment Variables

Create a `.env` file for each service using the following template (adjust if needed):

```env
# Common (Kafka)
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=message-topic

# Auth Service
PORT=9090
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=auth_service
JWT_SECRET=secret

# Ingestion Service
PORT=9091

# Persistence Service
PORT=9092
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=messages_db

# install depedencies
go mod tidy
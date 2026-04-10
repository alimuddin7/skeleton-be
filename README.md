# 🦴 Skeleton-BE Generator
[![Go Version](https://img.shields.io/github/go-mod/go-version/alimuddin7/skeleton-be)](https://golang.org/)
[![Release](https://img.shields.io/github/v/tag/alimuddin7/skeleton-be?label=release)](https://github.com/alimuddin7/skeleton-be/tags)
[![License](https://img.shields.io/github/license/alimuddin7/skeleton-be)](LICENSE)

**Skeleton-BE** is a production-grade, interactive CLI boilerplate generator for Go microservices. Optimized for high-performance and maintainability, it scaffolds projects based on **Fiber v3**, **GORM v2**, and **Zerolog**, strictly adhering to **Clean Architecture** principles.

---

## ✨ Key Features

- 🛠 **Interactive CLI**: Seamless onboarding experience using [Huh](https://github.com/charmbracelet/huh).
- 🏗 **Clean Architecture**: Standardized layers (Controller, Usecase, Repository, Model) ensuring clear separation of concerns.
- 🔌 **Plug & Play Infrastructure**: Instant integration for popular databases, caches, and message brokers.
- 📡 **gRPC Ready**: Built-in support for both gRPC Servers and Clients with automated proto management.
- 🏢 **Multi-Host Integration**: Scaffold robust clients for external service integrations.
- 🧪 **Observability**: Centralized logging with Zerolog, including automatic TraceID propagation across background workers.
- 🐳 **Containerization**: Multi-stage Dockerfiles and environment-specific Docker Compose configurations.
- 🚀 **DevOps Friendly**: Pre-configured GitLab CI/CD pipelines and SonarQube support.

---

## 🚀 Installation

Install the CLI globally:

```bash
go install github.com/alimuddin7/skeleton-be@latest
```

> [!TIP]
> Ensure your `$GOPATH/bin` is added to your system's `PATH`.

---

## 📚 Core Libraries
This generator leverages industry-standard, high-performance libraries to construct your microservices:

**Generator CLI Stack**:
- [**Cobra**](https://github.com/spf13/cobra): Powerful CLI application routing.
- [**Huh**](https://github.com/charmbracelet/huh): Beautiful terminal forms & interactive prompts.

**Generated Application Stack**:
- **Framework**: [Fiber v3](https://github.com/gofiber/fiber) (Ultra-fast HTTP routing) / Google gRPC
- **Database (ORM)**: [GORM v2](https://gorm.io/) (PostgreSQL & MySQL drivers)
- **Logging**: [Zerolog](https://github.com/rs/zerolog) (Zero-allocation JSON logger)
- **Validation**: [Validator v10](https://github.com/go-playground/validator) (Struct validation)
- **Messaging/Queue**: 
  - [NATS JetStream](https://github.com/nats-io/nats.go)
  - Kafka
  - [Asynq](https://github.com/hibiken/asynq) (Redis-based tasks)
- **Caching**: [Go-Redis v9](https://github.com/redis/go-redis)
- **Storage**: [MinIO go v7](https://github.com/minio/minio-go)

---

## 🛠 Usage Guide

### 1. Project Initialization
Scaffold a complete microservice in seconds via an interactive 9-step wizard:

```bash
skeleton-be init
```

The wizard covers:
1. **Project Name** & **Service Code** (Identifier)
2. **Project Type** (Backend, Scheduler, Worker, Publisher, gRPC) - *You can only select exactly **ONE** service type per project.*
3. **Database** (MySQL or PostgreSQL)
4. **Infra Modules** (Redis, Kafka, NATS, etc.)
5. **Messaging Roles** (Consumer, Publisher, or Both)
6. **External Hosts** & **gRPC Mode**

### 2. Available Infrastructure Modules
| Module | Description | Supported Roles |
| :--- | :--- | :--- |
| **MySQL / Postgre** | SQL Databases with GORM v2 | Primary DB |
| **Redis** | Standalone or Cluster mode | Cache / Queue |
| **NATS JetStream** | Cloud-native messaging | Consumer / Publisher |
| **Kafka** | High-throughput distributed queue | Consumer / Publisher |
| **Asynq** | Redis-based background processing | Worker / Client |
| **MinIO** | S3-compatible object storage | Storage |

### 3. Compatibility Matrix
Different Service Types support different architectural targets and CLI commands.

| Service Type (y) \ Modules & CLIs (x) | SQL / Redis / Minio | Message Brokers | Asynq Jobs | `add route` | `add crud` | `add host` | `add module` |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **Backend (REST API)** | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| **gRPC Server** | ✅ | ✅ | ✅ | ❌ | ❌ | ✅ | ✅ |
| **Worker / Consumer** | ✅ | ✅ (Consumer mode) | ❌ | ❌ | ❌ | ✅ | ✅ |
| **Publisher** | ✅ | ✅ (Publisher mode) | ❌ | ❌ | ❌ | ✅ | ✅ |
| **Scheduler** | ✅ | ✅ | ✅ | ❌ | ❌ | ✅ | ✅ |

### 4. Component Generation
Extend your existing project with specialized components:

```bash
# Add new infrastructure
skeleton-be add module kafka

# Generate full CRUD for an entity (Fiber Backend Only)
skeleton-be add crud user

# Add a simple route/feature stack (Fiber Backend Only)
skeleton-be add route healthcheck

# Integrate an external API
skeleton-be add host payment-gateway
```

### 4. Database Migrations
Create standardized SQL migration files:

```bash
skeleton-be migrate create add_status_to_users
```

---

## 📂 Project Structure

```text
├── cmd/                # Entry points (main.go)
├── configs/            # Config parsing (Env/YAML)
├── constants/          # App constants (TraceID, Error Codes)
├── controllers/        # Delivery layer (Fiber/gRPC handlers)
├── models/             # DTOs, Entities, and Validations
├── repositories/       # Data layer (SQL, Redis, NATS)
├── usecases/           # Pure Business Logic
├── hosts/              # External service clients
├── helpers/            # Shared utilities (Context, Logger, Auth)
├── routers/            # HTTP Route registrations
├── docker/             # Multi-stage Docker setups
└── errorcodes/         # Multilingual error messages (JSON)
```

---

## 🤝 Contribution

Contributions are welcome! Please feel free to submit Pull Requests or open Issues for feature requests.

---

*Developed with ❤️ by [ahmadfikrialimudin](https://github.com/ahmadfikrialimudin)*

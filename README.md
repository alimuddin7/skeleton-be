# 🦴 Skeleton-BE Generator
[![Go Version](https://img.shields.io/github/go-mod/go-version/alimuddin7/skeleton-be)](https://golang.org/)
[![Release](https://img.shields.io/github/v/tag/alimuddin7/skeleton-be?label=release)](https://github.com/alimuddin7/skeleton-be/tags)
[![License](https://img.shields.io/github/license/alimuddin7/skeleton-be)](LICENSE)

**Skeleton-BE** is a modern, interactive CLI boilerplate generator for Go microservices. Built on top of **Fiber v3**, **GORM v2**, and **Zerolog**, it follows Clean Architecture principles to help you scaffold production-ready services in seconds.

---

## ✨ Key Features

- 🛠 **Interactive CLI**: Smooth wizard experience powered by [Huh](https://github.com/charmbracelet/huh) and [Fang](https://github.com/charmbracelet/fang).
- 🏗 **Clean Architecture**: Separated layers for Controllers, Usecases, Repositories, and Models.
- 🔌 **Plug & Play Modules**: Easily add Redis, Kafka, NATS, MinIO, and more.
- 📡 **gRPC Support**: Built-in templates for gRPC Server and Client.
- 🏢 **Multi-Host Integration**: Ready-to-go templates for integrating with external API hosts.
- 🐳 **Docker Ready**: Pre-configured Dockerfile and Docker Compose for Dev, Staging, and Prod.
- 🚀 **GitLab CI/CD**: Complete CI/CD pipelines including SonarQube scanning.

---

## 🚀 Installation

Install the CLI globally using Go:

```bash
go install github.com/alimuddin7/skeleton-be@latest
```

> [!TIP]
> Make sure your `$GOPATH/bin` is in your system's `PATH` to run `skeleton-be` from anywhere.

---

## 🛠 Usage

### 1. Initialize a New Project
Start the 8-step interactive wizard to scaffold your service:

```bash
skeleton-be init
```

The wizard will guide you through:
1. **Project Name** - e.g., `payment-service`
2. **Service Code** - e.g., `OF01`, `OAG02`
3. **Project Type** - Backend, Scheduler, Worker, etc.
4. **Primary Database** - MySQL or PostgreSQL
5. **Additional Modules** - Select Redis, Kafka, etc.
6. **External API Hosts** - Input host names (comma-separated)
7. **Asynq** - Redis-based background queues
8. **gRPC Support** - Server, Client, or Both

### 2. Add Components to Existing Project
Keep your project growing with simple commands:

```bash
# Add a module (e.g., redis)
skeleton-be add module redis

# Add an external API host integration
skeleton-be add host payment-core

# Generate full CRUD (Controller, Usecase, Repository, Model, Routes)
skeleton-be add crud user --db mysql

# Add specific helper or route
skeleton-be add helper jwt
skeleton-be add route transaction
```

### 3. Database Migrations
Generate standardized migration files:

```bash
skeleton-be migrate create create_users_table
```

---

## 📂 Project Structure

Skeleton-BE generates a strict and clean directory structure:

```text
├── cmd/                # Entry points
├── configs/            # Configuration logic (Env/YAML)
├── constants/          # Application-wide constants
├── controllers/        # Delivery layer (Fiber handlers)
├── models/             # Business entities & DTOs (Request/Response)
├── repositories/       # Data layer (SQL, NoSQL, Cache)
├── usecases/           # Business logic layer
├── hosts/              # External API integrations
├── helpers/            # Cross-cutting utilities (Auth, HTTP, Logger)
├── routers/            # Route definitions
├── docker/             # Docker configurations
└── errorcodes/         # Standardized error definitions
```

---

## 🤝 Contribution

Feel free to open issues or submit pull requests. Let's make Go microservice development faster and cleaner!

---

*Authored by [ahmadfikrialimudin](https://github.com/alimuddin7)*

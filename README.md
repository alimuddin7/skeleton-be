# Skeleton Backend Generator

A dynamic boilerplate generator for Fiber v3 microservices. This tool helps you quickly scaffold production-ready Go microservices using Fiber v3, GORM v2, and Zerolog, following clean architecture patterns.

## Installation

To install the generator, run:

```bash
go install github.com/alimuddin7/skeleton-be@latest
```

Ensure that your `$GOPATH/bin` is in your system's `PATH`.

## Usage

### Initialize a New Project

To create a new project, simply run:

```bash
skeleton-be init
```

This will launch an interactive wizard to configure your new microservice. You can choose:
- Project Name
- Service Code
- Project Type (Backend, Scheduler, Worker, Publisher, gRPC)
- Database (MySQL, PostgreSQL)
- Additional modules (Redis, Kafka, NATS, MinIO, etc.)

Alternatively, you can use flags for non-interactive mode:

```bash
skeleton-be init --name my-service --code 01 --type Backend --db postgresql
```

### Add Modules to Existing Project

You can add new modules to an existing project:

```bash
skeleton-be add redis
skeleton-be add kafka
```

### Add a new CRUD Feature
Generate a full CRUD feature set (Controller, Usecase, Repository, Model, Routes):

```bash
skeleton-be add crud <name> --db <mysql|postgresql>
```

### Add a new Helper
Generate a new helper function stub:

```bash
skeleton-be add helper <name>
```

### Database Migrations
Manage database migrations:

```bash
# Create a new migration file
skeleton-be migrate create <name>
```

### Available Helpers
The following helpers are included by default:
- **HTTP Helper**: `net/http` wrapper for making external requests.
- **JWT Helper**: Token generation and parsing.
- **Auth Helper**: Password hashing and authentication.
- **Pagination Helper**: Pagination struct generator.
- **General Helper**: Env loading, date conversion, unique array, currency format, etc.
- **Meta Helper**: Standardized API response metadata.
- **Error Helper**: Error handling utilities.
- **Gorm Logger**: Custom GORM logger using Zerolog.
- **Health Check**: Health check response generator.

## Project Structure

### Add External Host Integration

Add an integration with an external service:

```bash
skeleton-be add host payment-service
```

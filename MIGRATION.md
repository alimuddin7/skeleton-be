# Migration Guide: skeleton-svc to skeleton-be

This document outlines the changes and usage differences between the legacy `go-archetype` based `skeleton-svc` and the new `cobra` based `skeleton-be` CLI.

## Overview

| Feature | Legacy (`skeleton-svc` / `go-archetype`) | New (`skeleton-be` CLI) |
| :--- | :--- | :--- |
| **Core Tool** | `go-archetype` | `cobra` (Native Go CLI) |
| **Installation** | `go install github.com/omeid/go-archetype@latest` | `go install github.com/alimuddin7/skeleton-be@latest` |
| **Initialization** | `go-archetype -t <repo> ...` | `skeleton-be init` |
| **Interactivity** | Limited / Flag-based | Fully Interactive (PromptUI) & Flag-based |
| **Module Support** | Pre-defined in template | Dynamic selection (Add modules as needed) |
| **Feature Generation**| Manual copy-paste or static | `skeleton-be add route` / `skeleton-be add crud` |
| **Maintenance** | Hard to update templates | Embedded templates, easy to version |

## Command Mapping

### 1. Initialize Project

**Legacy:**
```bash
go-archetype -s . -d ../my-service -t https://github.com/alimuddin7/skeleton-svc.git -- --name=my-service
```

**New:**
```bash
# Interactive
skeleton-be init

# Non-interactive
skeleton-be init --name my-service --code 01 --type Backend --db postgresql
```

### 2. Add Database/Infra

**Legacy:**
Usually involved uncommenting code or manually copying files.

**New:**
```bash
skeleton-be add redis
skeleton-be add kafka
skeleton-be add minio
```

### 3. Create New Endpoint/Feature

**Legacy:**
Manually create:
- `models/feature.go`
- `repositories/feature.go`
- `usecases/v1/feature.go`
- `controllers/v1/feature.go`
- Register manually in `router.go` and `wire.go` (if used).

**New:**
```bash
# Generate full CRUD stack
skeleton-be add crud product --db postgresql

# Generate basic feature stack
skeleton-be add route custom_feature
```

### 4. Database Migrations

**Legacy:**
Manual SQL file creation or external tool.

**New:**
```bash
skeleton-be migrate create init_schema
```

## Feature Status

| Component | Status | Notes |
| :--- | :--- | :--- |
| **Fiber v3** | ✅ Migrated | Fully supported in base template |
| **GORM v2** | ✅ Migrated | MySQL, PostgreSQL supported |
| **Databases** | ✅ Migrated | MySQL, Postgres, Redis, Mongo (planned) |
| **Message Broker** | ✅ Migrated | Kafka, NATS, Redis PubSub |
| **Storage** | ✅ Migrated | MinIO |
| **gRPC** | ✅ Migrated | Server & Client support |
| **Logging** | ✅ Migrated | Zerolog integration |
| **Docker** | ✅ Migrated | Dockerfile & Compose templates |
| **CI/CD** | ✅ Migrated | GitLab CI template |

## Next Steps

- Use `skeleton-be` for all new microservice generation.
- For existing services, you can continue using them as is, or use `skeleton-be add` commands to inject new modules (requires standard folder structure).

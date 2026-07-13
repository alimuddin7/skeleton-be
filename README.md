# skeleton-be

> **CLI boilerplate generator** untuk Go microservices berbasis **Fiber v3** — scaffold project siap produksi dalam hitungan detik.

---

## Daftar Isi

- [Tentang Project](#tentang-project)
- [Teknologi](#teknologi)
- [Arsitektur yang Dihasilkan](#arsitektur-yang-dihasilkan)
- [Struktur Direktori CLI](#struktur-direktori-cli)
- [Fitur](#fitur)
- [Kelebihan & Kekurangan](#kelebihan--kekurangan)
- [Instalasi](#instalasi)
- [Cara Penggunaan](#cara-penggunaan)
- [Referensi Command](#referensi-command)

---

## Tentang Project

`skeleton-be` adalah CLI tool yang mengotomatisasi pembuatan boilerplate Go microservice. Alih-alih menyalin template secara manual, cukup jalankan satu perintah dan dapatkan project yang sudah terstruktur rapi dengan integrasi infrastruktur (database, cache, messaging, storage), CI/CD pipeline, dan Docker siap pakai.

**Target pengguna:** Backend engineer yang ingin memulai microservice baru dengan standar Clean Architecture tanpa setup berulang.

---

## Teknologi

### CLI Layer

| Library | Versi | Fungsi |
|---|---|---|
| [Cobra](https://github.com/spf13/cobra) | v1.10 | Command & subcommand routing |
| [Charmbracelet Fang](https://github.com/charmbracelet/fang) | v0.4 | CLI entrypoint wrapper (help, version) |
| [Charmbracelet Huh](https://github.com/charmbracelet/huh) | v0.8 | Interactive form / wizard TUI |
| [Charmbracelet Bubbletea](https://github.com/charmbracelet/bubbletea) | v1.3 | TUI rendering engine (transitive) |
| [Lipgloss](https://github.com/charmbracelet/lipgloss) | v1.1 | Terminal UI styling (transitive) |

### Template Engine

| Teknologi | Keterangan |
|---|---|
| `text/template` (stdlib) | Render template Go ke file project |
| `embed.FS` | Template dikompilasi langsung ke dalam binary |
| `golang.org/x/text` | Transformasi teks (PascalCase, Title, dsb.) |

### Project yang Di-generate

| Komponen | Library |
|---|---|
| HTTP Framework | [Fiber v3](https://github.com/gofiber/fiber) |
| ORM | [GORM v2](https://gorm.io) |
| Logger | [Zerolog](https://github.com/rs/zerolog) |
| Database | MySQL · PostgreSQL |
| Cache | Redis Standalone · Redis Cluster |
| Messaging | NATS JetStream · Kafka · Asynq |
| Storage | MinIO |
| Scheduler | [robfig/cron](https://github.com/robfig/cron) |
| RPC | gRPC (Server / Client / Both) |
| CI/CD | GitLab CI/CD (build, deploy, sonar scanner) |
| Container | Docker + Docker Compose (dev, stg, prod) |

---

## Arsitektur yang Dihasilkan

Project yang di-generate mengikuti pola **Clean Architecture**:

```
your-service/
├── cmd/                        # Entrypoint aplikasi
├── configs/                    # Konfigurasi environment
├── constants/                  # Konstanta global
├── controllers/
│   └── v1/                     # HTTP handlers (Fiber)
├── usecases/
│   └── v1/                     # Business logic layer
├── repositories/
│   ├── mysql/                  # MySQL repository (jika dipilih)
│   ├── postgre/                # PostgreSQL repository (jika dipilih)
│   ├── redis/                  # Redis repository (jika dipilih)
│   ├── kafka/                  # Kafka producer/consumer
│   ├── nats/                   # NATS JetStream
│   └── asynq/                  # Asynq task queue
├── models/
│   └── dto/                    # Data Transfer Objects
├── routers/                    # Route registration
├── helpers/                    # Utilities (auth, middleware, http, logger)
├── errorcodes/                 # Error code mapping (JSON)
├── consumers/                  # Message consumer entrypoint
├── scheduler/                  # Cron job scheduler
├── grpc/
│   ├── server/                 # gRPC server
│   ├── client/                 # gRPC client
│   └── proto/                  # Protobuf definitions
├── hosts/                      # External API client wrappers
├── migrations/                 # SQL migration files
├── docker/                     # Dockerfile & docker-compose variants
├── .gitlab/                    # GitLab CI/CD pipeline configs
├── internal/app/               # App bootstrap & dependency wiring
├── skeleton.json               # Project state (untuk perintah add/remove)
├── .env.example
├── Makefile
└── .gitlab-ci.yml
```

---

## Struktur Direktori CLI

```
skeleton-be/                    # Source code CLI tool ini
├── cmd/
│   ├── root.go                 # Root command
│   ├── init.go                 # skeleton-be init (wizard 9-step)
│   ├── add.go                  # skeleton-be add (module/crud/route/host/helper)
│   ├── migrate.go              # skeleton-be migrate create
│   └── remove.go               # skeleton-be remove crud
├── internal/
│   └── generator/
│       ├── generator.go        # Core generator logic
│       └── templates/
│           ├── base/           # Base project templates
│           ├── domain/         # CRUD & feature domain templates
│           └── modules/        # Infra module templates (redis, kafka, dsb.)
└── main.go
```

---

## Fitur

### `init` — Interactive Project Wizard

9-step wizard yang menghasilkan project lengkap:

1. Nama project
2. Service code (identifier 2-digit, contoh: `OF01`)
3. Tipe project (multi-select): **Backend**, **Scheduler**, **Worker**, **Publisher**, **gRPC**
4. Primary database: **MySQL** atau **PostgreSQL**
5. Modul tambahan: **Redis**, **Redis Cluster**, **MinIO**
6. Messaging broker: **NATS JetStream**, **Kafka**, **Asynq**
7. Messaging role: **Consumer**, **Publisher**, **Both**
8. External API hosts (HTTP client wrapper per host)
9. gRPC support: **Server**, **Client**, **Both**

Setelah wizard selesai, generator akan:
- Membuat seluruh struktur direktori
- Merender semua template ke file Go yang valid
- Menjalankan `go mod tidy` otomatis
- Menginisialisasi git repository

### `add` — Extend Project yang Sudah Ada

| Subcommand | Fungsi |
|---|---|
| `add module [name]` | Tambah modul infrastruktur (redis, kafka, nats, dll) |
| `add crud [name]` | Generate stack CRUD lengkap (controller, usecase, repository, model, DTO) |
| `add route [name]` | Generate stack non-CRUD (controller + usecase saja) |
| `add host [name]` | Tambah external API client wrapper |
| `add helper [name]` | Tambah helper function stub |

### `migrate` — Migration Management

| Subcommand | Fungsi |
|---|---|
| `migrate create [name]` | Buat file `.up.sql` dan `.down.sql` bertimestamp |

### `remove` — Hapus Komponen

| Subcommand | Fungsi |
|---|---|
| `remove crud [name]` | Hapus seluruh file CRUD stack dan update state |

---

## Kelebihan & Kekurangan

### ✅ Kelebihan

| Aspek | Detail |
|---|---|
| **Zero setup** | Satu perintah menghasilkan project lengkap siap compile & run |
| **Single binary** | Semua template di-embed ke dalam binary via `embed.FS`, tidak butuh internet atau file eksternal |
| **Stateful** | `skeleton.json` menyimpan state project sehingga `add` dan `remove` bisa dijalankan kapan saja tanpa perlu reinit |
| **Idempotent** | `add module` mengecek apakah modul sudah ada sebelum menambahkan |
| **Clean Architecture** | Struktur yang di-generate memisahkan controller, usecase, repository secara tegas |
| **CI/CD out-of-the-box** | GitLab CI/CD dengan tahap build, deploy, sonar scanner sudah disertakan |
| **Multi-tipe project** | Satu project bisa sekaligus Backend + Worker + Scheduler |
| **Interactive & non-interactive** | Bisa dijalankan via wizard TUI atau flag langsung (cocok untuk CI/scripting) |

### ⚠️ Kekurangan

| Aspek | Detail |
|---|---|
| **Tidak ada unit test** | Generator belum memiliki test coverage untuk memvalidasi output template |

---

## Instalasi

### Prasyarat

- Go **1.21+**
- Git

### Install via Go

```bash
go install github.com/alimuddin7/skeleton-be@latest
```

Binary akan otomatis tersedia di `$GOPATH/bin`. Pastikan `$GOPATH/bin` sudah ada di `$PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Verifikasi

```bash
skeleton-be --help
```

---

## Cara Penggunaan

### 1. Membuat Project Baru (Interactive)

```bash
skeleton-be init
```

Wizard akan memandu 9 langkah konfigurasi. Setelah selesai, folder project langsung dibuat di direktori saat ini.

### 2. Membuat Project Baru (Non-Interactive / Scripted)

```bash
skeleton-be init \
  --name payment-service \
  --code PS01 \
  --type Backend \
  --db postgresql \
  --modules redis,minio \
  --hosts core-api,user-service \
  --grpc No
```

### 3. Menambah Modul ke Project yang Ada

```bash
# Masuk ke dalam project terlebih dahulu
cd payment-service

# Tambah Redis
skeleton-be add module redis

# Tambah NATS (akan ditanya role: Consumer/Publisher/Both)
skeleton-be add module nats
```

### 4. Generate CRUD Baru

```bash
skeleton-be add crud transaction
skeleton-be add crud product --db postgresql
```

Menghasilkan file:
- `controllers/v1/transaction.controller.go`
- `usecases/v1/transaction.usecase.go`
- `repositories/postgre/transaction.go`
- `models/transaction.go`
- `models/dto/transaction.go`

### 5. Generate Route Non-CRUD

```bash
skeleton-be add route healthcheck
skeleton-be add route login
```

### 6. Tambah External API Client

```bash
skeleton-be add host midtrans
skeleton-be add host xendit
```

Menghasilkan `hosts/midtrans/host.go` dengan HTTP client wrapper yang sudah terstruktur.

### 7. Buat Helper Baru

```bash
skeleton-be add helper password-hash
```

### 8. Buat Migration File

```bash
skeleton-be migrate create create_transactions_table
```

Menghasilkan:
- `migrations/20260609120000_create_transactions_table.up.sql`
- `migrations/20260609120000_create_transactions_table.down.sql`

### 9. Hapus CRUD Stack

```bash
skeleton-be remove crud transaction
```

---

## Referensi Command

```
skeleton-be
├── init          Inisialisasi project baru (wizard atau flag)
├── add
│   ├── module    Tambah modul infrastruktur
│   ├── crud      Generate stack CRUD entity
│   ├── route     Generate stack route non-entity
│   ├── host      Tambah external API client
│   └── helper    Tambah helper stub
├── migrate
│   └── create    Buat file migrasi SQL
└── remove
    └── crud      Hapus stack CRUD
```

### Flag Global `init`

| Flag | Shorthand | Keterangan |
|---|---|---|
| `--name` | `-n` | Nama project |
| `--code` | `-c` | Service code |
| `--type` | `-t` | Tipe project (Backend, Scheduler, Worker, Publisher, gRPC) |
| `--db` | `-d` | Primary database (mysql, postgresql) |
| `--modules` | `-m` | Modul tambahan (redis, kafka, nats, minio) |
| `--hosts` | `-H` | External API hosts, pisah koma |
| `--grpc` | `-g` | gRPC mode (No, Server, Client, Both) |

---

> Dibuat dengan ❤️ .

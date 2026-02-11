# Component Parity Report

Comparison between `skeleton-svc` (Source) and `skeleton-be` (Target).

## Infrastructure & Config
| Component | skeleton-svc | skeleton-be | Status |
| :--- | :--- | :--- | :--- |
| **Makefile** | Extensive (test, build, report, docker, mock, lint, proto, migration) | Basic (test, build, run, docker) | ⚠️ Missing `mock`, `lint`, `proto` commands |
| **Docker** | Dockerfile, Dockerfile.debug | Dockerfile, docker-compose | ✅ Improved (Compose added) |
| **CI/CD** | .gitlab-ci.yml | .gitlab-ci.yml | ✅ Present |
| **Docs** | `docs` folder | - | ⚠️ Missing default docs structure |
| **Mocks** | `mocks` folder & command | - | ⚠️ Missing mock generation |

## Modules (Databases/Brokers)
| Module | skeleton-svc | skeleton-be | Status |
| :--- | :--- | :--- | :--- |
| **MySQL** | ✅ | ✅ | Pariety |
| **PostgreSQL**| ✅ | ✅ | Pariety |
| **Redis** | ✅ | ✅ | Pariety |
| **Redis Cluster**| ✅ | ✅ | Pariety |
| **Redis Asynq** | ✅ | ❓ (Logic exists, templates missing?) | ⚠️ Needs Verification |
| **Google PubSub** | ✅ | ❌ | 🔴 Missing |
| **MQTT** | ✅ | ❌ | 🔴 Missing |
| **Cassandra** | ✅ | ❌ | 🔴 Missing |
| **MongoDB** | ✅ (implied) | ❌ | 🔴 Missing |
| **Kafka** | ✅ | ✅ | Pariety |
| **NATS** | ✅ | ✅ | Pariety |
| **MinIO** | ✅ | ✅ | Pariety |

## Core Components
| Component | skeleton-svc | skeleton-be | Status |
| :--- | :--- | :--- | :--- |
| **Middlewares**| Extensive | Basic (Logger, Recover, CORS) | ⚠️ Need to check feature parity |
| **Helpers** | Extensive | Basic | ⚠️ Significant gaps |
| **gRPC** | Client & Server | Client & Server | ✅ Pariety |
| **Scheduler** | ✅ | ✅ | ✅ Pariety |

## Helper & Function Analysis

| Helper | skeleton-svc | skeleton-be | Notes |
| :--- | :--- | :--- | :--- |
| **Auth/JWT** | `auth.helpers.go`, `jwt.helpers.go` | ❌ | 🔴 Missing Auth/JWT helpers |
| **HTTP** | `http.helpers.go`, `gin.helpers.go` | ❌ | 🔴 Missing HTTP helpers (Fiber equivalent needed) |
| **Pagination**| `pagination.helpers.go` | ❌ | 🔴 Missing Pagination helpers |
| **General** | `general.helpers.go` | `general_helpers.go` | ✅ Present, check content coverage |
| **Error** | `error.helpers.go` | `error_helpers.go` | ✅ Present |
| **Meta** | `meta.helpers.go` | `meta_helpers.go` | ✅ Present |
| **Health Check**| `hc.helpers.go` | `hc_helpers.go` | ✅ Present |
| **Logger** | `loggers.helpers.go`, `new-logger.helpers.go` | `logger.go` | ⚠️ Consolidated/Simplified? |
| **Gorm Logger** | `gorm-logger.helpers.go` | `gorm_logger.go` | ✅ Present |

## Action Items
1.  **Add Missing Modules**: Google PubSub, MQTT, Cassandra, MongoDB.
2.  **Verify Redis Asynq**: Ensure templates are present and correct.
3.  **Enhance Makefile**: Add `mock`, `lint`, `proto` commands.
4.  **Docs**: Add default `docs` folder or Swagger generation setup.
5.  **Add Helpers**: Port Auth/JWT, Pagination, and generic HTTP helpers (adapted for Fiber).

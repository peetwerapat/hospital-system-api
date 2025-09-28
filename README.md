## Setup & Run

### 1. Run Docker compose
```
docker compose up -d --build
```

### 2. Migrate Table
```
make migrate-up
```
---

## Project Structure

```
├── cmd/
│ └── server/
│   └── main.go                               # Application entry point
├── internal/
│ ├── domain/                                 # Core business entities
│ │ ├── hospital.go
│ │ ├── patient.go
│ │ └── staff.go
│ ├── usecase/                                # Business logic / use cases
│ │ ├── error.go
│ │ ├── patient_usecase.go
│ │ └── staff_usecase.go
│ ├── interface/
│ │ ├── controller/                           # Adapters (controllers, repositories)
│ │ │ ├── dto/                                # Data Transfer Objects
│ │ │ │ ├── patient_dto.go
│ │ │ │ └── staff_dto.go
│ │ │ ├── patient_controller.go
│ │ │ └── staff_controller.go
│ │ └── repository/                           # Repository interfaces
│ │   ├── hospital_repo_interface.go
│ │   ├── patient_repo_interface.go
│ │   └── staff_repo_interface.go
│ └── infrastructure/
│   ├── db/                                   # Database connection and setup
│   │ └── db.go
│   ├── repository_impl/                      # Repository implementations (GORM)
│   │ ├── hospital_repo_impl.go
│   │ ├── patient_repo_impl.go
│   │ └── staff_repo_impl.go
│   ├── router/                               # Gin router setup
│   │ └── router.go
│   └── di/                                   # Dependency injection setup
│     └── app.go
├── migrations/                               # SQL migration files
│ ├── 000001_init_table.up.sql
│ └── 000001_init_table.down.sql
├── nginx/                                    # Nginx config (reverse proxy)
│ └── conf.d/default.conf
├── pkg/
│ ├── config/                                 # Application configuration
│ │ └── config.go
│ ├── middleware/                             # Middleware
│ │ └── auth.go
│ ├── myJwt/                                  # JWT helpers
│ │ └── jwt.go
│ └── response/                               # API response formatter
│   └── response.go
├── test/
│ ├── mocks/                                  # Mock repository for unit tests
│ │ ├── mock_hospital_repo.go
│ │ ├── mock_patient_repo.go
│ │ └── mock_staff_repo.go
│ └── usecase/                                # Unit tests for use cases
│   ├── patient_usecase_test.go
│   └── staff_usecase_test.go
├── Dockerfile                              
├── docker-compose.yml
├── Makefile                                  # Make file migrate command
├── go.mod
├── go.sum
└── README.md
```

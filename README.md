# 🚀 Go Template - (Hexagonal + Clean Architecture)

Welcome, developers! 👋 This **Go Template** will accelerate your Go application development with a robust and maintainable architecture.

## 💫 Project Overview

This **Modular Monolith** combines Hexagonal (Ports/Adapters) and Clean Architecture principles to provide:

- 🔄 **Clean separation** between UI (HTTP) and persistence layers (DB, Cache)
- 🧠 **Domain-focused development** without infrastructure concerns
- 🔌 **Flexible infrastructure** - change databases or frameworks without affecting core logic
- 🧩 **Future scalability** - easily evolve into microservices when needed

## 🛠️ Technology Stack

- **Web Framework**: [Fiber v3](https://github.com/gofiber/fiber)
- **Dependency Injection**: [Uber Fx](https://github.com/uber-go/fx)
- **Database**: [GORM](https://gorm.io)
- **Cache / Session**: [go-redis v9](https://github.com/redis/go-redis)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Authentication**: JWT (github.com/golang-jwt/jwt/v5)
- **Validation**: go-playground/validator
- **Testing**: Testify + mockgen

## 🗂️ Project Structure

```
cmd/
  app/                # main.go (entry point)
internal/
  adapter/
    http/             # API endpoints and middleware
    persistence/      # Data models and mappers
    repository/       # Port implementations
  app/                # Module composition (Fx options)
  domain/
    auth/             # Authentication interfaces
    session/          # Session management interfaces
    user/             # Entities and repository interfaces
  service/
    auth/             # Authentication use cases
    user/             # User management use cases
  shared/             # Cross-cutting concerns (config, db, redis, jwt)
```

## 🚀 Getting Started

```bash
# 1. Clone and update dependencies
$ git clone https://github.com/ditthkr/go-template.git
$ cd go-template
$ go mod tidy

# 2. Configure your application
$ cp config.yaml.example config.yaml  # Remember to update DSN, Redis, JWT secret

# 3. Run development mode
$ make dev  # Loads ./config.yaml with ENV variable overrides
```

## 🌐 Core APIs

| Method | Path             | Auth | Description |
|--------|------------------|------|-------------|
| POST   | `/auth/register` | ✗    | Register with username and email |
| POST   | `/auth/login`    | ✗    | Login and receive JWT token |
| GET    | `/users/me`      | ✓    | Retrieve current user profile |

Response format:
```json
// Success
{"success":true, "data":{...}}
// Error
{"success":false, "message":"Error description"}
```

## 🧪 Testing

The project uses `gomock` for repository mocking and `testify` for assertions, supporting both unit and integration testing:

```bash
# Run all tests
$ go test ./...

# Test specific package
$ go test ./internal/service/auth -v

# Test specific function
$ go test ./internal/service/auth -run TestAuthService_Register
```

## 🙋‍♀️ Questions or Issues?

Please feel free to:
- Open an issue
- Submit a pull request
- Fork for customization

---
## 📝 MIT License
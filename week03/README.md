# Week 03 — “Add Postgres & Basic CRUD”

> Goal: integrate PostgreSQL into your app using pgxpool and implement basic Create/Read operations via a simple HTTP API built on Gin.

---

## 📝 Tasks & Steps
| # | Task | Hint / Command |
|---|------|----------------|
| 1	| Define a `User` model: `ID`, `Name`, `Email`, `CreatedAt`. |	Use Go struct and PostgreSQL schema. |
| 2	| Add pgxpool-based DB connection in `internal/db/db.go`. |	Use `pgxpool.ConnectConfig()`, handle errors. |
| 3	| Add DB config to `config.yaml`, support env/flag overrides.	| Extend your config struct with DB fields. |
| 4	| Refactor config loading to support DB config: `host`, `port`, `user`, `pass`, `name`. |	Respect override order: YAML < env < flags. |
| 5	| Add POST `/users`: read JSON, insert user into DB. |	Use `gin.Context.BindJSON`, `pgxpool.Exec` |
| 6 |	Add GET `/users/:id`: load user by ID.	| Use `gin.Context.Param("id")`, query with `QueryRow`. |
| 7 |	Add health check for DB connection. |	Create `internal/db/Health()` returning error. |
| 8 |	Add Postgres service in `docker-compose.yaml` with volume.	| Include `init.sql`, mount volume. Check examples |
| 9 |	Write integration test for insert + select via HTTP. |	Use HTTP test client + test DB. |
| 10 | Update CI to run Postgres and test your app against it. |	Use services: postgres in GitHub Actions. |
| 11 | Tag release `v0.3.0` and push.	| `git tag v0.3.0 && git push --tags` |

## 🔄 Architectural Extension (optional, but appreciated)

Create a UserStore interface and implement it using PostgreSQL.

This allows your API logic to be tested independently from the database (e.g., with SQLite or an in-memory map).
It enables better testing practices in CI and makes unit testing simpler using mock implementations.
This pattern is considered a professional standard — it helps keep your code flexible, testable, and decoupled from infrastructure.

Example:

```go
type User struct {
    ID        int64
    Name      string
    Email     string
    CreatedAt time.Time
}

type UserStore interface {
    CreateUser(ctx context.Context, u *User) error
    GetUserByID(ctx context.Context, id int64) (*User, error)
}
```

By injecting this interface into your HTTP handlers, you separate core logic from DB details — a key step toward clean architecture and scalable systems.

## 📁 Target File Structure

```
cmd/app/main.go
config.yaml
.env
Dockerfile
docker-compose.yaml
internal/config/load.go
internal/config/config.go
internal/server/server.go
internal/server/handlers.go
internal/db/db.go
internal/db/models.go
internal/db/db_test.go
sql/init.sql
.github/workflows/ci.yml
go.mod
```

## 📚 Useful References

- What is database schema? https://www.geeksforgeeks.org/dbms/database-schemas/
- PostreSQL data types: https://www.postgresql.org/docs/current/datatype.html
- Goose SQL Migration tool for Go: https://github.com/pressly/goose
- Gin quickstart: https://gin-gonic.com/en/docs/quickstart/

## ✅ Acceptance Criteria

- Postgres accessible via Docker Compose
- App connects to Postgres using pgxpool
- Config loaded with override order: flags > env > yaml
- POST /users accepts JSON, inserts into DB
- GET /users/:id fetches user from DB
- /healthz checks DB connectivity
- Integration test inserts and reads user
- CI pipeline spins up Postgres, runs tests
- Tag v0.3.0 is pushed

## 💡 Reflection Questions

- What makes a good database schema? How can poor schema design lead to long-term technical debt?
- Why use a connection pool instead of opening a new DB connection for every request? What are the performance and reliability tradeoffs?
- What could go wrong if schema migrations are done manually in production? How would you ensure safe, repeatable schema changes?
- In what ways can the database become a bottleneck in your system? How can you design around this?
- Why is it better to abstract your data access behind an interface? How does this affect testability, maintainability, and extensibility?
- What are the downsides of tightly coupling your HTTP handlers to your database logic? 
- How would you structure error responses for a REST API? What makes an API “developer-friendly”?
- When should you use SQLite over Postgres — and when should you avoid it? 
- What problems could arise from relying solely on configuration in .env or config.yaml? How can you make configuration secure and observable?
- How would you extend this app to support other storage backends (e.g., Redis, MySQL, file-based)? What changes would you not want to make in order to support that?

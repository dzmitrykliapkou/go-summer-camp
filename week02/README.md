# Week 02 — “Configurable HTTP Server in a Container”

> Goal: build a configurable HTTP server that runs via Docker Compose and supports flags, environment variables, and config files.

---

## 📝 Tasks & Steps

| # | Task | Hint / Command |
|---|------|----------------|
| 1  | **Refactor server startup into `internal/server/Start(cfg)`**                                                   | Start with a hardcoded config struct; isolate server logic. Define `port`, `greeting`, and `timeout` for the server. |
| 2  | **Add CLI flags** to configure `port`, `greeting`, and `timeout`.                                               | Use the `flag` package. Start with flags only.                 |
| 3  | **Add environment variable support** as a fallback if flags are not provided.                                   | Use `os.Getenv`, but only override missing flags.              |
| 4  | **Add YAML config file loading** as the base layer (lowest priority).                                           | Use `gopkg.in/yaml.v2`. Load `config.yaml` and fill defaults.  |
| 5  | **Refactor all config logic into internal/config.Load()**         | 🧠 Here’s the key architectural step: if config handling stays in main.go, it quickly becomes fragile and unreadable (“spaghetti code”). Refactor to keep main.go minimal, testable, and future-proof. |
| 6  | **Ensure proper config override order**: config.yaml < env vars < flags. | Use a `config.Load()` function that builds from the bottom up. Test different override scenarios. |
| 7  | **Expose a `/healthz` endpoint** returning 200 OK.                                                              | No logic – just useful in containers and CI.                   |
| 8  | **Add a Dockerfile** to build your app.                                                                         | Use `golang:1.22-alpine`. Simple build script is enough.       |
| 9  | **Add docker-compose.yaml** that runs your app with mounted `config.yaml` and `.env`.                           | Use `.env` to test env vars locally.                           |
| 10  | **Update CI** to build Docker image and run tests/lint.                                                         | Use dedicated docker GitHub Actions                       |
| 11 | **Tag release** as `v0.2.0` and push it.                                                                        | `git tag v0.2.0 && git push --tags`                            |

## 📌 About Configuration Handling

⚠️ It’s okay to keep flags, env vars, and config file loading in `main.go` while you're exploring each concept — this helps you stay focused on learning the mechanism. But once all three are added, your main.go will start to look like this:

```go
port := flag.String("port", "", ...)
flag.Parse()
if *port == "" {
    if env := os.Getenv("PORT"); env != "" {
        *port = env
    } else {
        // load from YAML...
    }
}
```

This quickly turns into **spaghetti code**: unreadable, tangled, hard to test, and impossible to reuse.

✅ Best practice: move all config-loading logic into `internal/config/load.go`. Let main.go stay small and clean — just wiring things together.

```go
func main() {
    cfg, err := config.Load()
    if err != nil { 
        log.Fatalf("Failed to load config: $v", err) 
    }
    server.Start(cfg)
}
```

## 📁 Target File Structure

```
cmd/app/main.go
internal/config/load.go
internal/config/config.go
internal/server/server.go
internal/server/server_test.go
config.yaml
.env
Dockerfile
docker-compose.yaml
.github/workflows/ci.yml
go.mod
```

## 📚 Useful References

- Flags: https://pkg.go.dev/flag
- Env: https://pkg.go.dev/os#Getenv
- YAML in Go: https://pkg.go.dev/gopkg.in/yaml.v2
- Timeouts in servers: https://blog.cloudflare.com/exposing-go-on-the-internet/#timeouts
- Dockerfile tips: https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
- Docker Compose: https://docs.docker.com/compose/
- GitHub Actions + Docker: https://docs.github.com/en/actions/how-tos/use-cases-and-examples/publishing-packages/publishing-docker-images

## ✅ Acceptance Criteria

- `config.Load()` implements source priority: `flags` > `env` > `yaml` > `default`.
- Server works with:
  ```
  go run ... --port=9000
  ```
- `.env` overrides `config.yaml` as base config
- `/` and `/healthz` work correctly.
- Dockerfile builds and runs the app.
- docker-compose up works with .env + config.yaml.
- CI builds, lints, and tests.
- Git tag `v0.2.0` is pushed.

## 💡 Reflection Questions

1. Why is it a bad idea to keep all config logic in main.go?
2. What makes "spaghetti code" dangerous in production apps?
3. What kinds of configuration are best handled with flags vs env vs file?
4. What can go wrong if your server has no timeouts or config validation?
5. How would you extend your config system to support JSON or TOML?

## 🧠 Reflection Theme: “Why configuration is a security problem”

We treat config as a boring setup detail. But what happens when someone puts secrets in a YAML file? Or misconfigures a public port? Or sets timeouts too high — or too low?

Read up on:
- Configuration drift
- Secrets in config files
- “Infrastructure as code” vs “configuration as code”
- Secure defaults and “the principle of least surprise”

Then reflect:
- What’s the risk of having an insecure or unpredictable config in production?
- How can config impact availability, security, and compliance?
- What practices can prevent config becoming a liability?

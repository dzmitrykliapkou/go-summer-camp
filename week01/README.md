# Week 01 – “Project Skeleton & First CI”

> Goal: turn one messy `main.go` into a testable, lint-clean Go module with a green GitHub Actions badge and a `v0.1.0` tag.

---

## 📝 Tasks & Steps

| # | Task | Hint / Command |
|---|------|----------------|
|1|**Create a Go module**|`go mod init github.com/<you>/go-camp` → `go mod tidy`|
|2|**Refactor** – move handler logic into a function `greet.Greet(name)` inside `internal/greet/`.|Separate business logic from `cmd/app/main.go`.|
|3|**Check the application** | Try to run the application `go run cmd/app/main.go` and check it in your browser at http://localhost:8080. Try adding `name` query parameter. 
|3|**Write table-driven unit tests** for `Greet`.|See Go Wiki “Table-Driven Tests”.|
|4|**Add CI** (`.github/workflows/ci.yml`) running `go vet`, `go test`, `golangci-lint run`.|Use `actions/setup-go@v5`.|
|5|**Fix lint findings** (misspell, unchecked errors, variable shadowing).| https://golangci-lint.run/welcome/install/ |
|6|**Tag semantic version** `v0.1.0` and push.|`git tag v0.1.0 && git push --tags`|

---

## 📚 Useful links

* Go Modules: <https://go.dev/doc/tutorial/create-module>
* Go Tests: <https://go.dev/doc/tutorial/add-a-test>
* Table-Driven Tests: <https://go.dev/wiki/TableDrivenTests>
* GitHub Actions for Go: <https://docs.github.com/actions/automating-builds-and-tests/building-and-testing-go>
* Semantic Versioning: <https://semver.org/>  
  *Go’s twist*: import paths include `/v2`, `/v3`, … for breaking releases.

---

## ✅ Acceptance criteria

* `go vet ./...`, `go test ./...`, and `golangci-lint run ./...` all pass locally **and** in CI.
* Repository contains a working workflow file and a CI badge in `README.md`.
* Code base is structured as:

  ```
  cmd/app/main.go
  internal/greet/greet.go
  internal/greet/greet_test.go
  go.mod
  ```

* Git tag **v0.1.0** exists and is pushed to GitHub.

---

### 💡 Reflection questions

1. Why is the result of `http.ListenAndServe` usually checked?
2. What benefit does a function like `Greet` give compared with inline string formatting?
3. How would CI differ for release tags versus every push?

# Repository Guidelines

## Project Structure & Module Organization

- `cmd/api/main.go`: API entrypoint (Gin) and route wiring (currently serves on `:8080`).
- `internal/modules/`: feature modules organized as `handler` (HTTP) → `service` (business logic) → `repository` (DB access) → `model` (types). Example: `internal/modules/auth/`.
- `internal/database/`: database connection helpers (Postgres DSN via env vars).
- `internal/utils/`: shared helpers (e.g., password hashing/verification).
- `migrations/`: SQL migrations (`YYYYMMDDHHMMSS_name.(up|down).sql`) targeting Postgres.
- `Dockerfile`: multi-stage build for a small runtime image.

## Build, Test, and Development Commands

- Run locally: `go run ./cmd/api` (starts the API on port `8080`).
- Build binary: `go build ./cmd/api` (outputs a platform-default binary in the current directory).
- Format: `go fmt ./...` (apply standard Go formatting).
- Static checks: `go vet ./...` (basic correctness checks).
- Tests: `go test ./...` (no tests are currently present; add `_test.go` as you implement coverage).
- Docker image: `docker build -t trackrr-backend .`
- Docker run: `docker run --rm -p 8080:8080 --env-file .env trackrr-backend`

## Coding Style & Naming Conventions

- Use `gofmt` defaults (tabs for indentation; no manual alignment).
- Keep package names short and lowercase (e.g., `auth`, `database`, `utils`).
- Prefer Go’s “small interfaces” pattern; name interfaces by behavior (`Repository`, `Service`) and constructors as `NewX`.

## Testing Guidelines

- Use Go’s built-in `testing` package with table-driven tests where practical.
- Test files must be named `*_test.go` and live next to the package under test.
- For DB-backed code, favor repository/service unit tests with mocks; keep integration tests explicit and documented.

## Commit & Pull Request Guidelines

- Commit history follows a Conventional Commits-like style (e.g., `feat: ...`, `feat(migrations): ...`); keep messages imperative and scoped when helpful.
- PRs should include: a clear description, how to test (commands + sample request), and any DB/migration notes (new tables, required env vars).

## Configuration & Security Tips

- Required DB env vars: `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`.
- Do not commit secrets; use `.env` locally and ensure production config is injected via environment.

## Agent/CI Note

- If your environment restricts writing to the default Go cache, set `GOCACHE`/`GOMODCACHE` to a repo-local directory before running Go commands.

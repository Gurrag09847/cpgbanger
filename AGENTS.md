# AGENTS.md

## Project overview

Self-hosted Go web server that scrapes product documentation PDFs from company websites
and writes results into Excel files. Frontend is a SvelteKit SPA embedded in the Go binary.

- **Go module name:** `illbruck` (root `go.mod` — name came from a Wails template, kept for history)
- **Frontend:** SvelteKit (`frontend/`) with `@sveltejs/adapter-static`, Tailwind v4, shadcn-svelte (bits-ui)
- **Build:** `go build` compiles Go + embeds `frontend/dist` via `//go:embed all:frontend/dist`

## Commands

| Command | Runs in | Purpose |
|----------|---------|---------|
| `go build -o tremco-server.exe .` | repo root | Build the server binary |
| `npm install` | `frontend/` | Install frontend deps |
| `npm run build` | `frontend/` | Vite production build → `frontend/dist` |
| `npm run dev` | `frontend/` | Vite dev server (frontend only, port 5173) |
| `npm run check` | `frontend/` | TypeScript typecheck (svelte-check) |
| `npm run lint` | `frontend/` | Prettier check |
| `npm run format` | `frontend/` | Prettier write |

## Key architecture

- **`main.go`** — HTTP server entrypoint. Embeds `frontend/dist`, sets up routes, SPA fallback.
- **`server.go`** — All HTTP handlers (upload, fetch, cancel, jobs list, SSE, download).
- **`fetcher.go`** — Core scraping logic. Receives progress callback instead of Wails events.
- **`app.go`** — Job struct, broadcaster (SSE fan-out), `FetchParams`.
- **`db.go`** — SQLite persistence (`modernc.org/sqlite`, pure Go, no CGO). Stores job history.
- **`images.txt`** — Image fetcher code (unmaintained; ignored by Go build).
- **`frontend/`** — SvelteKit SPA built with adapter-static (`fallback: 'index.html'`).

## API endpoints

| Method | Path | Purpose |
|--------|------|---------|
| POST | `/api/upload` | Multipart .xlsx upload → `{file_id, file_name}` |
| POST | `/api/fetch` | Start job (JSON body with `file_id` or `file_path`) → `{job_id}` |
| POST | `/api/cancel` | Cancel running job (`{job_id}`) |
| GET | `/api/jobs` | List all jobs (from SQLite) |
| GET | `/api/jobs/{id}` | Single job status |
| GET | `/api/jobs/{id}/sse` | SSE progress stream |
| GET | `/api/jobs/{id}/download` | Download processed .xlsx file |
| GET | `/*` | SPA fallback → serves `frontend/dist/index.html` |

## Frontend conventions

- **Tabs, not spaces** (Prettier `useTabs: true`)
- Single quotes, no trailing commas, print width 100
- Path alias `@/*` → `./src/lib/*`
- Form validation: `sveltekit-superforms` with `zod4Client` adapter (client-only, no server actions)

## Go quirks

- SQLite uses `modernc.org/sqlite` (CGo-free). DSN is just a file path.
- `excelize` column indices are 0-based for reading but `CoordinatesToCellName` expects 1-based
- Product code extraction always takes `artNum[:5]` (first 5 chars)
- HTTP scrape uses fixed CSS selectors — fragile if upstream sites change their HTML
- Uploaded files stored in `data/uploads/`, SQLite DB in `data/tremco.db`
- Server runs on `:8080` by default; override with `ADDR` env var

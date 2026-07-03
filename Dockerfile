FROM oven/bun:1-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN bun install
COPY frontend/ ./
RUN bun run build

FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o tremco-server .

FROM scratch
COPY --from=backend-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=backend-builder /app/tremco-server /tremco-server
EXPOSE 8080
ENV ADDR=:8080
CMD ["/tremco-server"]

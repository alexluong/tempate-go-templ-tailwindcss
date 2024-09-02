ARG app_env=production

# Fetch
FROM golang:1.23-alpine AS fetch
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

# Build assets
FROM node:18-alpine AS assets
WORKDIR /app
COPY ./web/package*.json .
RUN npm install
COPY ./web .
RUN npm run build

# Generate Templ
FROM ghcr.io/a-h/templ:latest AS templ
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build Go binary
FROM golang:1.23-alpine AS build
COPY --from=templ /app /app
COPY --from=assets /app/dist /app/web/dist
WORKDIR /app
RUN go build -o /app/main cmd/main/main.go

# Set up Deploy
FROM scratch AS env-production
WORKDIR /
COPY --from=build /app/main /main
ENTRYPOINT ["/main"]

# Set up Local Development
FROM build AS env-development
RUN go install github.com/air-verse/air@latest
CMD ["air"]

# Run
FROM env-${app_env}
EXPOSE 8090
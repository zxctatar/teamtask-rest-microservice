FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/app ./cmd/taskservice/main.go

FROM alpine:3 AS runtime
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 44048
ENTRYPOINT [ "./app" ]
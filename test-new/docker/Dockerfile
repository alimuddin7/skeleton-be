# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN apk update && apk add --no-cache alpine-sdk
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o test-new main.go

# Run stage
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"
COPY --from=builder /app/test-new .
COPY --from=builder /app/errorcodes.json .
EXPOSE 7878
ENTRYPOINT [ "/app/test-new" ]

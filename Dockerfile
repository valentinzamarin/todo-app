FROM node:18-alpine AS frontend
COPY web ./web
WORKDIR /web
RUN npm install && npm run build


FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/


COPY --from=builder /app/main .
COPY --from=frontend /static ./static

EXPOSE 8080

CMD ["./main"]

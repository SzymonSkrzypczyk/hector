FROM golang:1.24.11-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hector .

FROM alpine:latest

WORKDIR /app

RUN addgroup -S hector && adduser -S hector -G hector

COPY --from=builder /app/hector .
COPY --from=builder /app/themes ./themes
COPY --from=builder /app/examples ./examples

RUN chown -R hector:hector /app

USER hector

ENTRYPOINT ["./hector"]

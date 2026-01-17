FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hector .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/hector .
COPY --from=builder /app/themes ./themes
COPY --from=builder /app/examples ./examples

ENTRYPOINT ["./hector"]

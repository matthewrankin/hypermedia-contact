FROM golang:1.21-alpine as builder
WORKDIR /app
COPY go.mod go.sum* .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o bin/contacts ./cmd/web

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bin/contacts .
COPY --from=builder /app/ui/ ./ui/
CMD [ "./contacts", "-port", "4200", "-env", "development"]

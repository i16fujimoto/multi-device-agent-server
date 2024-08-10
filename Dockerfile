FROM golang:1.22.4 AS builder

WORKDIR /app

COPY . .
RUN go build -o main ./cmd/main.go


FROM gcr.io/distroless/base-debian12 as runner

COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]

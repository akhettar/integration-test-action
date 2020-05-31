FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -v -o readiness-check .

FROM alpine:latest

COPY --from=builder /app/readiness-check /readiness-check

ENTRYPOINT ["/readiness-check"]

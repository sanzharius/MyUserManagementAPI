FROM golang:1.20-alpine3.17 AS build-stage

WORKDIR /app

COPY ./../../go.mod ./../../go.sum ./
RUN go mod download

COPY ./../../cmd/ ./cmd/
COPY ./../../internal/ ./internal/
COPY ./../../configs/ ./configs/
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/usermanager

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/configs/.env /configs/
COPY --from=build-stage /app/app /


ENTRYPOINT ["/app"]

FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o app

FROM alpine:3.19 AS release

WORKDIR /

# Create log directory and set permissions before switching to nonroot user
RUN mkdir -p /tmp/app-logs && \
    chown -R nobody:nobody /tmp/app-logs && \
    chmod 777 /tmp/app-logs

COPY --from=build /app/app /app

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["/app"]
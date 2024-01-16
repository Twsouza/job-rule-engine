
# Builder stage
FROM golang:1.21 AS builder

ENV GOPROXY http://proxy.golang.org

WORKDIR /bin/

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./application/cmd/server.go

# Final stage
FROM alpine:3.16.7

RUN apk add --no-cache ca-certificates

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 3000

CMD exec /bin/app


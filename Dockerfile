# ---------
# Build stage
# ---------
FROM golang:1.23.8-alpine AS builder
WORKDIR /go/monogram

# Copy the Go project directory (adjust the context below if needed)
COPY go/monogram/ .

# Install essential tools (curl is used to install just)
RUN apk update && apk add --no-cache curl

# Install Just from its install script
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | sh -s -- --to /usr/local/bin

# Run the Justfile recipe to build the monogram executable
RUN just build

# ----------
# Runtime stage
# ----------
FROM alpine:3.18
WORKDIR /app

# Copy the monogram binary from the builder stage
COPY --from=builder /go/monogram/monogram /app/monogram

# Ensure the binary is executable
RUN chmod +x /app/monogram

# Set entrypoint to allow arguments to pass through
ENTRYPOINT ["/app/monogram"]
CMD []

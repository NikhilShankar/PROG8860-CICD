# Stage 1: Build stage
# Use specific Go version with Alpine for minimal size
FROM golang:1.21-alpine3.18 AS builder

# Set working directory
WORKDIR /app

# Copy go module files first (better layer caching)
COPY go.mod ./

# Copy source code
COPY main.go ./

# Build the application
# CGO_ENABLED=0 creates a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Final runtime stage
# Use minimal scratch or alpine image
FROM alpine:3.18

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1001 appgroup && \
    adduser -D -u 1001 -G appgroup appuser

# Set working directory
WORKDIR /app

# Copy only the compiled binary from builder stage
COPY --from=builder /app/main .

# Create data directory for volumes
RUN mkdir -p /app/data && chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Set environment variable
ENV PORT=8080

# Expose application port
EXPOSE 8080

# Define volume for persistent data
VOLUME ["/app/data"]

# Run the application
CMD ["./main"]
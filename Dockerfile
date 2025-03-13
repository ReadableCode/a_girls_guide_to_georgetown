# Stage 1: Build the Go backend
FROM golang:1.22.4-alpine AS build

WORKDIR /app

COPY backend/ backend/

WORKDIR /app/backend
RUN go mod tidy
RUN go build -o server

# Stage 2: Create the final image
FROM alpine:latest
WORKDIR /app

# Copy the built Go server from the previous stage
COPY --from=build /app/backend/server /app/backend/server

# Ensure the volume directories exist
RUN mkdir -p /app/logs
RUN mkdir -p /app/frontend

# Declare volumes for frontend and logs
VOLUME ["/app/frontend", "/app/logs"]

# Set the working directory to the backend
WORKDIR /app/backend

# Expose the port the server listens on
EXPOSE 8504

# Run the server
CMD ["./server"]

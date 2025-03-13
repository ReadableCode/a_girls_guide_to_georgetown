# Stage 1: Build the Go binary
FROM golang:1.22.4-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy backend and frontend files
COPY backend/ backend/
COPY frontend/ frontend/

# Change to the backend directory
WORKDIR /app/backend

# Download dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o server

# Stage 2: Create the final container
FROM alpine:latest

# Set working directory inside the final container
WORKDIR /root/

# Copy the built Go binary and frontend files from the builder stage
COPY --from=build /app/backend/server .
COPY --from=build /app/frontend /frontend

# Expose port 3001 for the application
EXPOSE 3001

# Run the server
CMD ["./server"]

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

# Copy the frontend separately
COPY frontend/ frontend/

WORKDIR /app/backend

EXPOSE 3001
CMD ["./server"]
# a_girls_guide_to_georgetown

## Hosting Locally

```bash
cd backend
go mod init portfolio  # if not already initialized
go mod tidy
go run main.go
```

## Hosting Locally with Docker

```bash
docker build -t a-girls-guide-to-georgetown .
docker run -it -p 3001:3001 -v "$PWD/frontend:/app/frontend" a-girls-guide-to-georgetown
```

## DIagnosing inside the container

```bash
docker run --rm -it a-girls-guide-to-georgetown sh
```

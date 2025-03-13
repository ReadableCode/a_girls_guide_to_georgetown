# a_girls_guide_to_georgetown

## Hosting Locally

```bash
cd backend
go mod init portfolio
go mod tidy
go run main.go
```

## Hosting Locally with Docker

```bash
docker build -t a-girls-guide-to-georgetown .
docker run -it -p 3001:3001 a-girls-guide-to-georgetown
```

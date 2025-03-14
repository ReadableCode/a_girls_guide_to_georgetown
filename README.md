# a_girls_guide_to_georgetown

## Initializing the Project

```bash
cd backend
go mod init portfolio
go mod tidy
```

## Formatting HTML files

```bash
# cd to the root of the project
npx prettier --write "frontend/templates/**/*.html"  
```

## Hosting Locally

```bash
cd backend
go mod tidy
go run main.go
```

## Hosting Locally with Docker

```bash
docker build -t a-girls-guide-to-georgetown .
docker run -it -p 8504:8504 -v "$PWD/frontend:/app/frontend" -v "$PWD/logs:/app/logs" a-girls-guide-to-georgetown
```

## Diagnosing inside the container

```bash
docker run --rm -it a-girls-guide-to-georgetown sh
```

## Bashing into the running container

```bash
sudo docker exec -it a-girls-guide-to-georgetown sh

```

## Documentation on reloads for page

[Docker Docs](../Docker/README.md)

### Rebuilding after backend changes in repo

```bash
sshelite
cd /home/jason/GitHub/Docker
sudo docker-compose -f docker_compose_projects.yaml build a_girls_guide_to_georgetown
sudo docker-compose -f docker_compose_projects.yaml up -d
```

### Rebuilding after changes to Dockerfile or Docker Compose

```bash
sshelite
cd /home/jason/GitHub/Docker
sudo systemctl stop docker_compose_projects.service
sudo docker-compose -f docker_compose_projects.yaml down --remove-orphans
sudo docker image prune -f
sudo docker-compose -f docker_compose_projects.yaml up --build --force-recreate -d
sudo systemctl start docker_compose_projects.service
```

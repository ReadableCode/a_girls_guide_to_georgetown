# A Girl's Guide to Georgetown

This repository contains the full codebase for a website designed to allow high school students in Georgetown, Texas, to manage the frontend independently while relying on a structured backend I built and maintain.

## Backend

- Built with **Go**, using **Fiber** for efficient, lightweight web serving.
- Dynamically serves pages without requiring container rebuilds, utilizing **Docker volumes**.
- Provides a structured API while handling page generation on demand.
- Designed for **flexibility**, allowing the frontend to evolve without requiring backend changes.
- Supports **customization** while ensuring stability and scalability.

## Frontend

- Managed entirely by the students, using **HTML, CSS, and JavaScript**.
- The frontend code is included in this repository, but **it is not my work**.
- The structure allows students to focus on frontend development and content updates without needing backend expertise.

## Hosting and Deployment

- Hosted on a **Linux (Ubuntu) server** running **Docker Compose**.
- Backend is containerized, built from the Go source code, and deployed via Docker Compose.
- Reverse-proxied through **SWAG (Secure Web Application Gateway)**.
- **Cloudflare** manages domain routing, with **DuckDNS** handling dynamic DNS updates.
- Uses **Docker volumes** to allow real-time updates to pages without restarting containers.

This project is intended to support student-led web development while maintaining a solid backend foundation.

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
# Windows
docker run -it -p 8504:8504 -v "$PWD/frontend:/app/frontend" -v "$PWD/logs:/app/logs" a-girls-guide-to-georgetown
# Linux
docker run -it -p 8504:8504 -v "$(pwd)/frontend:/app/frontend" -v "$(pwd)/logs:/app/logs" a-girls-guide-to-georgetown
```

## Diagnosing inside the container

```bash
# Windows
docker run -it -p 8504:8504 -v "$PWD/frontend:/app/frontend" -v "$PWD/logs:/app/logs" a-girls-guide-to-georgetown sh
# Linux
docker run -it -p 8504:8504 -v "$(pwd)/frontend:/app/frontend" -v "$(pwd)/logs:/app/logs" a-girls-guide-to-georgetown sh
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

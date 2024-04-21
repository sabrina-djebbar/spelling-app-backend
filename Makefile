GOCMD=go

docker-build:
	docker build -f Dockerfile . -t spelling-app:prod

docker-compose:
	docker-compose up -d

init-local-postgres:
	docker run --name pg-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
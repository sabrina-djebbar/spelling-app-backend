GOCMD=go

docker-build:
	docker build -f Dockerfile . -t spelling-app:prod

docker-compose:
	docker-compose up -d
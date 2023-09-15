handler:
	@go run .\client\main.go

processor:
	@go run .\server\main.go 

pg:
	@docker run --name aggr -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=pass -p 5432:5432 -d postgres

pgdb:
	@docker exec

network:
	@docker network create aggr-network

run:
	@docker compose up --build -d
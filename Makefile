all:
	sudo docker-compose up --build

up:
	psql -U postgres -c "CREATE DATABASE POCHTA_RUSSIA;" || true
	go run cmd/main/main.go
down:
	psql -U postgres -c "DROP DATABASE POCHTA_RUSSIA;" || true

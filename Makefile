postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15.2

psql:
	docker exec -it postgres15 psql

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat-db

dropDb:
	docker exec -it postgres15 dropdb go-chat-db

migrateUp:	
	migrate -path db/migrations -database "postgresql://postgres:je8hNb7eLKmlFkBrq9Gk@containers-us-west-94.railway.app:5795/railway" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgresql://postgres:je8hNb7eLKmlFkBrq9Gk@containers-us-west-94.railway.app:5795/railway" -verbose down

.PHONY: postgresinit  migrateUp migrateDown psql createDb dropDb

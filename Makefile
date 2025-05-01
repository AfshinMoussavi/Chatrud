postgresinit:
	docker run --name postgres_container -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres

postgres:
	docker exec -it postgres_container psql -U postgres

createdb:
	docker exec -it postgres_container createdb --username=postgres --owner=postgres chat_db

dropdb:
	docker exec -it postgres_container dropdb go-chat

migrateup:
	migrate -path models/migrations -database "postgres://postgres:1234@127.0.0.1:5433/chat_db?sslmode=disable" -verbose up

migratedown:
	migrate -path models/migrations -database "postgres://postgres:1234@127.0.0.1:5433/chat_db?sslmode=disable" -verbose down

redisinit:
	docker run --name redis_container -p 6379:6379 redis

rediscli:
	docker exec -it redis_container redis-cli

redislog:
	docker logs -f redis_container

stopredis:
	docker stop redis_container && docker rm redis_container

stoppostgres:
	docker stop postgres_container && docker rm postgres_container

swagger:
	swag init --generalInfo cmd/main.go

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown redisinit rediscli redislog stopredis stoppostgres swagger



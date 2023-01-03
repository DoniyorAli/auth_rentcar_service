migrate create -ext sql -dir ./migrations/postgres -seq -digits 2 create_article_table

migrate -path ./storage/migrations -database 'postgres://admin:qwerty123@localhost:5432/auth_service_db?sslmode=disable' up

migrate -path ./storage/migrations -database 'postgres://admin:qwerty123@localhost:5432/auth_service_db?sslmode=disable' down
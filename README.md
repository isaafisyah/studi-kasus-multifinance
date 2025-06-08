# Record Transaction API

API untuk mencatat transaksi, dibuat dengan Golang + Gin.

# docker

- docker-compose build
- docker-compose up -d
- docker logs multifinance
- docker exec -ti multifinance bash
- docker-compose stop (jika ingin stop docker)

# migration

- docker-compose run migrate
- docker-compose run migrate-down (untuk rollback)

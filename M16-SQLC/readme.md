# Configuração do Golang Migrate migrate (https://github.com/golang-migrate/migrate)

1. Instalar o command line (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```bash
# MACOS
brew install golang-migrate
```

# Manipulação de migrations

```bash
# Criar
migrate create -ext=sql -dir=sql/migrations -seq init

# Executar Migrations
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```

# Testanto banco de dados

```bash
docker-compose exec mysql bash

mysql -uroot -p courses

show tables;
```

# Configuração SQLC (https://sqlc.dev/)

1. Instalação command line (https://docs.sqlc.dev/en/latest/overview/install.html)

```bash
# MacOS
brew install sqlc
```

# Manipulando SQLC

```bash
sqlc generate
```

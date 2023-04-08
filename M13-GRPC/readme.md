# Pré requisitos

1. Instalar o compilador do protocol buffers (https://grpc.io/docs/protoc-installation/)

```bash
# Linux
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+

# Mac
brew install protobuf
protoc --version  # Ensure compiler version is 3+
```

2. Instalar plugins GRPC (https://grpc.io/docs/languages/go/quickstart/)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

# Compilação arquivo

```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.prot
```

# Instalação client Evans (client de coneção grpc)

1. verificar o seguinte link https://github.com/ktr0731/evans
2. Após instalado o evans iniciar o server e conectar.

```bash
# Iniciar o server
go run cmd/grpcServer/main.go

# Conectar no server iniciado
evans -r repl

# Selecionar o serviço
service CategoryService

# Criar registro
call CreateCategoryX
```

# Criando banco

```bash
sqlite3 db.sqlite

# criar tabelas
create table categories (id string, name string, description string);

create table courses (id string, name string, description string, category_id string);
```

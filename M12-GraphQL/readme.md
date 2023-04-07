# Criado esqueleto com a lib

- https://gqlgen.com/getting-started/

```bash
# Criando esqueleto server
go run github.com/99designs/gqlgen init

# Atualizar shema (Vai ler alteraçoes nos arquivos e atualizar)
go run github.com/99designs/gqlgen generate
```

# Criando banco

```bash
sqlite3 data.db

# criara tabela
create table categories (id string, name string, description string);
```

# Graphql

```bash
mutation createCategory {
  createCategory(input: {
    name: "Tecnologia",
    description: "Cursos tecnologia"
  }){
    id, name, description
  }
}
```

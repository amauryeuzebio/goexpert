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

# criar tabelas
create table categories (id string, name string, description string);

create table courses (id string, name string, description string, category_id string);
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

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
    }
  }
}

mutation createCourse {
  createCourse(input: {
    name: "Go",
    description: "Curso GoLang",
    categoryId: "e3df3a9b-a85e-4897-9b02-523b91cbbc59"
  }){
    id, name, description
  }
}

query queryCourses {
  courses {
    id
    name
    description

  }
}
```

# Configuração cobra-cli (https://github.com/spf13/cobra)

```bash
# Instalação command line do cobra
go install github.com/spf13/cobra-cli@latest

# Iniciar projeto cobra
cobra-cli init

# Adicionar comandos
cobra-cli add ping

# Par rodar o comando criado
go run main.go ping

# Criando comando filhos
cobra-cli add category
cobra-cli add create -p 'categoryCmd'
cobra-cli add list -p 'categoryCmd'
```


# Criando banco

```bash
sqlite3 data.db

# criar tabelas
create table categories (id string, name string, description string);

create table courses (id string, name string, description string, category_id string);
```

# Testar criação categoria 
```bash
go run main.go category create -n=Cat2 -d=Desc2
```


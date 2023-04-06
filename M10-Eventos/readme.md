# Instruções
1. Subir um container do rabbitMQ (docker-composer up -d)
2. Entrar no rabbit http://localhost:15672/ e criar uma fila e fazer o bind com a exchange amq.direct.
3. Executar o consumer (go run cmd/consumer/main.go)
4.  Executar o producer (go run cmd/producer/main.go)

### Esse exemplo o producer vai enviar uma mensagem "Hello World!" e o consumer vai dar um print em tela
## Desafio Clean Architecture

#### Deve-se rodar o docker para subir o serviço de banco de dados.
#### Se necessário, execute as migrations com o Makefile (make migrate).

#### Na pasta cmd/ordersystem execute o comando go run main.go wire_gen.go para subir todos os servidores (gRPC, GraphQL e REST).

## O gRPC server roda na porta 50051

Utilze o Evans (evans -r repl) para fazer as chamadas gRPC com o seguinte comando: 

call ListOrders

## O GraphQL server roda na porta 8080

Acesse http://localhost:8080/ e execute a seguinte query para listar as ordens:

query orders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}

## O Web server roda na porta 8000

Utilize o arquivo api/list_order.http para enviar a requisição de listagem ao servidor.

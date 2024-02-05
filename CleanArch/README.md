# List Orders

## Como Rodar o Código
Clone o repositório:

```bash
git clone https://github.com/danubiobwm/goExpert/tree/main/CleanArch
```
## Navegue até o diretório do projeto:

```bash
cd CleanArch
```

## Execute o código Go:

```bash
docker compose up --build -d

# Isso irá expor os serviços necessários para que a aplicação rode, como o mysql por exemplo.
```


```bash
go mod tidy
cd cmd/ordersystem
go run main.go wire_gen.go

# Isso irá expor os endpoints de rest, grpc e graphql necessários para as validações.

# Starting web server on port :8000
# Starting gRPC server on port 50051
# Starting GraphQL server on port 8080
```


## Exemplo de Uso

### REST
No arquivo `api/create_order.http`, execute ambos os exemplos para validar.

```
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "id":"a",
    "price": 100.5,
    "tax": 0.5
}

###
No arquivo `api/list_orders.http`,

GET http://localhost:8000/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

### GRAPHQL

Acesse no seu navegador a rota http://localhost:8080

E execute as seguintes queries para validação

```graphql
mutation createOrder {
  createOrder(input: {id: "ccc", Price: 10.20, Tax:2.0 }) {
    id,
    Price,
    Tax
  }
}
```

```graphql
query listOrders {
  listOrders {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

### GRPC
Com o evans: https://github.com/ktr0731/evans, Execute os seguintes comandos:

```bash
evans -r repl
```

```bash
package pb
```

```bash
service OrderService
```

```bash
call ListOrders
# ou
call CreateOrder
```

## Exemplo de saída bem-sucedida:

### REST

```json
[
  {
    "ID": "11",
    "Price": 10.2,
    "Tax": 2,
    "FinalPrice": 12.2
  },
  {
    "ID": "2",
    "Price": 12,
    "Tax": 1,
    "FinalPrice": 13
  },
]
```

### GRAPHQL

```json
{
  "data": {
    "listOrders": [
      {
        "id": "11",
        "Price": 10.2,
        "Tax": 2
      },
      {
        "id": "2",
        "Price": 12,
        "Tax": 1
      },
    ]
  }
}
```


### GRPC

```json
{
  "orders": [
    {
      "finalPrice": 12.2,
      "id": "11",
      "price": 10.2,
      "tax": 2
    },
    {
      "finalPrice": 13,
      "id": "2",
      "price": 12,
      "tax": 1
    }
  ]
}
```
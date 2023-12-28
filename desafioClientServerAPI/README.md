# GoExpert: Desafio 01

Desafio para aplicar o aprendizado sobre webserver http, contextos, banco de dados e manipulação de arquivos com Go.

## Client Server API

Projeto com duas aplicações desenvolvidas em Golang, `client` e `server`.

Os requisitos são:

- O `client` deverá realizar uma requisição HTTP no `server` solicitando a cotação do dólar.

- O `server` deverá consumir a API contendo o câmbio de Dólar e Real no endereço `https://economia.awesomeapi.com.br/json/last/USD-BRL`, e em seguida deverá retornar no formato JSON o resultado para o `client`.

- O `server` deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms, e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

- O `client` precisará receber do `server` apenas o valor atual do câmbio (campo "bid" do JSON). O `client` terá um timeout máximo de 300ms para receber o resultado.

- O `client` terá que salvar a cotação atual em um arquivo `cotacao.txt` no formato: Dólar: {valor}.

- O endpoint necessário gerado pelo `server` para este desafio será: `/cotacao` e a porta a ser utilizada pelo servidor HTTP será a `8080`.

### Iniciar aplicação `server`

Abrir uma janela de terminal e executar:

```
cd server && go run server.go
```

### Iniciar aplicação `client`

Abrir uma segunda janela de terminal e executar:

```
cd client && go run client.go
```


# O projeto de equipe de marketing para criação de campanhas

## Como rodar o projeto

### 1. Instale o [Docker](https://docs.docker.com/install/) e o [Docker Compose](https://docs.docker.com/compose/install/)

### 2. Clone o projeto

### 3. Rode o comando `docker-compose up` na raiz do projeto

### 4. Acesse o projeto em `http://localhost:8080`

### Envio de e-mails em massa para clientes

### Dependências utilizadas no projeto

#### - Para testes
#### - [Testify](https://github.com/stretchr/testify)

#### - Para Codar
#### - Gerador de ID globalmente exclusivo
#### - [xid](https://github.com/rs/xid)

### Para a fazer validação ulitizamos a biblioteca 
#### - [validator](https://pkg.go.dev/github.com/go-playground/validator/v10)

### Para criar dados fake utilizamos a biblioteca 
#### - [faker](https://pkg.go.dev/github.com/jaswdr/faker)

### Para construir serviços HTTP em Go usaremos a biblioteca [Chi](https://github.com/go-chi/chi)

- O go-chi/chi é um roteador leve e componível para construir serviços HTTP em Go, 
especialmente útil para grandes serviços de API REST. 
Ele é construído no novo pacote de contexto introduzido no Go 1.7 
e busca um design elegante e confortável para escrever servidores de API REST. 
Suas principais considerações de design incluem estrutura do projeto, manutenibilidade, produtividade do desenvolvedor e desconstrução de um grande sistema em partes menores. O roteador principal é pequeno, 
mas inclui subpacotes úteis/opcionais como middleware, render e docgen.

### O pacote [**Render**](https://github.com/go-chi/render) ajuda a gerenciar cargas úteis de solicitação/resposta HTTP.
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


#### [**Air**](https://github.com/cosmtrek/air) é uma ferramenta fantástica para desenvolvedores Go, proporcionando **recarga ao vivo** durante o desenvolvimento. Quando você está trabalhando em aplicativos Go, o **Air** recompila e executa automaticamente seu código toda vez que você faz alterações. É como ter um assistente mágico que observa seu código e o atualiza em tempo real! 🚀

Aqui estão algumas características-chave do **Air**:

1. **Saída de Log Colorida**: Desfrute de um display de log visualmente atraente.
2. **Comandos de Compilação Personalizáveis**: Ajuste seu processo de compilação conforme suas necessidades.
3. **Exclusão de Subdiretórios**: Especifique quais diretórios excluir da monitoração.
4. **Monitoramento Dinâmico de Diretórios**: Adicione novos diretórios para monitorar mesmo depois que o **Air** já foi iniciado.
5. **Processo de Compilação Aprimorado**: Uma experiência de compilação melhor.
6. **Sobrescrever Configuração**: Substitua as configurações de configuração por meio de argumentos de linha de comando.

Lembre-se de que o **Air** é tudo sobre conveniência de desenvolvimento e **recarga ao vivo**. Não é destinado a implantar código de produção em produção. Portanto, se você está trabalhando em um projeto Go, experimente o **Air** e deixe-o lidar com a recarga enquanto você se concentra no seu código! 😊

Você pode instalar o **Air** usando qualquer um destes métodos:

- **Recomendado**: Via `go install` (requer Go 1.18 ou superior):
    ```bash
    go install github.com/cosmtrek/air@latest
    ```

- Usando o script `install.sh` fornecido:
    ```bash
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
    ```

Saiba mais sobre o **Air** em seu [repositório GitHub](https://github.com/cosmtrek/air) e comece a aproveitar a magia da recarga ao vivo! ✨👟
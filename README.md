# CV Manager

## Descrição

O **CV Manager** é um projeto em desenvolvimento que visa gerenciar currículos de forma eficiente, utilizando tecnologias modernas como GraphQL, MongoDB e Go. Este repositório ainda está em fase inicial e pode conter funcionalidades incompletas ou sujeitas a alterações.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal do projeto.
  - <a href="https://go.dev/doc/" target="_blank">Documentação do Go</a>
- **GraphQL**: Para criação de APIs eficientes e flexíveis.
  - <a href="https://graphql.org/learn/" target="_blank">Documentação do GraphQL</a>
  - <a href="https://gqlgen.com/" target="_blank">GQLGen</a>: Ferramenta para gerar código GraphQL em Go.
- **MongoDB**: Banco de dados NoSQL para armazenamento de dados.
  - <a href="https://www.mongodb.com/docs/" target="_blank">Documentação do MongoDB</a>
- **Docker**: Para containerização do ambiente de desenvolvimento.
  - <a href="https://docs.docker.com/" target="_blank">Documentação do Docker</a>

## Estrutura do Projeto

- **`docker-compose.yml`**: Configuração do serviço MongoDB.
- **`go.mod` e `go.sum`**: Gerenciamento de dependências do Go.
- **`server.go`**: Ponto de entrada principal do servidor.
- **`config/`**: Configurações do projeto.
- **`db/`**: Integração com o banco de dados MongoDB.
- **`generated/`**: Código gerado automaticamente pelo GQLGen.
- **`model/`**: Modelos de dados.
- **`resolvers/`**: Resolvers para as queries e mutations do GraphQL.
- **`schemas/`**: Definições do esquema GraphQL.

## Como Executar

1. Certifique-se de ter o Docker e o Docker Compose instalados.
2. Suba o ambiente com o comando:
   ```bash
   docker compose up
   ```
3. Gere os esquemas GraphQL com o comando:
   ```bash
   make generate
   ```
4. Execute a aplicação com o comando:
   ```bash
   make run
   ```
5. O servidor estará disponível na porta configurada (padrão: `8080`).

## Comandos Disponíveis no Makefile

O projeto inclui um `makefile` com os seguintes comandos para facilitar o desenvolvimento:

- **`make build`**:

  - Compila o servidor Go e gera o binário na pasta `bin/`.
  - Útil para criar o executável do servidor.

- **`make run`**:

  - Compila o servidor (executa `make build`) e o inicia com as variáveis de ambiente configuradas.
  - Útil para rodar o servidor localmente.

- **`make generate`**:

  - Gera os esquemas GraphQL usando o `gqlgen`.
  - Útil para atualizar os resolvers e esquemas GraphQL após alterações.

- **`make clean`**:
  - Remove a pasta `bin/` e limpa os binários gerados.
  - Útil para limpar o ambiente de desenvolvimento.

O comando padrão é o `make run`, que compila e executa o servidor automaticamente.

## Aviso

Este projeto ainda **não está finalizado**. Algumas funcionalidades podem estar incompletas ou sujeitas a alterações.

# API Rede Social

Projeto de uma API REST para gestão de usuários e publicações em uma rede social, desenvolvida em Go. O objetivo principal é aplicar e consolidar conceitos de desenvolvimento backend, autenticação, persistência de dados, organização em camadas e criação de endpoints REST.

## Visão geral

A aplicação oferece funcionalidades para:

- cadastro de usuários;
- autenticação com token JWT;
- criação, listagem, edição e remoção de publicações;
- proteção de rotas privadas por middleware;
- manutenção de dados em banco relacional;
- estruturação modular do código para facilitar manutenção e evolução.

## Tecnologias utilizadas

- Go
- Gin Web Framework
- MySQL
- JWT (JSON Web Token)
- GORM ou acesso manual ao banco, conforme a implementação do projeto
- Arquitetura por camadas: controllers, models, repositories, router e middlewares

## Estrutura do projeto

```text
API-rede-social/
├── main.go
├── go.mod
├── README.md
├── sql/
│   ├── dados.sql
│   └── sql.sql
├── src/
│   ├── authentication/
│   ├── config/
│   ├── controllers/
│   ├── db/
│   ├── middlewares/
│   ├── models/
│   ├── repositories/
│   ├── respostas/
│   ├── router/
│   └── security/
└── ...
```

### Descrição dos principais diretórios

- `main.go` — ponto de entrada da aplicação.
- `src/controllers` — lógica dos endpoints da API.
- `src/models` — estruturas de dados do sistema.
- `src/repositories` — acesso e manipulação de dados no banco.
- `src/router` — definição das rotas e agrupamento de endpoints.
- `src/middlewares` — validação de autenticação e demais interceptadores.
- `src/security` — criptografia e regras de segurança.
- `src/config` — configuração geral da aplicação.
- `src/db` — conexão com o banco de dados.

## Requisitos

Antes de executar o projeto, verifique se você possui:

- Go instalado na máquina
- acesso a um banco MySQL
- Git para clonar o repositório
- ferramenta de testes ou cliente HTTP, como Postman ou cURL

## Como executar o projeto

### 1. Clone o repositório

```bash
git clone <url-do-repositorio>
cd API-rede-social
```

### 2. Instale as dependências

```bash
go mod download
```

### 3. Configure o banco de dados

Ajuste as variáveis de ambiente ou os valores de conexão no arquivo de configuração correspondente, incluindo:

- host do banco;
- porta;
- usuário;
- senha;
- nome do banco;
- timezone, se necessário.

### 4. Execute a aplicação

```bash
go run main.go
```

A API ficará disponível em um endereço local, normalmente em:

```text
http://localhost:8080
```

## Endpoints principais

### Autenticação

| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | `/login` | Realiza login e retorna o token de acesso |

### Usuários

| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | `/usuarios` | Cadastra um novo usuário |
| GET | `/usuarios` | Lista usuários |
| GET | `/usuarios/:id` | Busca um usuário por ID |
| PUT | `/usuarios/:id` | Atualiza dados de um usuário |
| DELETE | `/usuarios/:id` | Remove um usuário |

### Publicações

| Método | Endpoint | Descrição |
| --- | --- | --- |
| POST | `/publicacoes` | Cria uma nova publicação |
| GET | `/publicacoes` | Lista as publicações |
| GET | `/publicacoes/:id` | Busca uma publicação por ID |
| PUT | `/publicacoes/:id` | Atualiza uma publicação |
| DELETE | `/publicacoes/:id` | Remove uma publicação |

## Exemplos de requisições em português

### 1. Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@email.com",
    "senha": "123456"
  }'
```

Resposta esperada:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 2. Cadastro de usuário

```bash
curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com",
    "senha": "123456"
  }'
```

### 3. Listar publicações

```bash
curl -X GET http://localhost:8080/publicacoes \
  -H "Authorization: Bearer <token>"
```

### 4. Criar publicação

```bash
curl -X POST http://localhost:8080/publicacoes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "texto": "Olá, mundo! Esta é minha primeira publicação.",
    "usuario_id": 1
  }'
```

## Request examples in English

### 1. Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@email.com",
    "password": "123456"
  }'
```

Expected response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 2. Register a user

```bash
curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "john@email.com",
    "password": "123456"
  }'
```

### 3. List posts

```bash
curl -X GET http://localhost:8080/publicacoes \
  -H "Authorization: Bearer <token>"
```

### 4. Create a post

```bash
curl -X POST http://localhost:8080/publicacoes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "text": "Hello, world! This is my first post.",
    "user_id": 1
  }'
```

## Observações

- Este projeto foi criado com finalidade de estudo e prática de desenvolvimento backend em Go.
- A estrutura está organizada para facilitar crescimento e manutenção.
- Em um ambiente de produção, recomenda-se reforçar a segurança, aplicar validações mais rigorosas, usar boas práticas de criptografia e configurar variáveis de ambiente corretamente.
- O projeto pode servir como base para evoluções como comentários, curtidas, seguidores, upload de imagens e melhorias na arquitetura.

## English overview

This project is a social network API built in Go with a modular structure for handling users and posts. It was designed to strengthen practical knowledge in backend development, API design, authentication, database access, and layered application organization.

The API exposes REST endpoints for user registration and post management, using JWT authentication to restrict access to protected routes. The codebase is organized to be easy to extend and suitable as a starting point for more complete social media features.

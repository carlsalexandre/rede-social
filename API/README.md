# API — Rede Social

> 🇧🇷 [Português](#-português) · 🇺🇸 [English](#-english)

---

## 🇧🇷 Português

API REST em Go para uma rede social de estudo, responsável por toda a regra de negócio e persistência de dados: usuários, autenticação, publicações, curtidas e relação de seguir/seguidores. É consumida pelo projeto [`WebApp`](../WebApp/README.md), que fica na raiz do repositório.

### Visão geral

- Cadastro de usuário e login com token **JWT**.
- CRUD de publicações, com curtir/descurtir.
- Busca de usuários por nome/nick e de publicações por conteúdo (`LIKE`).
- Seguir, deixar de seguir, listar seguidores e seguindo.
- Atualização de dados do usuário (nome, nick, e-mail), com **regra de 3 meses** entre atualizações.
- Atualização de senha (exige a senha atual) e exclusão de conta.
- Todas as rotas privadas protegidas por middleware de autenticação (JWT no header `Authorization`).

### Tecnologias

- [Go](https://go.dev/)
- [gorilla/mux](https://github.com/gorilla/mux) — roteamento HTTP
- [MySQL](https://www.mysql.com/) via [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) (`database/sql` puro, sem ORM)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) — geração e validação de token
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — hash de senha
- [badoux/checkmail](https://github.com/badoux/checkmail) — validação de e-mail
- [joho/godotenv](https://github.com/joho/godotenv) — variáveis de ambiente

### Estrutura do projeto

```text
API/
├── main.go
├── go.mod
├── sql/
│   ├── sql.sql          # cria o banco e as tabelas
│   └── dados.sql        # dados de exemplo (opcional)
└── src/
    ├── authentication/  # geração/validação do JWT
    ├── config/          # leitura do .env
    ├── controllers/      # lógica de cada endpoint
    ├── db/               # conexão com o MySQL
    ├── middlewares/      # autenticação e logger das requisições
    ├── models/           # Usuario, Publicacao, Senha...
    ├── repositories/     # queries SQL
    ├── respostas/        # padronização das respostas JSON
    ├── router/           # definição e agrupamento das rotas
    └── security/         # hash e verificação de senha
```

### Como executar

**1. Pré-requisitos:** Go instalado e um MySQL acessível.

**2. Banco de dados** — rode o script (isso cria o banco `redesocial` do zero):

```bash
mysql -u root -p < sql/sql.sql
mysql -u root -p redesocial < sql/dados.sql   # opcional, dados de teste
```

**3. Variáveis de ambiente** — crie um arquivo `.env` na pasta `API/`:

```env
API_PORT=5000
DB_USUARIO=root
DB_SENHA=sua_senha
DB_NOME=redesocial
SECRET_KEY=uma_chave_secreta_qualquer
```

**4. Instale as dependências e rode:**

```bash
go mod download
go run main.go
```

A API sobe em `http://localhost:5000` (ou na porta definida em `API_PORT`).

### Endpoints

Rotas marcadas com 🔒 exigem o header `Authorization: Bearer <token>`.

#### Autenticação

| Método | Rota | Descrição |
| --- | --- | --- |
| POST | `/login` | Autentica e retorna o token JWT |

#### Usuários

| Método | Rota | Descrição |
| --- | --- | --- |
| POST | `/usuarios` | Cadastra um novo usuário |
| GET | `/usuarios?usuario=` 🔒 | Busca usuários por nome ou nick |
| GET | `/usuarios/{usuarioId}` 🔒 | Busca um usuário por ID |
| PUT | `/usuarios/{usuarioId}` 🔒 | Atualiza nome/nick/e-mail (bloqueado por 3 meses após a última atualização) |
| DELETE | `/usuarios/{usuarioId}` 🔒 | Exclui a conta |
| POST | `/usuarios/{usuarioId}/atualizar-senha` 🔒 | Troca a senha (exige a senha atual) |
| POST | `/usuarios/{usuarioId}/seguir` 🔒 | Passa a seguir o usuário |
| POST | `/usuarios/{usuarioId}/parar-de-seguir` 🔒 | Deixa de seguir |
| GET | `/usuarios/{usuarioId}/seguidores` 🔒 | Lista quem segue o usuário |
| GET | `/usuarios/{usuarioId}/seguindo` 🔒 | Lista quem o usuário segue |

#### Publicações

| Método | Rota | Descrição |
| --- | --- | --- |
| POST | `/publicacoes` 🔒 | Cria uma publicação |
| GET | `/publicacoes` 🔒 | Lista publicações do usuário e de quem ele segue |
| GET | `/publicacoes/buscar?termo=` 🔒 | Busca publicações por trecho do conteúdo |
| GET | `/publicacoes/{publicacaoId}` 🔒 | Busca uma publicação por ID |
| PUT | `/publicacoes/{publicacaoId}` 🔒 | Atualiza uma publicação (só o autor) |
| DELETE | `/publicacoes/{publicacaoId}` 🔒 | Exclui uma publicação (só o autor) |
| GET | `/usuarios/{usuarioId}/publicacoes` 🔒 | Lista publicações de um usuário específico |
| POST | `/publicacoes/{publicacaoId}/curtir` 🔒 | Curte a publicação |
| POST | `/publicacoes/{publicacaoId}/descurtir` 🔒 | Descurte a publicação |

> ⚠️ `/publicacoes/buscar` é registrada **antes** de `/publicacoes/{publicacaoId}` no código — se a ordem for invertida, o gorilla/mux tenta interpretar "buscar" como um ID.

### Exemplos de requisição

```bash
# Login
curl -X POST http://localhost:5000/login \
  -H "Content-Type: application/json" \
  -d '{"email": "usuario@email.com", "senha": "123456"}'

# Criar publicação (autenticado)
curl -X POST http://localhost:5000/publicacoes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"conteudo": "Olá, mundo!"}'

# Buscar publicações por conteúdo
curl -X GET "http://localhost:5000/publicacoes/buscar?termo=ola" \
  -H "Authorization: Bearer <token>"
```

### Banco de dados

Três tabelas: `usuarios`, `publicacoes` e `seguidores` (tabela de relação usuário↔seguidor). Todas as chaves estrangeiras usam `ON DELETE CASCADE` — excluir um usuário já apaga suas publicações e conexões automaticamente. O campo `usuarios.atualizado_em` guarda quando os dados do usuário foram editados pela última vez, usado na regra dos 3 meses.

### Observações

Projeto de estudo, feito para praticar Go no backend: rotas REST, autenticação com JWT, acesso ao banco sem ORM, arquitetura em camadas (controller → repository → banco). Em produção, valeria reforçar validações, usar variáveis de ambiente com mais cuidado e revisar as políticas de CORS.

---

## 🇺🇸 English

Go REST API for a study social network project, responsible for all business logic and data persistence: users, authentication, posts, likes and the follow/followers relationship. It's consumed by the [`WebApp`](../WebApp/README.md) project at the root of this repository.

### Overview

- User registration and login with a **JWT** token.
- Full CRUD for posts, plus like/unlike.
- Search users by name/nick and posts by content (`LIKE`).
- Follow, unfollow, list followers and following.
- Update user data (name, nick, email), with a **3-month rule** between updates.
- Change password (requires current password) and delete account.
- All private routes are protected by an authentication middleware (JWT in the `Authorization` header).

### Tech stack

- [Go](https://go.dev/)
- [gorilla/mux](https://github.com/gorilla/mux) — HTTP routing
- [MySQL](https://www.mysql.com/) via [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) (plain `database/sql`, no ORM)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) — token issuing and validation
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — password hashing
- [badoux/checkmail](https://github.com/badoux/checkmail) — email validation
- [joho/godotenv](https://github.com/joho/godotenv) — environment variables

### Project structure

```text
API/
├── main.go
├── go.mod
├── sql/
│   ├── sql.sql          # creates the database and tables
│   └── dados.sql        # sample data (optional)
└── src/
    ├── authentication/  # JWT generation/validation
    ├── config/          # reads the .env file
    ├── controllers/      # each endpoint's logic
    ├── db/               # MySQL connection
    ├── middlewares/      # auth and request logger
    ├── models/           # Usuario, Publicacao, Senha...
    ├── repositories/     # SQL queries
    ├── respostas/        # standardized JSON responses
    ├── router/           # route definitions and grouping
    └── security/         # password hashing/verification
```

### Running the project

**1. Requirements:** Go installed and a reachable MySQL instance.

**2. Database** — run the script (this creates the `redesocial` database from scratch):

```bash
mysql -u root -p < sql/sql.sql
mysql -u root -p redesocial < sql/dados.sql   # optional, sample data
```

**3. Environment variables** — create a `.env` file inside `API/`:

```env
API_PORT=5000
DB_USUARIO=root
DB_SENHA=your_password
DB_NOME=redesocial
SECRET_KEY=any_secret_key
```

**4. Install dependencies and run:**

```bash
go mod download
go run main.go
```

The API starts at `http://localhost:5000` (or whatever port is set in `API_PORT`).

### Endpoints

Routes marked 🔒 require the `Authorization: Bearer <token>` header.

#### Authentication

| Method | Route | Description |
| --- | --- | --- |
| POST | `/login` | Authenticates and returns the JWT token |

#### Users

| Method | Route | Description |
| --- | --- | --- |
| POST | `/usuarios` | Registers a new user |
| GET | `/usuarios?usuario=` 🔒 | Searches users by name or nick |
| GET | `/usuarios/{usuarioId}` 🔒 | Fetches a user by ID |
| PUT | `/usuarios/{usuarioId}` 🔒 | Updates name/nick/email (locked for 3 months after the last update) |
| DELETE | `/usuarios/{usuarioId}` 🔒 | Deletes the account |
| POST | `/usuarios/{usuarioId}/atualizar-senha` 🔒 | Changes the password (requires the current one) |
| POST | `/usuarios/{usuarioId}/seguir` 🔒 | Follows the user |
| POST | `/usuarios/{usuarioId}/parar-de-seguir` 🔒 | Unfollows |
| GET | `/usuarios/{usuarioId}/seguidores` 🔒 | Lists who follows the user |
| GET | `/usuarios/{usuarioId}/seguindo` 🔒 | Lists who the user follows |

#### Posts

| Method | Route | Description |
| --- | --- | --- |
| POST | `/publicacoes` 🔒 | Creates a post |
| GET | `/publicacoes` 🔒 | Lists the user's and followed users' posts |
| GET | `/publicacoes/buscar?termo=` 🔒 | Searches posts by content |
| GET | `/publicacoes/{publicacaoId}` 🔒 | Fetches a post by ID |
| PUT | `/publicacoes/{publicacaoId}` 🔒 | Updates a post (author only) |
| DELETE | `/publicacoes/{publicacaoId}` 🔒 | Deletes a post (author only) |
| GET | `/usuarios/{usuarioId}/publicacoes` 🔒 | Lists a specific user's posts |
| POST | `/publicacoes/{publicacaoId}/curtir` 🔒 | Likes the post |
| POST | `/publicacoes/{publicacaoId}/descurtir` 🔒 | Unlikes the post |

> ⚠️ `/publicacoes/buscar` is registered **before** `/publicacoes/{publicacaoId}` in the code — if the order were swapped, gorilla/mux would try to parse "buscar" as an ID.

### Request examples

```bash
# Login
curl -X POST http://localhost:5000/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@email.com", "senha": "123456"}'

# Create a post (authenticated)
curl -X POST http://localhost:5000/publicacoes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"conteudo": "Hello, world!"}'

# Search posts by content
curl -X GET "http://localhost:5000/publicacoes/buscar?termo=hello" \
  -H "Authorization: Bearer <token>"
```

### Database

Three tables: `usuarios`, `publicacoes` and `seguidores` (the user↔follower relationship table). All foreign keys use `ON DELETE CASCADE` — deleting a user automatically deletes their posts and connections. The `usuarios.atualizado_em` column stores when the user's data was last edited, used for the 3-month rule.

### Notes

A study project built to practice Go on the backend: REST routes, JWT authentication, plain SQL access with no ORM, layered architecture (controller → repository → database). For production, it would be worth tightening validation, handling environment variables more carefully, and reviewing CORS policy.
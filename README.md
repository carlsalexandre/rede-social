# Rede Social (Capybaras)

> 🇧🇷 [Português](#-português) · 🇺🇸 [English](#-english)

---

## 🇧🇷 Português

Rede social simples feita em **Go**, dividida em dois projetos independentes que conversam entre si:

- **[`API/`](API/README.md)** — API REST que guarda toda a regra de negócio e fala com o banco de dados (MySQL).
- **[`WebApp/`](WebApp/README.md)** — aplicação web que renderiza as páginas HTML e consome a API; é o que o usuário vê no navegador.

Projeto criado com fins de estudo, pra praticar Go no backend (API REST, autenticação com JWT, acesso a banco sem ORM) e no frontend server-side (`html/template`, sessão via cookie, Bootstrap sem framework de JS).

### Arquitetura

```text
┌────────────┐      HTML/AJAX      ┌────────────┐      HTTP + JWT      ┌───────────┐
│  Navegador │ ◄─────────────────► │   WebApp   │ ◄──────────────────► │    API    │ ◄──► MySQL
│            │   cookie de sessão  │  (Go, :8000) │   Bearer token      │ (Go, :5000) │
└────────────┘                     └────────────┘                       └───────────┘
```

O navegador nunca fala direto com a API nem guarda o token JWT — só interage com o WebApp, que mantém a sessão num cookie seguro e repassa as chamadas pra API usando o token por trás dos panos.

### Funcionalidades

- Cadastro, login e logout.
- Criar, curtir/descurtir e excluir publicações.
- Buscar pessoas e publicações pelo mesmo campo de busca.
- Perfil (o seu e o de qualquer pessoa), bloqueado pra quem não está logado.
- Seguir, deixar de seguir, ver listas de seguidores e seguindo.
- Editar nome/nick/e-mail (com regra de 3 meses entre atualizações), trocar senha, excluir conta.
- Interface com Bootstrap 5, alertas com SweetAlert2, sem depender de CDN — todas as bibliotecas ficam hospedadas dentro do próprio projeto.

### Estrutura do repositório

```text
rede-social/
├── API/            # API REST em Go + MySQL
│   └── README.md   # como rodar a API, endpoints, etc.
└── WebApp/         # cliente web server-side em Go
    └── README.md   # como rodar o WebApp, rotas, etc.
```

### Como rodar o projeto completo

1. **Banco de dados** — suba um MySQL e rode `API/sql/sql.sql` (veja detalhes no [README da API](API/README.md)).
2. **API** — configure o `.env` e rode `go run main.go` dentro de `API/`. Por padrão sobe em `http://localhost:5000`.
3. **WebApp** — configure o `.env` (apontando `API_URL` pra API do passo anterior) e rode `go run main.go` dentro de `WebApp/`. Por padrão sobe em `http://localhost:8000`.
4. Acesse `http://localhost:8000` no navegador.

Cada README específico (linkado acima) tem o passo a passo completo com as variáveis de ambiente.

### Tecnologias

| Camada | Tecnologias |
| --- | --- |
| API | Go, gorilla/mux, MySQL, JWT, bcrypt |
| WebApp | Go, gorilla/mux, html/template, Bootstrap 5, jQuery, SweetAlert2 |
| Banco | MySQL |

### Observações

Este é um projeto de estudo — a estrutura em duas aplicações separadas (API + cliente web) foi escolhida de propósito, pra praticar a comunicação entre serviços via HTTP em vez de tudo dentro de um único monólito. Funcionalidades como comentários de verdade, upload de foto de perfil e mensagens diretas ainda não foram implementadas.

---

## 🇺🇸 English

Simple social network built in **Go**, split into two independent projects that talk to each other:

- **[`API/`](API/README.md)** — REST API that holds all the business logic and talks to the database (MySQL).
- **[`WebApp/`](WebApp/README.md)** — web application that renders the HTML pages and consumes the API; this is what the user sees in the browser.

Built as a study project, to practice Go on the backend (REST API, JWT authentication, database access without an ORM) and on the server-side frontend (`html/template`, cookie-based sessions, Bootstrap with no JS framework).

### Architecture

```text
┌────────────┐      HTML/AJAX      ┌────────────┐      HTTP + JWT      ┌───────────┐
│  Browser   │ ◄─────────────────► │   WebApp   │ ◄──────────────────► │    API    │ ◄──► MySQL
│            │   session cookie    │  (Go, :8000) │   Bearer token      │ (Go, :5000) │
└────────────┘                     └────────────┘                       └───────────┘
```

The browser never talks directly to the API nor holds the JWT token — it only interacts with the WebApp, which keeps the session in a secure cookie and forwards calls to the API using the token behind the scenes.

### Features

- Sign-up, login and logout.
- Create, like/unlike and delete posts.
- Search for people and posts through the same search field.
- Profile page (your own and anyone else's), locked for logged-out visitors.
- Follow, unfollow, view followers and following lists.
- Edit name/nick/email (with a 3-month rule between updates), change password, delete account.
- Bootstrap 5 interface, SweetAlert2 alerts, no CDN dependency — every library is hosted inside the project itself.

### Repository structure

```text
rede-social/
├── API/            # REST API in Go + MySQL
│   └── README.md   # how to run the API, endpoints, etc.
└── WebApp/         # server-side web client in Go
    └── README.md   # how to run the WebApp, routes, etc.
```

### Running the whole project

1. **Database** — spin up a MySQL instance and run `API/sql/sql.sql` (see the [API README](API/README.md) for details).
2. **API** — set up the `.env` file and run `go run main.go` inside `API/`. It starts at `http://localhost:5000` by default.
3. **WebApp** — set up the `.env` file (pointing `API_URL` to the API from the previous step) and run `go run main.go` inside `WebApp/`. It starts at `http://localhost:8000` by default.
4. Open `http://localhost:8000` in your browser.

Each specific README (linked above) has the full step-by-step, including environment variables.

### Tech stack

| Layer | Technologies |
| --- | --- |
| API | Go, gorilla/mux, MySQL, JWT, bcrypt |
| WebApp | Go, gorilla/mux, html/template, Bootstrap 5, jQuery, SweetAlert2 |
| Database | MySQL |

### Notes

This is a study project — splitting it into two separate applications (API + web client) was a deliberate choice, to practice service-to-service communication over HTTP instead of a single monolith. Features like real comments, profile photo upload and direct messages haven't been implemented yet.

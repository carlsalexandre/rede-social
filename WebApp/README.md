# WebApp — Rede Social

> 🇧🇷 [Português](#-português) · 🇺🇸 [English](#-english)

---

## 🇧🇷 Português

Aplicação web em Go que serve as páginas HTML da rede social (renderizadas no servidor) e consome a [`API`](../API/README.md) para todas as regras de negócio. Não fala com o banco de dados diretamente — todo o acesso a dados passa pela API, via requisições HTTP autenticadas.

### Visão geral

- Login, cadastro e logout, com sessão guardada num cookie seguro (assinado/criptografado, não um JWT no navegador).
- Feed com publicações, criar, curtir/descurtir e excluir publicação.
- Busca de pessoas e de publicações (mesma barra de busca).
- Página de perfil (o seu e o de qualquer pessoa), com bloqueio de conteúdo pra quem não está logado.
- Seguir / deixar de seguir, lista de seguidores e seguindo.
- Configurações: editar nome/nick/e-mail (respeitando a regra de 3 meses da API), trocar senha, excluir conta.
- Alertas e confirmações com [SweetAlert2](https://sweetalert2.github.io/) no lugar do `alert()`/`confirm()` do navegador.

### Tecnologias

- [Go](https://go.dev/) com [gorilla/mux](https://github.com/gorilla/mux)
- `html/template` da biblioteca padrão do Go, para renderizar as páginas no servidor
- [gorilla/securecookie](https://github.com/gorilla/securecookie) — cookie de sessão
- [Bootstrap 5](https://getbootstrap.com/) — grid e componentes (navbar, collapse)
- [jQuery](https://jquery.com/) — requisições AJAX e manipulação do DOM
- [SweetAlert2](https://sweetalert2.github.io/) — alertas e confirmações
- CSS próprio (`assets/css/home.css`), sem framework de UI além do Bootstrap
- Todas as bibliotecas de frontend ficam hospedadas localmente em `assets/`, sem depender de CDN

### Estrutura do projeto

```text
WebApp/
├── main.go
├── go.mod
├── config/            # leitura do .env
├── assets/
│   ├── css/           # bootstrap.css + home.css (estilos das páginas)
│   └── js/            # bootstrap.js, jquery.js, sweetalert2 e os scripts próprios
├── views/
│   ├── login.html, cadastro.html, home.html, usuarios.html,
│   │   perfil.html, conexoes.html, configuracoes.html
│   └── templates/
│       └── publicacoes.html   # templates nomeados reaproveitados no feed, na busca e no perfil
└── src/
    ├── controllers/    # monta a requisição pra API e renderiza o template
    ├── cookies/        # ler/gravar/apagar o cookie de sessão
    ├── middlewares/     # exige login nas rotas privadas
    ├── models/          # Usuario, Publicacao, Senha (espelham os da API)
    ├── requisicoes/      # helper que já inclui o Bearer token nas chamadas à API
    ├── respostas/         # padronização das respostas JSON
    ├── router/            # rotas do WebApp + servidor de arquivos estáticos
    └── utils/             # carregamento dos templates HTML
```

### Como as páginas se conectam com a API

O WebApp nunca expõe o token JWT pro navegador. O fluxo é:

1. O navegador manda um formulário (login, publicar, seguir etc.) pro **WebApp**.
2. O controller do WebApp lê o token guardado no cookie de sessão e faz a chamada HTTP pra **API**, com `Authorization: Bearer <token>`.
3. A resposta da API é processada e uma página HTML (ou um JSON, no caso das chamadas via AJAX) é devolvida ao navegador.

### Como executar

**1. Pré-requisitos:** Go instalado e a [API](../API/README.md) já rodando (o WebApp depende dela pra tudo).

**2. Variáveis de ambiente** — crie um arquivo `.env` na pasta `WebApp/`:

```env
APP_PORT=8000
API_URL=http://localhost:5000
HASH_KEY=uma_chave_de_32_bytes_qualquer
BLOCK_KEY=outra_chave_de_32_bytes
```

**3. Instale as dependências e rode:**

```bash
go mod download
go run main.go
```

O WebApp sobe em `http://localhost:8000` (ou na porta definida em `APP_PORT`).

> 💡 O `HASH_KEY`/`BLOCK_KEY` são usados pelo `gorilla/securecookie` pra assinar e criptografar o cookie de sessão — qualquer string longa o suficiente funciona em desenvolvimento, mas em produção deveriam ser geradas aleatoriamente e mantidas em segredo.

### Rotas

| Método | Rota | Descrição |
| --- | --- | --- |
| GET | `/`, `/login` | Tela de login (redireciona pra `/home` se já estiver logado) |
| POST | `/login` | Efetua o login |
| GET | `/logout` | Apaga o cookie e volta pro login |
| GET | `/criar-conta` | Tela de cadastro |
| POST | `/usuarios` | Cria a conta |
| GET | `/home` | Feed principal |
| POST | `/publicacoes` | Cria uma publicação |
| DELETE | `/publicacoes/{id}` | Exclui uma publicação (só o autor) |
| POST | `/publicacoes/{id}/curtir` \| `/descurtir` | Curtir / descurtir |
| GET | `/usuarios?usuario=` | Busca pessoas e publicações |
| GET | `/usuarios/{id}` | Perfil de um usuário (público, mas sem publicações pra quem não está logado) |
| PUT | `/usuarios/{id}` | Atualiza nome/nick/e-mail |
| POST | `/usuarios/{id}/atualizar-senha` | Troca a senha |
| DELETE | `/usuarios/{id}` | Exclui a conta |
| POST | `/usuarios/{id}/seguir` \| `/parar-de-seguir` | Seguir / deixar de seguir |
| GET | `/usuarios/{id}/seguidores` \| `/seguindo` | Listas de conexões |
| GET | `/configuracoes` | Editar perfil, trocar senha, excluir conta |

### Detalhes de implementação que valem a pena conhecer

- **Templates nomeados**: o card de publicação existe em duas versões (`publicacao-com-permissao` e `publicacao-sem-permissao`, em `views/templates/publicacoes.html`) — a versão com a engrenagem de excluir só é usada quando `.AutorID == $.UsuarioID`. Esses templates são reaproveitados na Início, na busca e no perfil.
- **Cache dos arquivos estáticos**: o `fileServer` de `assets/` manda `Cache-Control: no-cache`, forçando o navegador a sempre confirmar com o servidor antes de reusar um CSS/JS antigo — evita o clássico problema de editar o CSS e o navegador continuar mostrando a versão velha.
- **Perfil sem login**: a rota `/usuarios/{id}` é pública, mas o controller confere se existe cookie de sessão **antes** de chamar a API — sem login, nenhuma publicação é buscada e a página mostra um aviso pedindo pra entrar.

### Observações

Projeto de estudo, focado em praticar renderização no servidor com Go (`html/template`), consumo de uma API própria, e montagem de uma interface com Bootstrap sem framework de frontend. Funcionalidades como comentários, upload de foto de perfil e mensagens diretas ficaram fora do escopo por enquanto.

---

## 🇺🇸 English

Go web application that serves the social network's HTML pages (server-side rendered) and consumes the [`API`](../API/README.md) for all business logic. It never talks to the database directly — every data access goes through the API, via authenticated HTTP requests.

### Overview

- Login, sign-up and logout, with the session kept in a secure cookie (signed/encrypted, not a JWT exposed to the browser).
- Feed with posts, creating, liking/unliking and deleting posts.
- Search for people and posts (same search bar).
- Profile page (your own and anyone else's), with content locked for logged-out visitors.
- Follow / unfollow, followers and following lists.
- Settings: edit name/nick/email (respecting the API's 3-month rule), change password, delete account.
- Alerts and confirmations via [SweetAlert2](https://sweetalert2.github.io/) instead of the browser's native `alert()`/`confirm()`.

### Tech stack

- [Go](https://go.dev/) with [gorilla/mux](https://github.com/gorilla/mux)
- Go standard library's `html/template`, for server-side page rendering
- [gorilla/securecookie](https://github.com/gorilla/securecookie) — session cookie
- [Bootstrap 5](https://getbootstrap.com/) — grid and components (navbar, collapse)
- [jQuery](https://jquery.com/) — AJAX requests and DOM manipulation
- [SweetAlert2](https://sweetalert2.github.io/) — alerts and confirmations
- Hand-written CSS (`assets/css/home.css`), no UI framework beyond Bootstrap
- All frontend libraries are hosted locally under `assets/`, no CDN dependency

### Project structure

```text
WebApp/
├── main.go
├── go.mod
├── config/            # reads the .env file
├── assets/
│   ├── css/           # bootstrap.css + home.css (page styles)
│   └── js/            # bootstrap.js, jquery.js, sweetalert2 and the app's own scripts
├── views/
│   ├── login.html, cadastro.html, home.html, usuarios.html,
│   │   perfil.html, conexoes.html, configuracoes.html
│   └── templates/
│       └── publicacoes.html   # named templates reused in the feed, search and profile
└── src/
    ├── controllers/    # builds the API request and renders the template
    ├── cookies/        # read/write/clear the session cookie
    ├── middlewares/     # requires login on private routes
    ├── models/          # Usuario, Publicacao, Senha (mirror the API's)
    ├── requisicoes/      # helper that already attaches the Bearer token to API calls
    ├── respostas/         # standardized JSON responses
    ├── router/            # WebApp routes + static file server
    └── utils/             # HTML template loading
```

### How pages talk to the API

The WebApp never exposes the JWT to the browser. The flow is:

1. The browser submits a form (login, post, follow, etc.) to the **WebApp**.
2. The WebApp controller reads the token stored in the session cookie and calls the **API** with `Authorization: Bearer <token>`.
3. The API's response is processed, and either an HTML page or a JSON payload (for AJAX calls) is returned to the browser.

### Running the project

**1. Requirements:** Go installed and the [API](../API/README.md) already running (the WebApp depends on it for everything).

**2. Environment variables** — create a `.env` file inside `WebApp/`:

```env
APP_PORT=8000
API_URL=http://localhost:5000
HASH_KEY=any_32_byte_key
BLOCK_KEY=another_32_byte_key
```

**3. Install dependencies and run:**

```bash
go mod download
go run main.go
```

The WebApp starts at `http://localhost:8000` (or whatever port is set in `APP_PORT`).

> 💡 `HASH_KEY`/`BLOCK_KEY` are used by `gorilla/securecookie` to sign and encrypt the session cookie — any long-enough string works for development, but in production they should be randomly generated and kept secret.

### Routes

| Method | Route | Description |
| --- | --- | --- |
| GET | `/`, `/login` | Login screen (redirects to `/home` if already logged in) |
| POST | `/login` | Logs in |
| GET | `/logout` | Clears the cookie and returns to login |
| GET | `/criar-conta` | Sign-up screen |
| POST | `/usuarios` | Creates the account |
| GET | `/home` | Main feed |
| POST | `/publicacoes` | Creates a post |
| DELETE | `/publicacoes/{id}` | Deletes a post (author only) |
| POST | `/publicacoes/{id}/curtir` \| `/descurtir` | Like / unlike |
| GET | `/usuarios?usuario=` | Searches people and posts |
| GET | `/usuarios/{id}` | A user's profile (public, but posts are hidden for logged-out visitors) |
| PUT | `/usuarios/{id}` | Updates name/nick/email |
| POST | `/usuarios/{id}/atualizar-senha` | Changes password |
| DELETE | `/usuarios/{id}` | Deletes the account |
| POST | `/usuarios/{id}/seguir` \| `/parar-de-seguir` | Follow / unfollow |
| GET | `/usuarios/{id}/seguidores` \| `/seguindo` | Connection lists |
| GET | `/configuracoes` | Edit profile, change password, delete account |

### Implementation details worth knowing

- **Named templates**: the post card exists in two versions (`publicacao-com-permissao` and `publicacao-sem-permissao`, inside `views/templates/publicacoes.html`) — the version with the delete gear icon is only used when `.AutorID == $.UsuarioID`. These templates are reused across the home feed, search results and profile.
- **Static asset caching**: the `fileServer` serving `assets/` sends `Cache-Control: no-cache`, forcing the browser to always revalidate with the server before reusing an old CSS/JS file — this avoids the classic problem of editing a CSS file and the browser still showing the stale version.
- **Logged-out profile view**: the `/usuarios/{id}` route is public, but the controller checks for a session cookie **before** calling the API — without login, no posts are fetched and the page shows a message asking the visitor to sign in.

### Notes

A study project focused on practicing server-side rendering with Go (`html/template`), consuming a custom-built API, and assembling a UI with Bootstrap without a frontend framework. Features like real comments, profile photo upload and direct messages are still out of scope.
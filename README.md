# 🚀 Barber Analytics Pro — Backend v2.0

**Backend Go para SaaS de Gestão de Barbearias**

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-In%20Development-yellow)](https://github.com/andviana23/barber-analytics-backend-v2)

---

## 📋 Sobre o Projeto

Backend v2.0 do **Barber Analytics Pro** — migração arquitetural de MVP (React + Supabase) para plataforma SaaS escalável com:

- ✅ **Clean Architecture** + **DDD** + **SOLID**
- ✅ **Multi-tenancy** column-based (tenant_id)
- ✅ **PostgreSQL 14+** (Neon serverless)
- ✅ **JWT RS256** para autenticação
- ✅ **Cron jobs** para automação (Asaas sync, alertas)
- ✅ **Echo v4** (framework HTTP)
- ✅ **SQLC** (type-safe SQL)

---

## 🏗️ Arquitetura

```
barber-analytics-backend-v2/
├── cmd/
│   └── api/
│       └── main.go              # Entry point
├── internal/
│   ├── config/                  # Configuração
│   ├── domain/                  # Entidades e lógica de negócio
│   │   ├── entity/              # User, Tenant, Receita, etc
│   │   ├── valueobject/         # Email, Money, etc
│   │   └── service/             # Domain Services
│   ├── application/             # Use Cases
│   │   ├── dto/                 # Data Transfer Objects
│   │   ├── mapper/              # Domain ↔ DTO
│   │   └── usecase/             # Auth, Financial, etc
│   └── infrastructure/          # Implementações externas
│       ├── http/                # Handlers HTTP
│       ├── repository/          # PostgreSQL repositories
│       ├── external/            # Asaas API client
│       └── scheduler/           # Cron jobs
├── migrations/                  # SQL migrations
├── tests/                       # Testes (unit, integration, e2e)
├── .env.example                 # Exemplo de variáveis de ambiente
├── Dockerfile                   # Multi-stage build
├── docker-compose.yml           # Dev environment
└── Makefile                     # Comandos comuns
```

---

## 🚀 Quick Start

### Pré-requisitos

- Go 1.22+
- Docker & Docker Compose
- PostgreSQL 14+ (ou Neon account)
- Make (opcional)

### Setup Local

```bash
# 1. Clonar repositório
git clone https://github.com/andviana23/barber-analytics-backend-v2.git
cd barber-analytics-backend-v2

# 2. Instalar dependências
go mod download

# 3. Configurar ambiente
cp .env.example .env
# Editar .env com suas credenciais

# 4. Subir banco de dados (Docker)
docker-compose up -d db

# 5. Rodar migrations
make migrate-up

# 6. Iniciar servidor
go run cmd/api/main.go
# ou
make run
```

**Servidor rodando em:** http://localhost:8080

---

## 🛠️ Tecnologias

| Categoria | Stack |
|-----------|-------|
| **Linguagem** | Go 1.22+ |
| **Framework HTTP** | Echo v4 |
| **Database** | PostgreSQL 14+ (Neon) |
| **ORM/Queries** | SQLC (type-safe SQL) |
| **Auth** | JWT (RS256) + Refresh Tokens |
| **Validation** | go-playground/validator/v10 |
| **Scheduler** | robfig/cron/v3 |
| **Logger** | Zap (structured JSON) |
| **Migrations** | golang-migrate/migrate |
| **Testing** | testing.T + testify |
| **CI/CD** | GitHub Actions |
| **Containerização** | Docker + Docker Compose |

---

## 📚 Documentação

Documentação completa em `/docs`:

- [ARQUITETURA.md](../docs/ARQUITETURA.md) — Clean Architecture + DDD
- [GUIA_DEV_BACKEND.md](../docs/GUIA_DEV_BACKEND.md) — Setup e convenções
- [API_REFERENCE.md](../docs/API_REFERENCE.md) — Endpoints documentados
- [DOMAIN_MODELS.md](../docs/DOMAIN_MODELS.md) — Entities e Value Objects
- [BANCO_DE_DADOS.md](../docs/BANCO_DE_DADOS.md) — Schema e migrations
- [MODELO_MULTI_TENANT.md](../docs/MODELO_MULTI_TENANT.md) — Estratégia multi-tenancy

---

## 🔐 Segurança

- ✅ **Multi-tenancy**: Column-based com tenant_id em todas as queries
- ✅ **JWT RS256**: Tokens assinados com chave privada
- ✅ **Rate Limiting**: 30 req/s por IP, 100 req/s global
- ✅ **HTTPS**: SSL/TLS obrigatório (Let's Encrypt)
- ✅ **Auditoria**: Logs de todas as operações críticas
- ✅ **LGPD**: Compliance com right to be forgotten

---

## 🧪 Testes

```bash
# Rodar todos os testes
go test ./... -v

# Testes com coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Testes de integração
go test ./tests/integration -v

# Testes E2E
go test ./tests/e2e -v
```

**Target:** >80% coverage

---

## 🚢 Deploy

### Docker

```bash
# Build imagem
docker build -t barber-api:latest .

# Rodar container
docker run -p 8080:8080 --env-file .env barber-api:latest
```

### Docker Compose (Produção)

```bash
docker-compose -f docker-compose.prod.yml up -d
```

### CI/CD

- **Build automático** em push para `develop` e `main`
- **Deploy automático** em push para `main` (via GitHub Actions)
- **Testes obrigatórios** antes de merge

---

## 📊 Endpoints Principais

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/health` | Health check |
| POST | `/auth/login` | Login (JWT) |
| POST | `/auth/refresh` | Refresh token |
| GET | `/financial/receitas` | Listar receitas |
| POST | `/financial/receitas` | Criar receita |
| GET | `/financial/despesas` | Listar despesas |
| POST | `/financial/despesas` | Criar despesa |
| GET | `/financial/cashflow` | Fluxo de caixa |
| GET | `/subscriptions` | Listar assinaturas |
| POST | `/subscriptions` | Criar assinatura |

**Documentação completa:** [API_REFERENCE.md](../docs/API_REFERENCE.md)

---

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma feature branch (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'feat: Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

**Leia:** [CONTRIBUTING.md](CONTRIBUTING.md) para mais detalhes.

---

## 📝 Convenções de Código

- **Arquivos:** `snake_case` (user_repository.go)
- **Tipos:** `PascalCase` (User, UserRepository)
- **Funções públicas:** `PascalCase` (CreateUser)
- **Funções privadas:** `camelCase` (validateEmail)
- **Commits:** [Conventional Commits](https://www.conventionalcommits.org/)

---

## 📄 Licença

Este projeto está sob a licença **MIT**. Veja [LICENSE](LICENSE) para mais detalhes.

---

## 👥 Time

- **Arquiteto:** Andrey Viana
- **Backend Lead:** Andrey Viana
- **DevOps:** Andrey Viana

---

## 🎯 Roadmap

- [x] Fase 0: Fundamentos (Nov 14-17)
- [ ] Fase 1: DevOps Base (Nov 17-24)
- [ ] Fase 2: Backend Core (Nov 24-Dez 08)
- [ ] Fase 3: Módulos Backend (Dez 08-Jan 05)
- [ ] Fase 4: Frontend (Dez 08-Jan 05) — paralelo
- [ ] Fase 5: Migração (Jan 05-Feb 02)
- [ ] Fase 6: Hardening (Feb 02-16)

**Meta:** MVP 2.0 live em **Janeiro 16-23, 2025** 🎯

---

## 📞 Contato

- **GitHub:** [@andviana23](https://github.com/andviana23)
- **Projeto:** [Barber Analytics Pro v2.0](https://github.com/andviana23/barber-analytics-backend-v2)

---

**Feito com ❤️ e Go**

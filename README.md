# User API

## Descrição
Este projeto é uma API de autenticação de usuários, utilizando JWT para gerenciamento de sessões. Ele permite registro, login, logout e recuperação de detalhes do usuário.

## Tecnologias Utilizadas
- Go (Golang)
- GORM para ORM
- Fiber para o framework web
- MySQL como banco de dados
- JWT para autenticação

## Pré-requisitos
É necessário ter instalado em sua máquina:
- Go (versão 1.16 ou superior)
- MySQL
- Git

## Configuração do Ambiente

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/marimunari/user-api
   cd user-api
2. **Crie um arquivo .env com as seguintes informações:**
 ```bash
   # Database configuration
    DB_HOST=
    DB_PORT=
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    DB_URL=user:password@tcp(server_address:port)/management
    
    # JWT secrets
    JWT_SECRET=
    REFRESH_JWT_SECRET=
3. **Execute o seguinte comando para instalar as bibliotecas necessárias:**
 ```bash
    go mod tidy
4. **Execute as migrações do banco de dados: O projeto inclui uma função de migração automática. Ao executar o servidor pela primeira vez, o banco de dados será configurado automaticamente.**

## Executando a Aplicação
  ```bash
     go run main.go
A API estará disponível em http://localhost:8080.

## Endpoints
```bash
    /api/register: Registra um novo usuário
    /api/login: Faz login do usuário
    /api/logout: Faz logout do usuário
    /api/user: Exibe detalhes do usuário (rota protegida)
    /api/refresh-token: Atualiza o token do usuário

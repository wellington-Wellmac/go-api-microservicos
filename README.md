Gestão de Usuários - Arquitetura de Microserviços em Go

Este projeto implementa uma API RESTful com microserviços em Golang, utilizando Docker e gRPC para comunicação interna.

1. Estrutura do Projeto

go-projects/
├── user-service/
│   ├── main.go
│   ├── handlers.go
│   ├── models.go
│   ├── routes.go
│   ├── Dockerfile
├── auth-service/
│   ├── main.go
│   ├── handlers.go
│   ├── routes.go
│   ├── Dockerfile
├── email-service/
│   ├── main.go
│   ├── handlers.go
│   ├── routes.go
│   ├── Dockerfile
├── docker-compose.yml
└── README.md


2. Configuração e Execução

Antes de rodar os serviços, é necessário instalar o Docker e o Docker Compose.

2.1. Construindo e rodando os contêineres

Execute o comando abaixo no terminal para subir os serviços:

docker-compose up --build

Isso criará e iniciará os seguintes serviços:

User Service (localhost:8081)

Auth Service (localhost:8082)

Email Service (localhost:8083)


Para parar os serviços:

docker-compose down


3. Comunicação entre os Microserviços

Os microserviços comunicam-se usando gRPC e REST:

O user-service gerencia os usuários.

O auth-service autentica os usuários e gera JWT tokens.

O email-service envia e-mails quando necessário.


Fluxo de comunicação:

1. O usuário cria uma conta no user-service.


2. O auth-service gera um JWT ao fazer login.


3. O email-service recebe um pedido do user-service para enviar um e-mail de boas-vindas.


4. Testes no Postman

4.1. Testando o user-service

Listar usuários

Método: GET

URL: http://localhost:8081/users


Buscar um usuário por ID

Método: GET

URL: http://localhost:8081/users/1


Criar um novo usuário

Método: POST

URL: http://localhost:8081/users

Body (JSON):


{
  "username": "joao",
  "email": "joao@example.com"
}



4.2. Testando o auth-service

Gerar token JWT

Método: POST

URL: http://localhost:8082/auth/login?username=joao



4.3. Testando o email-service

Enviar um e-mail

Método: POST

URL: http://localhost:8083/email/send

Body (JSON):


{
  "to": "usuario@example.com",
  "subject": "Bem-vindo!",
  "body": "Olá, seja bem-vindo ao nosso sistema!"
}



5. Monitoramento

5.1. Verificar logs dos contêineres

docker logs <nome_do_container>

Para listar os containers em execução:

docker ps

5.2. Acessar um contêiner

docker exec -it <nome_do_container> sh



6. Melhorias Futuras

Implementar banco de dados PostgreSQL para persistência de usuários.

Criar um serviço de notificações para acompanhar os eventos do sistema.

Melhorar autenticação com OAuth 2.0.

Adicionar testes unitários.


7. Conclusão

Este projeto demonstra a construção de microserviços escaláveis em Golang, utilizando Docker, gRPC, JWT e APIs RESTful. 
Ele serve como base para sistemas distribuídos modernos.

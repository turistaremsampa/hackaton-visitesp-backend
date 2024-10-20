
# Backend do Aplicativo de Turismo

Este projeto é um backend baseado em Go para gerenciar usuários, feedbacks e sugestões em um aplicativo de recomendações turísticas. Ele utiliza Gin como framework web e GORM para gerenciamento de banco de dados.

## Estrutura do Projeto

```
tourism-app-backend/
├── main.go               # Ponto de entrada da aplicação
│   
├── controllers/
│   ├── userController.go     # Controlador de usuários
│   ├── feedbackController.go # Controlador de feedbacks
│   └── suggestionController.go # Controlador de sugestões
├── models/
│   ├── user.go               # Modelo de usuário
│   ├── feedback.go           # Modelo de feedback
│   └── suggestion.go         # Modelo de sugestão
├── services/
│   ├── userService.go        # Lógica de negócios de usuários
│   ├── feedbackService.go    # Lógica de negócios de feedbacks
│   └── suggestionService.go  # Lógica de negócios de sugestões
├── repositories/
│   ├── userRepository.go     # Consultas ao banco de dados para usuários
│   ├── feedbackRepository.go # Consultas ao banco de dados para feedbacks
│   └── suggestionRepository.go # Consultas ao banco de dados para sugestões
├── routes/
│   └── routes.go             # Definições de rotas
├── database/
│   └── database.go           # Conexão e configuração do banco de dados
├── go.mod                    # Módulos Go
└── go.sum                    # Checksum dos módulos Go
```

## Instalação

### Pré-requisitos

- Go (v1.20 ou superior)
- MySQL

### Clonar o repositório

```bash
git clone https://github.com/turistaremsampa/hackaton-visitesp-backend.git
cd hackaton-visitesp-backend
```

### Instalar as dependências

```bash
go mod tidy
```

### Rodar a aplicação

```bash
go run .
```

A aplicação será executada em `http://localhost:8080`.

## Endpoints da API

### Usuário

- **POST** `/user/login`: Entrar com e-mail e senha 
- **POST** `/users`: Criar um novo usuário
- **GET** `/users`: Listar todos os usuários

### Feedback

- **POST** `/feedbacks`: Criar um novo feedback
- **GET** `/feedbacks`: Listar todos os feedbacks
- **GET** `/feedbacks/:id`: Obter feedback por ID

### Sugestão

- **POST** `/suggestions`: Criar uma nova sugestão
- **GET** `/suggestions`: Listar todas as sugestões
- **GET** `/suggestions/:id`: Obter sugestão por ID

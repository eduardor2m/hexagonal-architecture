# CRUD em Golang usando a arquitetura hexagonal
### Este repositório contém uma implementação de um CRUD básico usando a linguagem Golang e a arquitetura hexagonal.

## Arquitetura hexagonal
A arquitetura hexagonal, também conhecida como ports and adapters, é um estilo arquitetural que tem como objetivo tornar o código mais flexível e independente de infraestrutura. Essa arquitetura é baseada em três conceitos principais:

O núcleo da aplicação: é onde se encontram as regras de negócio e a lógica da aplicação.
Adapters: são os componentes que permitem a comunicação do núcleo com o mundo externo, como por exemplo os adaptadores de banco de dados, adaptadores de interface de usuário, adaptadores de API, etc.
Ports: são as interfaces que definem como o núcleo da aplicação se comunica com os adaptadores.
A arquitetura hexagonal é muito útil para projetos de médio e grande porte, pois torna o código mais modular, testável e escalável.

Executando o projeto
Para executar o projeto, é necessário ter o Golang instalado na sua máquina. Você pode baixar a última versão do Golang aqui.

Para clonar o repositório e executar o projeto, execute os seguintes comandos:

```bash
$ git clone https://github.com/eduardor2m/hexagonal-architecture.git
$ cd hexagonal-architecture
$ go run ./src/cmd/server/main.go 
```
Endpoints
O projeto contém os seguintes endpoints:

- GET /users: Retorna todos os usuários cadastrados
- GET /users/:id: Retorna um usuário pelo ID
- POST /users: Cria um novo usuário
- PUT /users/:id: Atualiza um usuário existente
- DELETE /users/:id: Remove um usuário pelo ID
 
## Contribuindo

Sinta-se à vontade para contribuir com o projeto. Basta fazer um fork do repositório, criar uma nova branch com a sua contribuição e abrir um pull request.

Licença
Este projeto está licenciado sob a Licença MIT. Leia o arquivo LICENSE para mais informações.

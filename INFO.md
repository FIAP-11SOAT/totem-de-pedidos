# ℹ️ Informações do Projeto - Totem de Pedidos

## Video

[Fase 1 - Totem de Pedido (DDD + Hexagonal + Docker)](https://github.com/FIAP-11SOAT/totem-de-pedidos)

## DDD - Domain-Driven Design

O mapeamento do domínio (event storming, contextos, etc) está disponível no Miro:  
[https://miro.com/app/board/uXjVIHWL0sE=/](https://miro.com/app/board/uXjVIHWL0sE=/)

## Como rodar o projeto

1. **Pré-requisitos**:  
   - Docker e Docker Compose instalados.

2. **Subindo o ambiente**:  
   Execute o comando abaixo na raiz do projeto:
   ```sh
   docker compose up --build
   ```
   Isso irá subir os containers da aplicação, banco de dados e executar as migrations automaticamente.

3. **Acessando a aplicação**:  
   O backend estará disponível em:  
   ```
   http://localhost:8080
   ```

## Documentação da API

O projeto disponibiliza três visualizações para a documentação OpenAPI:

- **Swagger UI**:  
  [http://localhost:8080/docs/swagger](http://localhost:8080/docs/swagger)

- **Redoc**:  
  [http://localhost:8080/docs/redoc](http://localhost:8080/docs/redoc)

- **Scalar**:  
  [http://localhost:8080/docs/scalar](http://localhost:8080/docs/scalar)

O arquivo OpenAPI bruto pode ser acessado em:  
[http://localhost:8080/docs/openapi.yaml](http://localhost:8080/docs/openapi.yaml)

## Repositório

O código-fonte está disponível em:  
[https://github.com/FIAP-11SOAT/totem-de-pedidos](https://github.com/FIAP-11SOAT/totem-de-pedidos)

## Outras informações

- O projeto segue arquitetura hexagonal e princípios de DDD.
- Utiliza PostgreSQL como banco de dados.
- Integração de pagamentos via Mercado Pago.
- Testes automatizados com `testify`.
- Variáveis de ambiente podem ser configuradas no arquivo `.env` (veja `.env.example`).
- O serviço expõe endpoints RESTful para clientes, pedidos, produtos, categorias e pagamentos.
- Documentação DDD e decisões arquiteturais estão detalhadas no [README.md](README.md).
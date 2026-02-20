# Orchestrator

API leve de orquestração para fluxos de ordem de serviço.  
Coordena chamadas para serviços externos (`os-service-api`, billing, execution e entity APIs) e aplica ações de compensação em caso de falhas.

## Fluxos Principais

- Iniciar orquestração de ordem de serviço
  - Cria Ordem de Serviço
  - Cria orçamento
  - Inicia execução
  - Compensa etapas anteriores se uma etapa seguinte falhar
- Cancelar orquestração de ordem de serviço
  - Carrega dados da Ordem de Serviço
  - Libera peças/suprimentos
  - Cancela execução
  - Cancela orçamento/pagamento
  - Cancela Ordem de Serviço

## Endpoints HTTP

- `GET /health`
- `POST /orchestrator/v1/service-orders`
- `POST /orchestrator/v1/service-orders/{id}/cancel`

## Variáveis de Ambiente

- `PORT` (padrão: `8080`)
- `OS_BASE_URL` (padrão: `http://os-service:8080`)
- `OS_AUTH_TOKEN` (opcional; quando definido, é enviado como token `Bearer` para a API de OS)
- `BILLING_BASE_URL` (padrão: `http://billing-service:8080`)
- `EXEC_BASE_URL` (padrão: `http://execution-service:8080`)
- `ENTITY_BASE_URL` (padrão: `http://entity-api-service:8080`)

## Execução

```bash
go run ./cmd/api
```

## Testes e Cobertura

Execute os testes focados na orquestração com cobertura:

```bash
go test ./internal/usecase ./internal/adapter/http/handlers ./internal/adapter/http/routes ./internal/adapter/clients -coverprofile=coverage -covermode=atomic
go tool cover -func coverage
```

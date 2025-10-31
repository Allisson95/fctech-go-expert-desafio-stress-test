# Stressr - CLI para Testes de Carga HTTP

Sistema CLI em Go para realizar testes de carga em serviços web com relatório detalhado.

## Características

- Executa requests HTTP concorrentes contra uma URL especificada
- Distribui requisições de acordo com o nível de concorrência definido
- Gera relatório com métricas detalhadas:
  - Tempo total de execução
  - Quantidade total de requests realizados
  - Quantidade de requests com status HTTP 200
  - Distribuição de outros códigos de status HTTP (404, 500, etc.)
  - Quantidade de erros de rede/timeout

## Instalação

### Via Go Install

```bash
go install github.com/allisson95/fctech-go-expert-desafio-stress-test@latest
```

### Build Local

```bash
git clone https://github.com/allisson95/fctech-go-expert-desafio-stress-test.git
cd fctech-go-expert-desafio-stress-test
go build -o stressr
```

## Uso

### Sintaxe

```bash
stressr --url=<URL> --requests=<total> --concurrency=<concorrência>
```

### Parâmetros

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requests a executar (padrão: 1)
- `--concurrency`: Número de chamadas simultâneas (padrão: 1)

### Exemplos

#### Teste simples com 100 requests e 10 workers concorrentes

```bash
./stressr --url=http://google.com --requests=100 --concurrency=10
```

#### Teste com 1000 requests e 50 workers

```bash
./stressr --url=https://example.com --requests=1000 --concurrency=50
```

### Exemplo de Saída

```
=== Stress Test Report ===
Target URL: http://example.com
Total time: 3.823239518s
Total requests: 100
Status codes:
  200: 100
Errors (network/timeouts/etc): 0
```

## Uso via Docker

### Build da Imagem

```bash
docker build -t stressr .
```

### Executar via Docker

```bash
docker run stressr --url=http://google.com --requests=1000 --concurrency=10
```

### Exemplo com URL externa

```bash
docker run stressr --url=https://jsonplaceholder.typicode.com/posts --requests=500 --concurrency=25
```

## Desenvolvimento

### Executar Testes

```bash
go test ./... -v
```

### Estrutura do Projeto

```
.
├── cmd/
│   └── root.go          # Comando principal da CLI (stressr)
├── internal/
│   └── stress/
│       ├── stress.go    # Lógica do stress test
│       └── stress_test.go # Testes unitários
├── main.go              # Entry point
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

## Tecnologias Utilizadas

- **Go 1.24**: Linguagem de programação
- **Cobra**: Framework CLI
- **Docker**: Containerização
- **net/http**: Cliente HTTP nativo do Go

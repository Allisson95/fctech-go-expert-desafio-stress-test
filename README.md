# Stressr ⚡️ — Testes de Carga HTTP

Aperte o pedal e descubra como sua API se comporta sob pressão — simples, rápido e sem complicação.

## ✨ O que ele faz

- Em poucas palavras, ele:
- Executa requests HTTP concorrentes contra uma URL especificada
- Distribui requisições de acordo com o nível de concorrência definido
- Gera relatório com métricas detalhadas:
  - Tempo total de execução
  - Quantidade total de requests realizados
  - Quantidade de requests com status HTTP 200
  - Distribuição de outros códigos de status HTTP (404, 500, etc.)
  - Quantidade de erros de rede/timeout

## 🚀 Início Rápido (Docker 🐳)

O jeito mais rápido de começar é com a imagem oficial no DockerHub. Um comando e bora testar:

```bash
docker run -it --rm allissonabn/stressr --url=https://example.com --requests=100 --concurrency=10
```

### ⚙️ Parâmetros

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requests a executar (padrão: 1)
- `--concurrency`: Número de chamadas simultâneas (padrão: 1)

### 🧪 Exemplos práticos

**Teste básico com 1000 requests:**
```bash
docker run -it --rm allissonabn/stressr --url=http://google.com --requests=1000 --concurrency=10
```

**Teste com API externa:**
```bash
docker run -it --rm allissonabn/stressr --url=https://jsonplaceholder.typicode.com/posts --requests=500 --concurrency=25
```

**Teste de alta carga:**
```bash
docker run -it --rm allissonabn/stressr --url=https://example.com --requests=10000 --concurrency=100
```

**Teste de resistência (stress test pesado):**
```bash
docker run -it --rm allissonabn/stressr --url=https://api.example.com/health --requests=50000 --concurrency=500
```

### 🧾 Exemplo de saída

```
=== Stress Test Report ===
Target URL: http://example.com
Total time: 3.823239518s
Total requests: 100
Status codes:
  200: 100
Errors (network/timeouts/etc): 0
```

## 📦 Outras formas de instalação

### 🐹 Via Go Install

Se você já tem Go por aí, dá pra usar o binário direto:

```bash
go install github.com/allisson95/fctech-go-expert-desafio-stress-test@latest
```

Depois, é só rodar:

```bash
stressr --url=https://example.com --requests=100 --concurrency=10
```

### 🛠️ Build local

Quer contribuir ou mexer no código? Facinho:

```bash
git clone https://github.com/allisson95/fctech-go-expert-desafio-stress-test.git
cd fctech-go-expert-desafio-stress-test
go build -o stressr
./stressr --url=https://example.com --requests=100 --concurrency=10
```

### 🐳 Docker — build local da imagem

Prefere gerar sua própria imagem? Manda bala:

```bash
docker build -t stressr .
docker run -it --rm stressr --url=https://example.com --requests=1000 --concurrency=10
```

## 🛠️ Desenvolvimento

### ✅ Executar testes

```bash
go test ./... -v
```

### 📁 Estrutura do projeto

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

## 🧰 Tecnologias

- **Go 1.24** — Linguagem de programação
- **Cobra** — Framework CLI
- **Docker** — Containerização
- **net/http** — Cliente HTTP nativo do Go

## 💡 Dicas importantes

- Use com responsabilidade: não teste serviços sem autorização.
- Comece pequeno e vá aumentando a carga aos poucos (evita sustos e bloqueios).
- Tenha atenção a limites de rate limit/Firewall do destino.
- Em ambientes corporativos, valide proxies e regras de rede antes do teste.

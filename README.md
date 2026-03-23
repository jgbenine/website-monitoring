# Website Monitoring (Go) — Em desenvolvimento

Projeto simples em Go para **monitorar a disponibilidade de sites** via linha de comando, fazendo requisições HTTP e avaliando o `StatusCode` retornado.

## O que esse programa faz

- Exibe um menu no terminal:
  - `1` inicia o monitoramento
  - `2` (placeholder) ver logs
  - `0` encerra o programa
- Quando o monitoramento inicia, o programa:
  - percorre uma lista de URLs (hardcoded no código)
  - faz `GET` em cada URL
  - imprime no terminal se o site respondeu com `200` (OK) ou se houve problema (outro status)
  - repete o ciclo `numberOfMonitoring` vezes, aguardando `delayOfMonitoring` segundos entre os ciclos

## Estrutura do projeto

- `monitoring.go`: código do CLI e lógica de monitoramento.

## Requisitos

- Go instalado (recomendado `>= 1.18`)
- Acesso à internet para testar as URLs

## (Opcional) Usar `go.mod`

Hoje o projeto roda com um único arquivo (`monitoring.go`). Se você quiser transformar em um módulo Go (facilita `go run .` e organização do projeto), você pode iniciar assim:

```bash
go mod init website-monitoring
```

## Como executar

Na raiz do projeto:

```bash
go run monitoring.go
```

Depois, escolha uma opção no menu (por exemplo, `1` para iniciar o monitoramento).

## Como compilar (build)

```bash
go build -o website-monitoring monitoring.go
./website-monitoring
```

## Como personalizar

Edite `monitoring.go`:

- Quantidade de ciclos:
  - `const numberOfMonitoring = 2`
- Intervalo entre ciclos (em segundos):
  - `const delayOfMonitoring = 5`
- Lista de sites:
  - slice `websites := []string{ ... }` dentro de `initialMonitoring()`

## Observações e limitações atuais

- A opção **“View logs” (2)** ainda não está implementada (apenas imprime uma mensagem).
- O código atualmente ignora o erro de `http.Get` (uso de `_`), então falhas de rede/DNS podem não ser reportadas da melhor forma.
- Não há timeout explícito para as requisições HTTP (pode ficar aguardando dependendo do ambiente).

## Ideias de evolução

- Registrar logs em arquivo (e ler/mostrar na opção `2`)
- Ler a lista de URLs de um arquivo (ex.: `sites.txt`) ou via flags (`-site`, `-delay`, `-cycles`)
- Adicionar `http.Client` com timeout
- Fazer monitoramento concorrente (goroutines) e consolidar resultados

## go-mail-service

Projeto simples em Go para envio de emails via API e utilitários CLI.

Este repositório contém:

- Um servidor HTTP (em `cmd/api`) com handlers para gerar/validar API keys e enviar emails.
- Um cliente CLI (em `cmd/cli`) para gerar API keys e enviar requests de envio de email.
- Lógica de geração/validação de API keys em `internal/apikey` (salva em `data/apikeys.json`).
- Serviço de envio de email em `internal/services` e pacote `pkg/mailer`.

## Pré-requisitos

- Go 1.20+ instalado

## Build rápido

Na raiz do repositório, para compilar a CLI:

```bash
go build -o emailcli ./cmd/cli
```

Para rodar o servidor API (necessário se quiser testar o endpoint HTTP de envio):

```bash
go run ./cmd/api
```

Observação: alguns exemplos abaixo usam o binário `emailcli` gerado; se preferir, você também pode usar `go run` dentro de `cmd/cli`.

## Gerar uma API Key (CLI)

Gera uma nova chave e grava em `data/apikeys.json` (caminho relativo ao diretório onde o binário é executado).

```bash
./emailcli apikey generate
# Exemplo de saída: sk_3f5b9c2a... (a chave é impressa no stdout)
```

Guarde essa chave para autenticar requests ao endpoint de envio de email.

## Enviar email (CLI)

O comando `send` faz um POST para o endpoint HTTP especificado (por padrão `http://localhost:8080/api/send-email`).

Flags disponíveis (em `cmd/cli/send.go`):
- `--to`      : destinatário (string)
- `--subject` : assunto
- `--html`    : corpo em HTML
- `--from`    : remetente
- `--key`     : API key (Bearer)
- `--url`     : URL do endpoint (padrão: `http://localhost:8080/api/send-email`)

Exemplo de uso (usando a chave gerada anteriormente):

```bash
./emailcli send \
  --to user@example.com \
  --subject "Teste via CLI" \
  --html "<p>Olá via CLI</p>" \
  --from "noreply@example.com" \
  --key sk_3f5b9c2a... \
  --url http://localhost:8080/api/send-email
```

Observações:

- O comando `send` atual realiza a requisição HTTP e não imprime a resposta/erros detalhados por padrão — para melhor feedback, considere inspecionar o servidor API ou aprimorar o cliente CLI para mostrar status e corpo da resposta.
- O endpoint HTTP do servidor espera um JSON com campos `to` (array), `subject`, `html` e `from` e a autenticação via header `Authorization: Bearer <API_KEY>`.

## Arquivos importantes

- `internal/apikey/apikey.go` — geração e validação de chaves (persiste em `data/apikeys.json`).
- `cmd/cli` — implementação da CLI (comandos `send` e `apikey generate`).
- `cmd/api` — servidor HTTP e handlers (inclui handler para `GenerateKey`).

## Boas práticas e próximos passos

- Para produção, não salve chaves em um arquivo JSON local; use um banco de dados ou outro armazenamento seguro.
- Adicionar logs e tratamento de erros mais detalhado no comando `send` (para mostrar HTTP status, corpo e headers).
- Adicionar testes unitários para `internal/apikey` (geração + validação) e testes de integração para o fluxo CLI -> API.

## Contato

Se quiser que eu melhore o `send` para exibir resposta/erros, ou adicione uma flag de saída para `apikey generate` (ex.: `--out arquivo`), diga qual prefere e eu implemento.

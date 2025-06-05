# API com Golang

## Instalando o CompileDaemon

> Serve para rodar o código sem precisar reiniciar o servidor

```bash
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -command="./main"
```

---

### 📁 **Para que serve a pasta `vendor` no Go**

A pasta `vendor` é usada para armazenar **cópias locais das dependências** do seu projeto Go. Ou seja, em vez de buscar as bibliotecas do cache global (`~/go/pkg/mod`), o Go pode usar as versões salvas localmente dentro do projeto.

---

### 🧠 Por que usar o `vendor`?

1. ✅ **Build reprodutível:** garante que todas as máquinas (ou servidores) usem exatamente as mesmas versões das dependências.
2. 🚫 **Ambientes sem internet:** você pode compilar mesmo sem acesso à internet.
3. 🧳 **Distribuição offline:** útil em empresas que precisam auditar o código de terceiros.
4. 💥 **Evita conflitos com outras versões no cache do sistema.**

---

### 📦 Como gerar a pasta `vendor`

Se você quiser criar a pasta `vendor` com todas as dependências:

```bash
go mod vendor
```

Isso vai:

- Copiar todos os pacotes necessários (que estão no `go.mod`) para a pasta `./vendor`.

---

### ⚙️ Como usar o `vendor` durante a build

Se você quiser **forçar o uso do `vendor`**, use:

```bash
go build -mod=vendor
```

Ou:

```bash
go run -mod=vendor main.go
```

> Sem isso, o Go pode usar o cache global (em `go/pkg/mod`), ignorando a pasta `vendor`.

---

### 🗂️ Estrutura exemplo com `vendor`

```text
meu-projeto/
├── go.mod
├── go.sum
├── main.go
└── vendor/
    └── github.com/
        └── gin-gonic/
            └── gin/
```

---

### 🔥 Preciso usar `vendor` sempre?

**Não.**
O Go gerencia bem as dependências com o `go.mod` e `go.sum`.
Mas usar `vendor` **pode ser útil em ambientes corporativos, CI/CD ou deploy offline.**

---

Dá pra fazer isso sim — com algumas **observações importantes**.

### ✅ **Resumo rápido:**

> O Go **não baixa automaticamente as dependências na pasta `vendor`**, **a menos que você mande ele fazer isso** com `go mod vendor`.

Porém, com **um script ou alias**, dá pra simular isso e criar um comportamento tipo “modo vendor automático”.

---

### 🧠 Por que o Go não usa `vendor` por padrão?

Porque o `vendor` é considerado uma **opção especial**, usada principalmente em:

- ambientes corporativos,
- CI/CD,
- builds offline.

Por padrão, o Go usa o cache global (`~/go/pkg/mod`) — mais rápido e centralizado.

---

### 🚀 Mas e se eu quiser **forçar o uso do `vendor` sempre**?

Você pode:

#### 🛠️ 1. Gerar o `vendor` manualmente:

Sempre que instalar ou atualizar libs:

```bash
go get github.com/alguma/lib
go mod tidy
go mod vendor
```

Ou, em um comando só (via `.bat`):

```bat
@echo off
go get %1
go mod tidy
go mod vendor
```

Salve como `get.bat` e use assim:

```bash
get.bat github.com/gin-gonic/gin
```

---

#### 🧪 2. Rodar seu projeto usando o `vendor`:

Sempre que for compilar ou rodar:

```bash
go run -mod=vendor main.go
```

Ou num script `dev.bat`:

```bat
@echo off
go mod vendor
CompileDaemon -command="go run -mod=vendor main.go"
```

---

### 💡 Dica bônus: Forçar `vendor` no `go build`

Você também pode **forçar o uso de `vendor` por padrão** durante builds com:

```bash
go build -mod=vendor
```

---

### 📌 Conclusão

| Objetivo              | Comando                                         |
| --------------------- | ----------------------------------------------- |
| Instalar libs         | `go get pacote && go mod tidy && go mod vendor` |
| Rodar usando `vendor` | `go run -mod=vendor main.go`                    |
| Build usando `vendor` | `go build -mod=vendor`                          |
| Automatizar tudo      | Criar scripts `.bat` com esses comandos         |

---

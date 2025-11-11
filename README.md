# 🌊 Echo Surf
> **Um navegador de linha de comando feito para desenvolvedores.**

> Explore, analise e baixe conteúdo da web direto do terminal

---

## Sobre o projeto

O **Echo Surf** é um navegador de linha de comando escrito em **Go** que permite:

- Acessar e analisar páginas web
- Testar endpoints (GET, POST, PUT, PATCH, DELETE)
- Visualizar **título**, **links**, **imagens** e **headers**
- Fazer **downloads de vídeos e áudios** (via `yt-dlp`)
- Funcionar como uma ferramenta de estudo e inspeção da web

---

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/LucasPaulo001/Echo-Surf.git
   cd Echo-Surf
   ```

2. Instale as dependências
   ```bash
     go mod tidy
   ```

3. (Opcional, mas recomendado)
   ```bash
     sudo apt install yt-dlp ffmpeg
   ```

## Uso básico

1. Visualizar informações de uma página:
```bash
   go run ./cmd/main.go --url https://example.com
```

2. Exibir headers HTTP
```bash
   go run ./cmd/main.go --url https://exemple.com --headers
```

3. Download de vídeos e áudios (yt-dlp integrado)
```bash
   go run ./cmd/main.go --url https://exemple.com --download mp4 / mp3
```

## Testando APIs com o Echo Surf

1. Exemplo de teste
```bash
   go run ./cmd/main.go --url-test <link> --method <Passa o método após a flag> --body '{json}' --headers-test <headers após a flag> --json 
```
> --url-test => endpoint para ser testado

> --method => (GET, POST, PUT, PATCH, DELETE) *caso omitido será GET*

> --body => '{"email": "email@gmai.com", "password": "password123"}' *<= exemplo de body*

> --headers-test => "Content-Type:application/json"

> > se caso o header precisar de token "Content-Type:application/json;Authorization: Bearer 'token'"

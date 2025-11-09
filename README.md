# 🌊 Echo Surf
> **Um navegador de linha de comando feito para desenvolvedores.**

> Explore, analise e baixe conteúdo da web direto do terminal

---

## Sobre o projeto

O **Echo Surf** é um navegador de linha de comando escrito em **Go** que permite:

- Acessar e analisar páginas web
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
